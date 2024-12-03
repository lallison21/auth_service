package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lallison21/auth_service/internal/config/jwt_utils"
	"github.com/lallison21/auth_service/internal/config/logger"
	"github.com/lallison21/auth_service/internal/config/password"
	"github.com/lallison21/auth_service/internal/config/storage"
)

type Config struct {
	Grpc     GrpcConfig          `env:"GRPC"`
	Logger   logger.Logger       `env:"LOGGER"`
	Postgres storage.Postgres    `env:"POSTGRES"`
	Password password.Password   `env:"PASSWORD"`
	JWT      jwt_utils.JWTConfig `env:"JWT"`
}

type GrpcConfig struct {
	AppHost string `env:"GRPC_HOST" required:"true" env-default:"0.0.0.0"`
	AppPort string `env:"GRPC_PORT" required:"true" env-default:"60000"`
}

func MustEnv() *Config {
	cfg := Config{}

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(fmt.Errorf("[config.Env] read env: %w", err))
	}

	return &cfg
}
