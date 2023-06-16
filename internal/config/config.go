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
		AccessTokenExpTime  int64  `config:"JWT_ACCESS_EXPIRATION_TIME"`
		RefreshTokenExpTime int64  `config:"JWT_REFRESH_EXPIRATION_TIME"`
		SecretKey           string `config:"JWT_SECRET"`
	}

	Redis struct {
		Host string `config:"REDIS_HOST"`
		Port string `config:"REDIS_PORT"`
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

	err = config.From(configPath).To(&conf.Redis)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
