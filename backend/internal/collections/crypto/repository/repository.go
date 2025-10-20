package repository

import (
	collection "github.com/Jeje-Gab/cyber-av/backend/internal/collections/crypto"
)

// Implementação exemplo (memória). Mantém o padrão "temperature".
type memoryRepo struct{}

func NewMemory() collection.Repository {
	return &memoryRepo{}
}

// Exemplos de métodos quando a interface crescer:
// func (r *memoryRepo) SaveHistory(ctx context.Context, h *entity.History) (error, context.Context, int) {
// 	return nil, ctx, 0
// }
