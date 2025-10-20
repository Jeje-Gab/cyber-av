import axios from "axios";

const api = axios.create({
    baseURL: process.env.REACT_APP_API_URL || "http://localhost:8080",
});

export const enviarMensagemApi = (mensagem) =>
    api.post("/mensagem", { mensagem_clara: mensagem });

export const buscarMensagensApi = () => api.get("/mensagens");

export const decryptMensagemApi = (ciphertextBase64) =>
    api.post("/decrypt", { ciphertext_base64: ciphertextBase64 });

export default api;
