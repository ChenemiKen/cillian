package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() error {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	conStr := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", dbUsername, dbPassword,
		dbHost, dbName)
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		return err
	}

	return testDB(db)
}

func testDB(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}

func ConnectDb(db *sql.DB) error {
	if err := db.Conn()
}
