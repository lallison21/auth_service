package application

import (
	"github.com/lallison21/auth_service/internal/config/config"
	"github.com/lallison21/auth_service/internal/config/logger"
	"github.com/rs/zerolog"
)

type Application struct {
	cfg *config.Config
	log *zerolog.Logger
}

func New(cfg *config.Config) (*Application, error) {
	log := logger.New(cfg.Logger)

	return &Application{
		cfg: cfg,
		log: log,
	}, nil
}

func (a *Application) RunApi() {
	a.log.Info().Msgf("[RunApi] service started")
}
