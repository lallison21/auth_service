package grpc

import (
	"fmt"
	"github.com/lallison21/auth_service/internal/config/config"
	"github.com/lallison21/auth_service/pkg/grpc_stubs/auth_service"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	auth_service.UnimplementedAuthServiceServer
}

func RunServer(log *zerolog.Logger, cfg config.GrpcConfig) {
	addr := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Errorf("[application.RunApi] failed to listen: %w", err))
	}

	s := grpc.NewServer()
	_, _ = s, lis

	log.Info().Msgf("[application.RunApi] listening on %s", addr)
	auth_service.RegisterAuthServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(fmt.Errorf("[application.RunApi] failed to serve: %w", err))
	}
}
