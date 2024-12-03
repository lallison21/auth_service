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

func (s *Service) Login(ctx context.Context, login *models.LoginUserDto) (*models.Tokens, error) {
	existingUser, err := s.repository.GetUserByUsernameOrEmail(ctx, login.Username, login.Password)
	if err != nil {
		return nil, fmt.Errorf("[service.Login] get user: %w", err)
	}

	passMatch, err := s.passUtils.ComparePassword(login.Password, existingUser.Password)
	if err != nil {
		return nil, fmt.Errorf("[service.Login] verify password: %w", err)
	}
	if !passMatch {
		return nil, fmt.Errorf("[service.Login] verify password: %w", app_errors.ErrWrongCredentials)
	}

	accessToken, _, err := s.jwtUtils.GenerateAccessToken(existingUser.Id)
	if err != nil {
		return nil, fmt.Errorf("[service.Login] generate access token: %w", err)
	}
	refreshToken, _, err := s.jwtUtils.GenerateRefreshToken(existingUser.Id)
	if err != nil {
		return nil, fmt.Errorf("[service.Login] generate refresh token: %w", err)
	}

	userTokens := &models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return userTokens, nil
}
