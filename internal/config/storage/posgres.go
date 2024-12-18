package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Postgres struct {
	PostgresHost     string        `env:"POSTGRES_HOST" env-default:"localhost"`
	PostgresqlPort   string        `env:"POSTGRES_PORT" env-default:"5432"`
	PostgresUser     string        `env:"POSTGRES_USER" env-default:"postgres"`
	PostgresPassword string        `env:"POSTGRES_PASSWORD" env-default:"postgres"`
	PostgresDb       string        `env:"POSTGRES_DB" env-default:"postgres"`
	MaxIdleConnTime  time.Duration `env:"MAX_IDLE_CONN_TIME" env-default:"5m"`
	MaxConn          string        `env:"MAX_CONN" env-default:"10"`
	ConnMaxLifetime  time.Duration `env:"CONN_MAX_LIFETIME" env-default:"10m"`
}

const (
	maxConn           = 50
	healthCheckPeriod = 3 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
)

func NewPostgres(cfg Postgres) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresqlPort,
		cfg.PostgresUser,
		cfg.PostgresDb,
		cfg.PostgresPassword,
	)

	poolConfig, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = maxConn
	poolConfig.HealthCheckPeriod = healthCheckPeriod
	poolConfig.MaxConnIdleTime = maxConnIdleTime
	poolConfig.MaxConnLifetime = maxConnLifetime
	poolConfig.MinConns = minConns

	connPool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	return connPool, nil
}
