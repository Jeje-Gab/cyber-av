package crypto

import (
	"context"

	"github.com/Jeje-Gab/cyber-av/backend/internal/dto"
)

// Segue o padrão do "temperature": Usecase com métodos da coleção.
// Mantém a assinatura (err, ctx, *resp) como no teu exemplo.
type Usecase interface {
	Encrypt(ctx context.Context, in dto.EncryptReq) (error, context.Context, *dto.EncryptResp)
	Decrypt(ctx context.Context, in dto.DecryptReq) (error, context.Context, *dto.DecryptResp)
}
