package database

import (
	"UrlShorter/internal/config"
	inMem "UrlShorter/internal/database/inMemory"
	"UrlShorter/internal/database/postgres"
	"UrlShorter/internal/models"
	"flag"
)

type database interface {
	GetUrl(models.Url) (models.Url, error)
	CreateUrl(models.Url) error
	GetMaxUrl() (models.Url, error)
}

func NewDatabase(cfg *config.Config) database {
	dFlag := *flag.Bool("d", false, "В зависимости от флага -d данные будут храниться либо в памяти(если нет флага -d), либо в PostgreSQl(если флаг присутствует)")
	flag.Parse()

	//Убрать !
	if !dFlag {
		return postgres.NewPostgres(cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.SslMode, cfg.TimeZone, cfg.DbName)
	}

	return inMem.NewInMemory()
}
