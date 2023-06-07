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

	JWT struct {
		TokenExpirationTime int64  `config:"JWT_EXPIRATION_TIME"`
		SecretKey           string `config:"JWT_SECRET"`
	}
}

func GetConfig(configPath string) (*Config, error) {
	var conf Config
	err := config.From(configPath).To(&conf.DB)
	if err != nil {
		return nil, err
	}

	err = config.From(configPath).To(&conf.JWT)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
