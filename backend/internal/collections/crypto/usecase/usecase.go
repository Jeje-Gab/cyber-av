package usecase

import (
	"context"

	collection "github.com/Jeje-Gab/cyber-av/backend/internal/collections/crypto"
	"github.com/Jeje-Gab/cyber-av/backend/internal/dto"
)

// Usecase implementa crypto.Usecase e despacha p/ EncryptUC / DecryptUC.
type Usecase struct {
	encrypt *EncryptUC
	decrypt *DecryptUC
}

// NewUseCase injeta dependências (repo e demais serviços quando existirem)
// e retorna algo que cumpre a interface de topo collection.Usecase.
func NewUseCase(repo collection.Repository) collection.Usecase {
	return &Usecase{
		encrypt: NewEncryptUC(repo),
		// se Encrypt depender de Decrypt (ou vice-versa), injete aqui
		decrypt: NewDecryptUC(repo),
	}
}

// ===== Implementação da interface de topo =====

func (u *Usecase) Encrypt(ctx context.Context, in dto.EncryptReq) (error, context.Context, *dto.EncryptResp) {
	return u.encrypt.Execute(ctx, in)
}

func (u *Usecase) Decrypt(ctx context.Context, in dto.DecryptReq) (error, context.Context, *dto.DecryptResp) {
	return u.decrypt.Execute(ctx, in)
}
