package main

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"

	cryptoHandler "github.com/Jeje-Gab/cyber-av/backend/internal/collections/crypto/handler"
	cryptoRepo "github.com/Jeje-Gab/cyber-av/backend/internal/collections/crypto/repository"
	cryptoUC "github.com/Jeje-Gab/cyber-av/backend/internal/collections/crypto/usecase"

	"github.com/Jeje-Gab/cyber-av/backend/pkg/config"
)

func main() {
	cfg := config.Load()

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover(), middleware.Logger())

	repo := cryptoRepo.NewMemory()
	uc := cryptoUC.NewUseCase(repo)
	h := cryptoHandler.NewHTTP(uc)

	e.POST("/encrypt", h.Encrypt)
	e.POST("/decrypt", h.Decrypt)

	addr := ":" + cfg.Port
	if err := e.Start(addr); err != nil {
		panic(fmt.Errorf("server error: %w", err))
	}
}
