package usecase

import (
	"context"
	"errors"

	collection "github.com/Jeje-Gab/cyber-av/backend/internal/collections/crypto"
	"github.com/Jeje-Gab/cyber-av/backend/internal/dto"
)

type DecryptUC struct {
	repo collection.Repository
}

func NewDecryptUC(repo collection.Repository) *DecryptUC {
	return &DecryptUC{repo: repo}
}

func (u *DecryptUC) Execute(ctx context.Context, in dto.DecryptReq) (error, context.Context, *dto.DecryptResp) {
	if in.Key == "" {
		return errors.New("key is required"), ctx, nil
	}

	// TODO: implementar lógica de decrypt
	return nil, ctx, &dto.DecryptResp{Plaintext: ""}
}
