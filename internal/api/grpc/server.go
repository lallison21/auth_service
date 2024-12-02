package grpc

import (
	"fmt"
	"github.com/lallison21/auth_service/internal/api"
	"github.com/lallison21/auth_service/internal/config/config"
	"github.com/lallison21/auth_service/pkg/grpc_stubs/auth_service"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	auth_service.UnimplementedAuthServiceServer
	service api.Service
	logger  *zerolog.Logger
}

func RunServer(cfg *config.GrpcConfig, log *zerolog.Logger, service api.Service) error {
	addr := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Errorf("[application.RunApi] failed to listen: %w", err))
	}

	s := grpc.NewServer()

	log.Info().Msgf("[application.RunApi] listening on %s", addr)
	auth_service.RegisterAuthServiceServer(s, &server{
		service: service,
		logger:  log,
	})
	if err := s.Serve(lis); err != nil {
		panic(fmt.Errorf("[application.RunApi] failed to serve: %w", err))
	}

	return nil
}
