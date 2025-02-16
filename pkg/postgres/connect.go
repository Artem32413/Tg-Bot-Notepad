package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type DBInfo struct {
	user     string
	password string
	dbName   string
	dbport   string
	dbhost   string
}

func Connect() (*sqlx.DB, error) {
	
	var inf DBInfo
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	inf.LoadEnvValues()
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslMode port=%s host=%s",
		inf.user, inf.password, inf.dbName, inf.dbport, inf.dbhost)
	conn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	fmt.Println(conn)
	return conn, nil
}
