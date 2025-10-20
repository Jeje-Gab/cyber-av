package dto

type EncryptReq struct {
	Plaintext string `json:"plaintext"`
	Key       string `json:"key"`
}
type EncryptResp struct {
	Ciphertext string `json:"ciphertext"`
}

type DecryptReq struct {
	Ciphertext string `json:"ciphertext"`
	Key        string `json:"key"`
}
type DecryptResp struct {
	Plaintext string `json:"plaintext"`
}
