package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func Connect() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
	}

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")

	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":3306)/" + db_name

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
