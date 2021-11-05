package config

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	HOST     = os.Getenv("GARAGE_API_DB_HOST")
	PORT     = os.Getenv("GARAGE_API_DB_PORT")
	NAME     = os.Getenv("GARAGE_API_DB_NAME")
	USER     = os.Getenv("GARAGE_API_DB_USER")
	PASSWORD = os.Getenv("GARAGE_API_DB_PASSWORD")
)

func GetReaderSqlx() *sql.DB {
	var DB_CONNECTION = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		HOST, USER, NAME, PASSWORD, PORT)

	db, err := sql.Open("postgres", DB_CONNECTION)
	if err != nil {
		panic(err.Error())
	}
	return db
}
