package config

import (
	"fmt"
	"log/slog"

	"github.com/kkyr/fig"
)

type Config struct {
	Host     string `fig: "host"`
	User     string `fig: "user"`
	Password string `fig: "password"`
	Port     string `fig: "port"`
	SslMode  string `fig: "sslMode"`
	TimeZone string `fig: "timeZone"`
	DbName   string `fig: "dbName"`
	RunIp    string `fig: "runIp"`
	RunPort  string `fig: "runPort"`
}

func Init() *Config {
	cfg := Config{}

	//загружаем из конфиг файла
	if err := fig.Load(&cfg,
		fig.File("config.yaml"),                             //имя конфиг файла
		fig.Dirs("configs", "../../configs", "../configs")); //путь к конфиг файлу
	err != nil {
		panic(fmt.Sprintf("func config.Init: %s", err))
	}

	slog.Info("Config loaded")
	return &cfg
}
