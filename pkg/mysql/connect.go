package mysql

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func Connect() (*sqlx.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	user := os.Getenv("MY_NAME")
	password := os.Getenv("MY_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslMode port=%s host=%s",
		user, password, dbName, dbport, dbhost)
	conn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
