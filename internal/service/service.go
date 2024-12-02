package service

import (
	"context"
	"github.com/lallison21/auth_service/internal/models"
)

type Service struct {
	repository Repository
	passUtils  PasswordUtils
	jwtUtils   JWTUtils
}

func New(repo Repository, passUtils PasswordUtils, jwtUtils JWTUtils) *Service {
	return &Service{
		repository: repo,
		passUtils:  passUtils,
		jwtUtils:   jwtUtils,
	}
}

func (s *Service) Register(ctx context.Context, newUser *models.CreateUserDto) (int, error) {
	return 0, nil
}

func (s *Service) Login(ctx context.Context, newUser *models.LoginUserDto) (models.Tokens, error) {
	return models.Tokens{}, nil
}
