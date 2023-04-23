package main

import (
	"fmt"

	"smart-door/internal/applications/initiator"
	"smart-door/internal/config"
	"smart-door/pkg/logging"
	"smart-door/pkg/migrations"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()

	if err != nil {
		fmt.Errorf("failed to create logger: %v", err)
	}
	defer logger.Sync() //nolint:errcheck
	appLogger := logging.NewLogger(logger, "benches")
	undoLogger := otelzap.ReplaceGlobals(appLogger)
	defer undoLogger()

	logger.Info("config initializing")
	cfg := config.GetConfig()
	migrateManager := migrations.NewManager(fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host, cfg.PostgreSQL.Port, cfg.PostgreSQL.Database, "disable",
	))

	err = migrateManager.Migrate()
	if err != nil {
		logger.Fatal("migrate: ", zap.Error(err))
	}
	a, err := initiator.NewApp(cfg, appLogger)
	if err != nil {
		appLogger.Fatal("Error create app", zap.Error(err))
	}
	logger.Info("running application")
	a.Run(cfg)
}
