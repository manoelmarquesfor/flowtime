package config

import (
	"errors"

	"github.com/BurntSushi/toml"
)

var ErrFalhaCarregarConfig = errors.New("falha ao carregar o arquivo de configuração")

func LoadConfig() (*Config, error) {
	var cfg Config

	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		return nil, ErrFalhaCarregarConfig
	}

	return &cfg, nil
}
