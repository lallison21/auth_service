package main

import (
	"github.com/lallison21/auth_service/internal/application"
	"github.com/lallison21/auth_service/internal/config/config"
)

func main() {
	cfg := config.MustEnv()

	app, err := application.New(cfg)
	if err != nil {
		panic(err)
	}

	app.RunApi()
}
