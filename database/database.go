package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseUrl string) (*pgxpool.Pool, error) {
	var err error

	connection, err := pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		return nil, err
	}

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// defer cancel()

	// oniChan := make(chan bool, 1)
	// go func(ch chan bool) {
	// 	Postgres.Connection.Ping(context.Background())
	// 	ch <- true
	// }(oniChan)

	// select {
	// case <-ctx.Done():
	// 	return nil, fmt.Errorf("Database Connection Timeout")
	// case <-oniChan:
	return connection, nil
	// }
}
