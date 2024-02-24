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

	// Парсим файл конфигурации
	//config, err := configuration.NewConfig(configPath)
	//if err != nil {
	//	logger.Fatalf("failed to Config: %v", err)
	//}

	// Клиент для реализации бизнес-логики
	//serviceClient := service.NewGlobalService(transfer.NewTransfer())

	//app.Initialization(serviceClient)

	//// Реализация методов grpc
	//implementationServer := role_service.NewRoleService(serviceClient)
	//
	// Создаём экземпляр grpc сервера
	//grpcClient := grpc.NewServerGRPC(config.GrpcConfig.Port, serviceClient, logger)

	// Запускаем компонент grpc сервера
	//app.NewApp(logger, grpcClient).Run(context.Background())

	startApp, err := app.NewApp(context.Background(), configPath, logger)
	if err != nil {
		logger.Fatalf("failed to app: %v", err)
		return
	}

	startApp.Run1(context.Background())
}
