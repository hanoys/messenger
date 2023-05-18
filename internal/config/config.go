package config

import (
	"github.com/JeremyLoy/config"
)

type Config struct {
	DB struct {
		User     string `config:"DB_USER"`
		Password string `config:"DB_PASSWORD"`
		Host     string `config:"DB_HOST"`
		Port     string `config:"DB_PORT"`
		Name     string `config:"DB_NAME"`
		URL      string `config:"DB_URL"`
	}
}

func GetConfig() (*Config, error) {
	var conf Config
	err := config.FromEnv().To(&conf.DB)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
