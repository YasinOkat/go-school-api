package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URI")
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to verify connection to database!", err)
	}
}
