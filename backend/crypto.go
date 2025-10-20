package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

var aead cipher.AEAD

func initAEAD() error {
	kb64 := os.Getenv("AES_KEY_B64")
	if kb64 == "" {
		return errors.New("AES_KEY_B64 ausente (use 32 bytes em base64)")
	}
	key, err := base64.StdEncoding.DecodeString(kb64)
	if err != nil {
		return err
	}
	if l := len(key); l != 16 && l != 24 && l != 32 {
		return errors.New("chave inválida: tamanho deve ser 16/24/32 bytes")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	return initGCM(block)
}

func initGCM(block cipher.Block) error {
	var err error
	aead, err = cipher.NewGCM(block)
	return err
}

func encryptAES(plaintext string) (string, error) {
	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ct := aead.Seal(nil, nonce, []byte(plaintext), nil)
	// Guardamos nonce||ct em base64
	out := append(nonce, ct...)
	return base64.StdEncoding.EncodeToString(out), nil
}

func decryptAES(b64 string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", err
	}
	if len(raw) < aead.NonceSize() {
		return "", errors.New("ciphertext curto")
	}
	nonce, ct := raw[:aead.NonceSize()], raw[aead.NonceSize():]
	pt, err := aead.Open(nil, nonce, ct, nil)
	if err != nil {
		return "", err
	}
	return string(pt), nil
}
