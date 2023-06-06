package db

import (
	"os"
	"fmt"
	"server/utils"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func ConnectDB() (*Database, error) {
	err := godotenv.Load()

	if err != nil {
		utils.Logger(4, "db.go   -> godotenv.Load() -> Gagal membaca file .env")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOST")
	portnumb := os.Getenv("DB_PORT")
	dtbsname := os.Getenv("DB_NAME")
	debugger := os.Getenv("DEBUG_MODE")

	if debugger == "1" {
		utils.Logger(3, "db.go   -> godotenv.Load()")
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, portnumb, dtbsname))

	if debugger == "1" {
		if err != nil {
			utils.Logger(4, fmt.Sprintf("db.go   -> sql.Open() -> %s", err.Error()))
		} else{
			utils.Logger(3, "db.go   -> sql.Open()")
		}
	}

	err = db.Ping()
	if debugger == "1" {
		if err != nil {
			utils.Logger(4, fmt.Sprintf("db.go   -> db.Ping() -> %s", err.Error()))
		} else{
			utils.Logger(3, "db.go   -> db.Ping()")
		}
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}