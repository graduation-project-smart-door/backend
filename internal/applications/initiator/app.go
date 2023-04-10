package initiator

import (
	_ "smart-door/docs"

	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net"
	"net/http"
	"smart-door/internal/config"
	userPolicy "smart-door/internal/policy/user"
	userRepository "smart-door/internal/repository/postgres/user"
	userService "smart-door/internal/service/user"
	userHandlers "smart-door/internal/transport/httpv1/user"
	"smart-door/pkg/auth"
	postgresClient "smart-door/pkg/client/postgres"
	"time"
)

type App struct {
	config     *config.Config
	logger     *zap.Logger
	router     *mux.Router
	httpServer *http.Server
}

func NewApp(config *config.Config, logger *zap.Logger) (*App, error) {
	logger.Info("router initializing")
	router := mux.NewRouter()

	logger.Info("swagger initializing")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	logger.Info("database initializing")
	postgresConfig := postgresClient.NewConfig(
		config.PostgreSQL.Username, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database)
	database, errInitPostgres := postgresClient.NewClient(postgresConfig)

	if errInitPostgres != nil {
		logger.Fatal("create postgres client", zap.Error(errInitPostgres))
	}

	logger.Info("minio initializing")
	minioClient, err := minio.New(config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Minio.AccessKey, config.Minio.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		logger.Fatal("creating minio client", zap.Error(err))
	}
	err = minioClient.MakeBucket(context.Background(), config.Minio.Bucket, minio.MakeBucketOptions{
		Region:        "",
		ObjectLocking: false,
	})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(context.Background(), config.Minio.Bucket)
		if errBucketExists == nil && exists {
			logger.Info("we already own", zap.String("bucket", config.Minio.Bucket))
		} else {
			logger.Fatal("run minio client: ", zap.Error(err))
		}
	}

	logger.Info("redis initializing")
	_ = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	_, err = auth.NewManager(config.SigningKey)
	if err != nil {
		logger.Fatal("init auth manager", zap.Error(err))
	}

	// Пользователи
	logger.Info("user handlers initializing")
	appUserRouter := router.PathPrefix("/api/v1/users").Subrouter()
	appUserRepository := userRepository.NewRepository(database)
	appUserService := userService.NewService(logger, appUserRepository)
	appUserPolicy := userPolicy.NewPolicy(appUserService)
	appUserHandler := userHandlers.NewHandler(appUserPolicy)
	appUserHandler.Register(appUserRouter)

	return &App{config: config, logger: logger, router: router}, nil
}

func (app *App) startHTTP() {
	app.logger.Info("start HTTP")

	var listener net.Listener

	app.logger.Info(fmt.Sprintf("bind application to host: %s and port: %s", app.config.Listen.BindIP, app.config.Listen.Port))
	var err error
	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", app.config.Listen.BindIP, app.config.Listen.Port))
	if err != nil {
		app.logger.Fatal("", zap.Error(err))
	}

	c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodOptions, http.MethodDelete},
		AllowedOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Location", "Charset", "Access-Control-Allow-Origin", "Content-Type", "content-type", "Origin", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Location", "Authorization", "Content-Disposition"},
		Debug:              false,
	})

	handler := c.Handler(app.router)

	app.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	app.logger.Info("application completely initialized and started")

	if err := app.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			app.logger.Warn("server shutdown")
		default:
			app.logger.Fatal("error run listener", zap.Error(err))
		}
	}
	err = app.httpServer.Shutdown(context.Background())
	if err != nil {
		app.logger.Fatal("", zap.Error(err))
	}
}

func (app *App) Run() {
	app.startHTTP()
}
