package db

import (
	"context"
	"fmt"
	"github.com/Hullaah/stage2/models"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func CreateQueryEngine() *models.Queries {
	godotenv.Load("/.env")
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		database = os.Getenv("DB_NAME")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PWD")
	)
	ctx := context.Background()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, database)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}
	return models.New(conn)
}
