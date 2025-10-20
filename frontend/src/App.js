import React, { useEffect, useState } from "react";
import { enviarMensagemApi, buscarMensagensApi, decryptMensagemApi } from "./api";

export default function App() {
    const [mensagem, setMensagem] = useState("");
    const [mensagens, setMensagens] = useState([]);
    const [dec, setDec] = useState({}); // id -> plaintext

    async function carregarMensagens() {
        const res = await buscarMensagensApi();
        setMensagens(res.data);
    }

    async function enviar() {
        if (!mensagem.trim()) return;
        await enviarMensagemApi(mensagem.trim());
        setMensagem("");
        await carregarMensagens();
    }

    async function decrypt(id, b64) {
        const res = await decryptMensagemApi(b64);
        setDec((d) => ({ ...d, [id]: res.data.mensagem_clara }));
    }

    useEffect(() => { carregarMensagens(); }, []);

    return (
        <div style={{ padding: 24, fontFamily: "system-ui, sans-serif", maxWidth: 780, margin: "0 auto" }}>
            <h2>Demo AES-GCM (AEAD)</h2>
            <p style={{ opacity: 0.8, marginTop: -8 }}>
                Backend: {process.env.REACT_APP_API_URL || "http://localhost:8080"}
            </p>

            <div style={{ display: "flex", gap: 8 }}>
                <input
                    value={mensagem}
                    onChange={(e) => setMensagem(e.target.value)}
                    placeholder="Digite a mensagem"
                    style={{ flex: 1, padding: 10, borderRadius: 8, border: "1px solid #ddd" }}
                />
                <button onClick={enviar} style={{ padding: "10px 16px", borderRadius: 8 }}>
                    Criptografar & Salvar
                </button>
            </div>

            <h3 style={{ marginTop: 24 }}>Mensagens</h3>
            {!mensagens.length && <div>Nenhuma mensagem ainda.</div>}
            <ul style={{ listStyle: "none", padding: 0 }}>
                {mensagens.map((m) => (
                    <li key={m.id} style={{ padding: 12, border: "1px solid #eee", borderRadius: 8, marginBottom: 12 }}>
                        <div><b>Criptografada (base64):</b> <code>{m.mensagem_criptografada}</code></div>
                        <div><b>Original (DB):</b> {m.mensagem_clara}</div>
                        <div style={{ marginTop: 8 }}>
                            <button onClick={() => decrypt(m.id, m.mensagem_criptografada)} style={{ padding: "6px 10px", borderRadius: 6 }}>
                                Decriptar via API
                            </button>
                        </div>
                        {dec[m.id] && (
                            <div style={{ marginTop: 8, background: "#f7f7f7", padding: 8, borderRadius: 6 }}>
                                <b>Decriptada (API):</b> {dec[m.id]}
                            </div>
                        )}
                    </li>
                ))}
            </ul>
        </div>
    );
}
