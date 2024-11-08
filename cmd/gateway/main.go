package main

import (
	"context"
	"gateway/internal/app"
	"gateway/internal/app/configuration"
	"os"
	"path/filepath"
)

func main() {

	configFileName := "config.local.yaml"
	configFolder := "config"

	if IsDeploy := os.Getenv("IS_DEPLOY"); IsDeploy != "" {
		configFileName = "config.yaml"
	}

	if path := os.Getenv("CONFIG_PATH"); path != "" {
		configFolder = path
	}

	// Файл с конфигурацией проекта
	configPath := filepath.Join(configFolder, configFileName)

	// Создаём логгер
	logger := configuration.NewLogger()

	startApp, err := app.NewApp(context.Background(), configPath, logger)
	if err != nil {
		logger.Fatalf("failed to app: %v", err)
		return
	}

	startApp.Run1(context.Background())
}
