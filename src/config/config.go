package config

import (
	"os"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := os.Getenv("APP_DB_USERNAME")
	dbPass := os.Getenv("APP_DB_PASSWORD")
	dbName := os.Getenv("APP_DB_NAME")
	db, err = sql.Open(dbDriver, dbUser + ":" + dbPass + "@/" + dbName + "?parseTime=true")

	if err != nil {
        log.Fatal(err)
    }

	return
}