package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type Db struct {
}

var conn = tools.CreateConnection(Db{})


func (db Db) CreateConnection() *pgxpool.Pool {
	url := "postgres://postgres:89908990aSa@127.0.0.1/leftstat"
	conn, err := pgxpool.New(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
