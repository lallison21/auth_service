package application

import (
	"fmt"
	"github.com/lallison21/auth_service/internal/api/grpc"
	"github.com/lallison21/auth_service/internal/config/config"
	"github.com/lallison21/auth_service/internal/config/jwt_utils"
	"github.com/lallison21/auth_service/internal/config/logger"
	"github.com/lallison21/auth_service/internal/config/password"
	"github.com/lallison21/auth_service/internal/config/storage"
	"github.com/lallison21/auth_service/internal/repository"
	"github.com/lallison21/auth_service/internal/service"
	"github.com/lallison21/auth_service/version"
	"github.com/rs/zerolog"
)

type Application struct {
	cfg     *config.Config
	log     *zerolog.Logger
	service *service.Service
}

func New(cfg *config.Config) (*Application, error) {
	log := logger.New(cfg.Logger)

	postgresDb, err := storage.NewPostgres(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("[application.New] connection to postgres: %w", err)
	}

	authRepo := repository.New(postgresDb)

	passUtils := password.New(&cfg.Password)
	jwtUtils := jwt_utils.New(&cfg.JWT)

	authService := service.New(authRepo, passUtils, jwtUtils)

	return &Application{
		cfg:     cfg,
		log:     log,
		service: authService,
	}, nil
}

func (a *Application) RunApi() {
	a.log.Info().Msgf("[application.RunApi] version: %s name: %s commit: %s build time: %s",
		version.Version,
		version.Name,
		version.Commit,
		version.BuildTime,
	)

	grpc.RunServer(a.log, a.cfg.Grpc)
}
