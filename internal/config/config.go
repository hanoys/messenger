package config

import (
	"log"

	"github.com/JeremyLoy/config"
)

type Config struct {
	DB struct {
        User     string `config:"DB_USER"`
        Password string `config:"DB_PASSWORD"`
        Host     string `config:"DB_HOST"`
        Port     string `config:"DB_PORT"`
        DBName   string `config:"DB_NAME"`
	}
}

func GetConfig() *Config {
    var conf Config
    err := config.FromEnv().To(&conf.DB)
    if err != nil {
        log.Fatalf("configuration error: %v", err)
    }

    return &conf
}
