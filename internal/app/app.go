package app

import (
	"UrlShorter/internal/config"
	"UrlShorter/internal/database"
	"UrlShorter/internal/services"
	"UrlShorter/internal/transport"
	"log/slog"

	"github.com/labstack/echo/v4"
)

func Run() {
	cfg := config.Init()

	db := database.NewDatabase(cfg)
	service := services.NewService(cfg, db, db)
	transport := transport.NewTransport(service)

	e := echo.New()

	e.GET("/:shortUrl", transport.H.GetUrl)
	e.POST("/", transport.H.CreateUrl)

	slog.Info("Server running")

	err := e.Start(cfg.RunIp + ":" + cfg.RunPort)
	if err != nil {
		panic(err)
	}
}
