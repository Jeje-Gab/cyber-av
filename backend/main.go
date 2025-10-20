// main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	// Carrega variáveis do .env (opcional, mas prático no Windows)
	_ = godotenv.Load()

	// Banco + Cripto (usam DATABASE_URL e AES_KEY_B64 do .env)
	initDB()
	if err := initAEAD(); err != nil {
		log.Fatal("falha ao iniciar AEAD:", err)
	}

	// Rotas
	mux := http.NewServeMux()
	mux.HandleFunc("/mensagem", handlePostMensagem)        // POST cria+criptografa
	mux.HandleFunc("/mensagens", handleGetMensagens)       // GET lista
	mux.HandleFunc("/decrypt", handlePostDecrypt)          // POST decripta ciphertext arbitrário
	mux.HandleFunc("/decrypt-by-id", handleGetDecryptByID) // GET decripta pelo id salvo
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	addr := ":" + pick("PORT", "8080")
	log.Println("Servidor iniciado em", addr)

	// CORS global
	handler := corsMiddleware(mux)
	log.Fatal(http.ListenAndServe(addr, handler))
}

func pick(env, def string) string {
	if v := os.Getenv(env); v != "" {
		return v
	}
	return def
}

// ------------------------- CORS -------------------------

// CORS global simples e eficaz para dev local (3000 -> 8080)
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Como você não usa cookies/credenciais, "*" é suficiente
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			// Responde preflight e encerra
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ------------------------- Helpers -------------------------

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

// ------------------------- Handlers -------------------------

// POST /mensagem
// body: { "mensagem_clara": "..." }
func handlePostMensagem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var in struct {
		MensagemClara string `json:"mensagem_clara"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	ct, err := encryptAES(in.MensagemClara)
	if err != nil {
		http.Error(w, "erro ao criptografar", http.StatusInternalServerError)
		return
	}

	if err := inserirMensagem(in.MensagemClara, ct); err != nil {
		http.Error(w, "erro ao salvar", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{
		"mensagem_criptografada": ct,
	})
}

// GET /mensagens
func handleGetMensagens(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}
	msgs, err := buscarMensagens()
	if err != nil {
		http.Error(w, "erro ao buscar", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, msgs)
}

// POST /decrypt
// body: { "ciphertext_base64": "..." }
func handlePostDecrypt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var in struct {
		Ciphertext string `json:"ciphertext_base64"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	pt, err := decryptAES(in.Ciphertext)
	if err != nil {
		http.Error(w, "falha na decriptação", http.StatusBadRequest)
		return
	}

	// Evita cachear plaintext em proxies/navegador
	w.Header().Set("Cache-Control", "no-store")
	writeJSON(w, http.StatusOK, map[string]string{
		"mensagem_clara": pt,
	})
}

// GET /decrypt-by-id?id=123
func handleGetDecryptByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "parâmetro 'id' é obrigatório", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "parâmetro 'id' inválido", http.StatusBadRequest)
		return
	}

	m, err := buscarMensagemPorID(id)
	if err != nil {
		http.Error(w, "mensagem não encontrada", http.StatusNotFound)
		return
	}

	pt, err := decryptAES(m.MensagemCriptografada)
	if err != nil {
		http.Error(w, "falha na decriptação", http.StatusBadRequest)
		return
	}

	// Evita cachear plaintext em proxies/navegador
	w.Header().Set("Cache-Control", "no-store")
	writeJSON(w, http.StatusOK, map[string]any{
		"id":             m.ID,
		"mensagem_clara": pt,
	})
}
