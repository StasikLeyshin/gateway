package main

import (
	"context"
	"gateway/internal/app"
	"gateway/internal/app/configuration"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

const (
	configFolder = "config"
	logFolder    = "log"
)

func main() {
	var deployFolder string

	configFileName := "config.local.yaml"

	levelLogger := zap.DebugLevel
	fileNameLogger := "app.log"

	if IsDeploy := os.Getenv("IS_DEPLOY"); IsDeploy != "" {
		configFileName = "config.yaml"
		levelLogger = zap.InfoLevel
	}

	if path := os.Getenv("DEPLOY_PATH"); path != "" {
		deployFolder = path
	}

	// Путь до файла с конфигурацией проекта
	configPath := filepath.Join(deployFolder, configFolder, configFileName)

	// Путь до файла с логами
	loggerPath := filepath.Join(deployFolder, logFolder, fileNameLogger)

	// Создаём логгер
	logger := configuration.NewLogger("main", levelLogger, loggerPath)

	startApp, err := app.NewApp(context.Background(), configPath, logger)
	if err != nil {
		logger.Fatalf("failed to app: %v", err)
		return
	}

	startApp.Run1(context.Background())
}
