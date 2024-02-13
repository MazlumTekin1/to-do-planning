package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Initialize() *pgxpool.Pool {
	connectionPath := "postgresql://postgres:12345@localhost:5432/postgres"

	poolConfig, err := pgxpool.ParseConfig(connectionPath)
	if err != nil {
		log.Println("Unable to parse DATABASE_URL", "error", err)
		os.Exit(1)
	}

	connection, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Println("Unable to create connection pool, error:", err)
		os.Exit(1)
	} else {
		log.Println("Database Connection Created")
	}
	return connection
}
