package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:postgres@localhost:5432/cyberav?sslmode=disable"
	}
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("erro ao abrir conexão:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("banco inacessível:", err)
	}
	log.Println("📦 PostgreSQL conectado")
}

type Mensagem struct {
	ID                    int    `json:"id"`
	MensagemClara         string `json:"mensagem_clara"`
	MensagemCriptografada string `json:"mensagem_criptografada"`
}

func inserirMensagem(mensagemClara, mensagemCriptografada string) error {
	_, err := db.Exec(
		`INSERT INTO mensagens (mensagem_clara, mensagem_criptografada) VALUES ($1,$2)`,
		mensagemClara, mensagemCriptografada,
	)
	return err
}

func buscarMensagens() ([]Mensagem, error) {
	rows, err := db.Query(`SELECT id, mensagem_clara, mensagem_criptografada FROM mensagens ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Mensagem
	for rows.Next() {
		var m Mensagem
		if err := rows.Scan(&m.ID, &m.MensagemClara, &m.MensagemCriptografada); err == nil {
			out = append(out, m)
		}
	}
	return out, nil
}

func buscarMensagemPorID(id int) (*Mensagem, error) {
	row := db.QueryRow(`SELECT id, mensagem_clara, mensagem_criptografada FROM mensagens WHERE id = $1`, id)
	var m Mensagem
	if err := row.Scan(&m.ID, &m.MensagemClara, &m.MensagemCriptografada); err != nil {
		return nil, err
	}
	return &m, nil
}
