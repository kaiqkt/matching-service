package di

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	username := GetEnv("POSTGRES_USERNAME")
	password := GetEnv("POSTGRES_PASSWORD")
	databaseName := GetEnv("POSTGRES_DATABASE")
	ip := GetEnv("POSTGRES_IP")
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", username, password, ip, databaseName)

	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
