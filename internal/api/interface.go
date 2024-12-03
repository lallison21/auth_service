package api

import (
	"context"
	"github.com/lallison21/auth_service/internal/models"
)

type Service interface {
	Register(ctx context.Context, newUser *models.CreateUserDto) (int, error)
	Login(ctx context.Context, newUser *models.LoginUserDto) (*models.Tokens, error)
}
