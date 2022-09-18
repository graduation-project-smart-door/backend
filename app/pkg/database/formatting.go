package database

import "fmt"

func DatabaseParametersToDSN(engine string, host string, database string, user string, password string) string {
	// example: postgresql://localhost/mydb?user=other&password=secret
	return fmt.Sprintf("%s://%s/%s?user=%s&password=%s", engine, host, database, user, password)
}
