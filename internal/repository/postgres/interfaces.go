package postgres

import (
	"database/sql"
)

type Client interface {
	Begin() (*sql.Tx, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	NamedExec(query string, arg any) (sql.Result, error)
}
