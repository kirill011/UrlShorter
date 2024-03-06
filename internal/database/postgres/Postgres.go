package postgres

import (
	"UrlShorter/internal/models"
	"fmt"
	"log/slog"

	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func (base *Postgres) GetUrl(short models.Url) (models.Url, error) {

	var url models.Url

	result := base.db.Select("FullUrl").Where(&short).First(&url)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return models.Url{}, models.ErrUrlNotFound
		}
		return models.Url{}, result.Error
	}

	return url, nil
}

func (base *Postgres) CreateUrl(url models.Url) error {
	result := base.db.Create(&url)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (base *Postgres) GetMaxUrl() (models.Url, error) {

	var url models.Url

	result := base.db.Select("ShortUrl").Last(&url)
	if result.Error != nil {
		return models.Url{}, result.Error
	}

	return url, nil
}

func NewPostgres(host, user, password, port, sslmode, timezone, dbName string) *Postgres {
	connString := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s TimeZone=%s", host, user, password, port, sslmode, timezone)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: slogGorm.New(),
	})

	if err != nil {
		panic(fmt.Sprintf("func postgres.NewPostgres: %s", err))
	}

	databaseString := fmt.Sprintf("%s dbname=%s", connString, dbName)
	dbNameExists := false
	db.Raw("SELECT exists(select * FROM pg_database WHERE datname = ?)", dbName).Scan(&dbNameExists)

	//Если на сервере нет базы данных dbName, создаём её
	if !dbNameExists {
		sql := fmt.Sprintf("CREATE DATABASE \"%s\";", dbName)
		db.Exec(sql)

		slog.Info("Database created")
	}

	db, err = gorm.Open(postgres.Open(databaseString), &gorm.Config{
		Logger: slogGorm.New(),
	})
	if err != nil {
		panic(fmt.Sprintf("func postgres.NewPostgres: %s", err))
	}

	err = db.AutoMigrate(&models.Url{})
	if err != nil {
		panic(fmt.Sprintf("func postgres.NewPostgres: %s", err))
	}

	return &Postgres{db: db}
}
