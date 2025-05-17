package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


var DB *sql.DB

func InitDB () {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"),)

	var err error
	DB, err = sql.Open("mysql",dsn)

	if err != nil{
		fmt.Print("Error: ", err)
		panic("Cannot connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	fmt.Print("Successfully connected to the database!!")
}

