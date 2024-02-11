package main

import (
	"gateway/internal/app/startup"
	"gateway/internal/service/service"
	"gateway/internal/transfer"
)

func main() {

	// Файл с конфигурацией проекта
	configPath := "config/config.local.yaml" //os.Getenv("CONFIG_PATH")

	// Создаём логгер
	logger := startup.NewLogger()

	// Парсим файл конфигурации
	config, err := startup.NewConfig(configPath)
	if err != nil {
		logger.Fatalf("failed to Config: %v", err)
	}

	// Клиент для реализации бизнес-логики
	serviceClient := service.NewGlobalService(transfer.NewTransfer())

	//app.Initialization(serviceClient)

	//// Реализация методов grpc
	//implementationServer := role_service.NewRoleService(serviceClient)
	//
	// Создаём экземпляр grpc сервера
	//grpcClient := grpc.NewServerGRPC(config.GrpcConfig.Port, serviceClient, logger)

	// Запускаем компонент grpc сервера
	//app.NewApp(logger, grpcClient).Run(context.Background())
}
