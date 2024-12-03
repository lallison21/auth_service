package service

import (
	"context"
	"github.com/lallison21/auth_service/internal/models"
)

type Repository interface {
	Register(ctx context.Context, newUser *models.UserDao) (int, error)
	GetUserByUsernameOrEmail(ctx context.Context, username, email string) (*models.UserDao, error)
}

type PasswordUtils interface {
	GeneratePassword(password string) (string, error)
	ComparePassword(password, hash string) (bool, error)
}

type JWTUtils interface {
	GenerateAccessToken(userId int) (string, int64, error)
	GenerateRefreshToken(userId int) (string, int64, error)
	VerifyToken(token string) (int, int64, error)
}
