package grpc

import (
	"context"
	"github.com/lallison21/auth_service/internal/models"
	"github.com/lallison21/auth_service/pkg/grpc_stubs/auth_service"
)

func (s *server) Register(ctx context.Context, request *auth_service.RegisterRequest) (*auth_service.RegisterResponse, error) {
	newUser := models.EmptyCreateUserDto().FromGRPC(request)

	id, err := s.service.Register(ctx, newUser)
	if err != nil {
		s.logger.Error().Msgf("[Register]: %v", err)
		return nil, err
	}

	return &auth_service.RegisterResponse{
		UserId: int32(id),
	}, nil
}

func (s *server) Login(ctx context.Context, request *auth_service.LoginRequest) (*auth_service.LoginResponse, error) {
	loginUser := models.EmptyLoginUserDto().FromGRPC(request)

	tokens, err := s.service.Login(ctx, loginUser)
	if err != nil {
		s.logger.Error().Msgf("[Login]: %v", err)
		return nil, err
	}

	return tokens.ToGRPC(), nil
}
