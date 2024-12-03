package service

import (
	"context"
	"fmt"
	"github.com/lallison21/auth_service/internal/app_errors"
	"github.com/lallison21/auth_service/internal/models"
	"net/mail"
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
	if newUser.Password != newUser.PasswordConfirm {
		return -1, fmt.Errorf("[service.Register] confirm password: %w", app_errors.ErrPassAndConfirmDoseNotMatch)
	}

	hashedPassword, err := s.passUtils.GeneratePassword(newUser.Password)
	if err != nil {
		return -1, fmt.Errorf("[service.Register] generating password: %w", err)
	}

	_, err = mail.ParseAddress(newUser.Email)
	if err != nil {
		return -1, fmt.Errorf("[service.Register] valid email: %w", err)
	}

	newUserDao := &models.UserDao{
		Username: newUser.Username,
		Password: hashedPassword,
		Email:    newUser.Email,
	}
	newUserId, err := s.repository.Register(ctx, newUserDao)
	if err != nil {
		return -1, fmt.Errorf("[service.Register] register user: %w", err)
	}

	return newUserId, nil
}

func (s *Service) Login(ctx context.Context, newUser *models.LoginUserDto) (models.Tokens, error) {
	return models.Tokens{}, nil
}
