package mysql

import "os"

func (db *DBInfo) LoadEnvVars() {
	db.user = os.Getenv("MY_NAME")
	db.password = os.Getenv("MY_PASSWORD")
	db.dbName = os.Getenv("DB_NAME")
	db.dbport = os.Getenv("DB_PORT")
	db.dbhost = os.Getenv("DB_HOST")
}
