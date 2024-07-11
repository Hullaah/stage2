package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Hullaah/stage2/models"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

const dropAll = `
	DELETE FROM membership;
	DELETE FROM "user";
	DELETE FROM organisation;
`

var dbConn = initializeDB()

func initializeDB() *pgx.Conn {
	godotenv.Load("../.env")
	var (
		host     = os.Getenv("STAGE2_POSTGRESQL_HOST")
		port     = os.Getenv("STAGE2_POSTGRESQL_PORT")
		database = os.Getenv("STAGE2_POSTGRESQL_DB")
		user     = os.Getenv("STAGE2_POSTGRESQL_USER")
		password = os.Getenv("STAGE2_POSTGRESQL_PWD")
	)
	ctx := context.Background()
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, database)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func CreateQueryEngine() *models.Queries {
	environment := os.Getenv("STAGE2_ENV")
	if environment == "test" {
		dbConn.Exec(context.Background(), dropAll)
	}
	return models.New(dbConn)
}
