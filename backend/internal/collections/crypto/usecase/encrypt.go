package usecase

import (
	"context"
	"errors"

	collection "github.com/Jeje-Gab/cyber-av/backend/internal/collections/crypto"
	"github.com/Jeje-Gab/cyber-av/backend/internal/dto"
)

type EncryptUC struct {
	repo collection.Repository
}

func NewEncryptUC(repo collection.Repository) *EncryptUC {
	return &EncryptUC{repo: repo}
}

// Execute segue a mesma assinatura de retorno do teu padrão.
func (u *EncryptUC) Execute(ctx context.Context, in dto.EncryptReq) (error, context.Context, *dto.EncryptResp) {
	if in.Key == "" {
		return errors.New("key is required"), ctx, nil
	}

	// TODO: implementar lógica de encrypt (ex.: chamar pkg/crypto/AES)
	// resp := &dto.EncryptResp{Ciphertext: "..."}
	// return nil, ctx, resp

	return nil, ctx, &dto.EncryptResp{Ciphertext: ""}
}
