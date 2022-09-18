package main

import (
	"fmt"
	"go.uber.org/zap"
	"smart-door/app/internal/applications/initiator"
	"smart-door/app/internal/config"
	"smart-door/app/pkg/logging"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Errorf("failed to create logger: %v", err)
	}
	defer logger.Sync() // nolint:errcheck
	appLogger := logging.NewLogger(logger, "smart-door")
	logger.Info("init config")
	cfg := config.GetConfig()

	a, err := initiator.NewApp(cfg, appLogger)
	if err != nil {
		appLogger.Fatal("Error create app", zap.Error(err))
	}
	logger.Info("Running Application")
	a.Run()
}
