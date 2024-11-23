package application

import (
	"fmt"
	"github.com/lallison21/auth_service/internal/config/config"
	"github.com/lallison21/auth_service/internal/config/logger"
	"github.com/lallison21/auth_service/internal/config/storage"
	"github.com/lallison21/auth_service/internal/repository"
	"github.com/lallison21/auth_service/version"
	"github.com/rs/zerolog"
)

type Application struct {
	cfg *config.Config
	log *zerolog.Logger
}

func New(cfg *config.Config) (*Application, error) {
	log := logger.New(cfg.Logger)

	postgresDb, err := storage.NewPostgres(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("[application.New] connection to postgres: %w", err)
	}

	_ = repository.New(postgresDb)

	return &Application{
		cfg: cfg,
		log: log,
	}, nil
}

func (a *Application) RunApi() {
	a.log.Info().Msgf("[RunApi] service started")
	a.log.Info().Msgf("[RunApi] version: %s name: %s commit: %s build time: %s", version.Version, version.Name, version.Commit, version.BuildTime)
}
