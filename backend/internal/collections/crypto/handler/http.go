package handler

import (
	"net/http"

	collection "github.com/Jeje-Gab/cyber-av/backend/internal/collections/crypto"
	"github.com/Jeje-Gab/cyber-av/backend/internal/dto"
)

type HTTP struct {
	uc collection.Usecase
}

func NewHTTP(uc collection.Usecase) *HTTP {
	return &HTTP{uc: uc}
}

func (h *HTTP) Encrypt(c echo.Context) error {
	var in dto.EncryptReq
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	err, ctx, out := h.uc.Encrypt(c.Request().Context(), in)
	_ = ctx // mantém a assinatura padrão (ctx encadeável)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, out)
}

func (h *HTTP) Decrypt(c echo.Context) error {
	var in dto.DecryptReq
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	err, ctx, out := h.uc.Decrypt(c.Request().Context(), in)
	_ = ctx
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, out)
}
