CREATE TABLE IF NOT EXISTS mensagens (
    id SERIAL PRIMARY KEY,
    mensagem_clara TEXT,
    mensagem_criptografada TEXT,
    created_at TIMESTAMPTZ DEFAULT now()
);
