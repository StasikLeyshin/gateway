package main

import (
	"context"
	"gateway/internal/app"
	"gateway/internal/app/configuration"
)

func main() {

	// Файл с конфигурацией проекта
	configPath := "config/config.local.yaml" //os.Getenv("CONFIG_PATH")

	// Создаём логгер
	logger := configuration.NewLogger()

	startApp, err := app.NewApp(context.Background(), configPath, logger)
	if err != nil {
		logger.Fatalf("failed to app: %v", err)
		return
	}

	startApp.Run1(context.Background())
}
