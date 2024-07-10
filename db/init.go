package db

import (
	"context"
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"
)

func InitializeDB() *pgx.Conn {
	godotenv.Load()
	var (
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		database = os.Getenv("DB_NAME")
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PWD")
	)
	ctx := context.Background()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, database)
	conn, err := pgx.Connect(ctx, connString)
}

