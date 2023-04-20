package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppConfig struct {
		SigningKey                  string `env:"SIGNING_KEY"`
		IsDev                       bool   `env:"IS_DEV"`
		LifetimeAccessTokenMinutes  int    `env:"LIFETIME_ACCESS_TOKEN_MINUTES" env-default:"15"`
		LifetimeRefreshTokenMinutes int    `env:"LIFETIME_REFRESH_TOKEN_MINUTES" env-default:"30"`
	}

	TelegramBot struct {
		BaseURL string `env:"TELEGRAM_BOT_BASE_URL"`
	}

	Listen struct {
		BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port   string `env:"PORT" env-default:"8000"`
	}

	PostgreSQL struct {
		Username string `env:"POSTGRES_USER" env-required:"true"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
		Host     string `env:"POSTGRES_HOST" env-required:"true"`
		Port     string `env:"POSTGRES_PORT" env-required:"true"`
		Database string `env:"POSTGRES_DB" env-required:"true"`
	}

	Minio struct {
		Endpoint  string `env:"MINIO_ENDPOINT" env-default:"minio:9000"`
		AccessKey string `env:"MINIO_ACCESS_KEY"`
		SecretKey string `env:"MINIO_SECRET_KEY"`
		Bucket    string `env:"MINIO_BUCKET" env-default:"smart-door"`
		UseSSL    bool   `env:"MINIO_USE_SSL" env-default:"true"`
	}

	Redis struct {
		Host     string `env:"REDIS_HOST" env-default:"redis:6379"`
		Password string `env:"REDIS_PASSWORD" env-required:"true"`
		DB       int    `env:"REDIS_DB" env-default:"0"`
	}

	Images struct {
		PublicEndpoint string `env:"IMAGES_PUBLIC_ENDPOINT" env-default:"http://localhost:9000"`
	}
}

func GetConfig() *Config {
	log.Print("Get config")

	instance := &Config{}

	if err := cleanenv.ReadConfig(".env", instance); err != nil {
		helpText := "Error read env"
		help, _ := cleanenv.GetDescription(instance, &helpText)
		log.Print(help)
		log.Fatal(err)
	}
	return instance
}
