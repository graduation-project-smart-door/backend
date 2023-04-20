package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

type Client interface {
	Begin() (*sql.Tx, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	NamedExec(query string, arg any) (sql.Result, error)
}

type config struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func NewConfig(username string, password string, host string, port string, database string) *config {
	return &config{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func NewClient(cfg *config) (Client, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.Username, cfg.Password,
		cfg.Host, cfg.Port, cfg.Database, "disable",
	)

	db, errConnect := otelsqlx.Open("postgres", dsn,
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
		otelsql.WithDBName(cfg.Database))
	if errConnect != nil {
		log.Fatalf("Failed conntction to database: %v\n", errConnect)
		return nil, errConnect
	}

	return db, nil
}
