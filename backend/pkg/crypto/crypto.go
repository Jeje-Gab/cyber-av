package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func EncryptAES(plaintext, key []byte) (string, error) {
	if len(key) != 32 {
		return "", errors.New("key must be 32 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptAES(cipherB64 string, key []byte) ([]byte, error) {
	raw, err := base64.StdEncoding.DecodeString(cipherB64)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesGCM.NonceSize()
	if len(raw) < nonceSize {
		return nil, errors.New("invalid ciphertext")
	}
	nonce, ciphertext := raw[:nonceSize], raw[nonceSize:]
	return aesGCM.Open(nil, nonce, ciphertext, nil)
}
