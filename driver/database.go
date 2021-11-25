package driver

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// ConnectDB : Database connection
func ConnectDB() (db *sql.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("couldn't able to load .env file : %v", err)
	}
	dbDriver := os.Getenv("DATABASE_DRIVER")
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	port := os.Getenv("PORT")
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(:"+port+")/"+dbName)
	if err != nil {
		return nil, fmt.Errorf("couldn't able to open Mysql : %v", err)
	}
	return db, nil
}
