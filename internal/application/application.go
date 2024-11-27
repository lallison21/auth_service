package application

import (
	"fmt"
	"github.com/lallison21/auth_service/internal/config/config"
	"github.com/lallison21/auth_service/internal/config/logger"
	"github.com/lallison21/auth_service/internal/config/storage"
	"github.com/lallison21/auth_service/internal/repository"
	"github.com/lallison21/auth_service/internal/service"
	"github.com/lallison21/auth_service/pkg/grpc_stubs/auth_service"
	"github.com/lallison21/auth_service/version"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
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
	authService := service.New(authRepo)

	return &Application{
		cfg:     cfg,
		log:     log,
		service: authService,
	}, nil
}

type server struct {
	auth_service.UnimplementedAuthServiceServer
}

func (a *Application) RunApi() {
	a.log.Info().Msgf("[application.RunApi] version: %s name: %s commit: %s build time: %s",
		version.Version,
		version.Name,
		version.Commit,
		version.BuildTime,
	)

	addr := fmt.Sprintf("%s:%s", a.cfg.Grpc.AppHost, a.cfg.Grpc.AppPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Errorf("[application.RunApi] failed to listen: %w", err))
	}

	s := grpc.NewServer()
	_, _ = s, lis

	a.log.Info().Msgf("[application.RunApi] listening on %s", addr)
	auth_service.RegisterAuthServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(fmt.Errorf("[application.RunApi] failed to serve: %w", err))
	}
}
