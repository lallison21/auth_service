package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type Logger struct {
	LogIndex  string `env:"LOG_INDEX" required:"true" env-default:"auth_service"`
	IsDebug   bool   `env:"IS_DEBUG" required:"true" env-default:"false"`
	LogToFile bool   `env:"LOG_TO_FILE" required:"true" env-default:"true"`
}

func New(cfg Logger) *zerolog.Logger {
	logger := log.With().Str("service", cfg.LogIndex).Logger()

	if cfg.IsDebug {
		logger = logger.Level(zerolog.DebugLevel)
	} else {
		logger = logger.Level(zerolog.InfoLevel)
	}

	if cfg.LogToFile {
		file, err := os.Create("auth_service.log")
		if err != nil {
			log.Error().Err(err).Msg("Failed to create auth_service.log")
		}
		logger = logger.Output(file)
	} else {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	return &logger
}
