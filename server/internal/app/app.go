package app

import (
	"fmt"

	"github.com/rustem-yam/esoft/internal/config"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type application struct {
	logger *zap.Logger
	cfg    *config.Config
}

func New() (*application, error) {
	l, err := zap.NewProduction()

	if err != nil {
		return nil, err
	}

	var cfg config.Config

	err = cleanenv.ReadConfig("config.yml", &cfg)

	if err != nil {
		return nil, err
	}

	fmt.Println(cfg)

	return &application{
		logger: l,
		cfg:    &cfg,
	}, nil
}
