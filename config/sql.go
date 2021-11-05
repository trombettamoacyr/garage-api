package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

const (
	CAR_TABLE = "CREATE TABLE public.car (id uuid NOT NULL, model varchar NOT NULL, brand varchar NOT NULL, " +
		"hp int NOT NULL, license varchar NOT NULL, insurance_value varchar NULL, owner_id varchar NULL);"

	CAR_TABLE_INDEX = "CREATE INDEX IF NOT EXISTS car_id_idx ON public.car (id);"
)

var (
	HOST     = os.Getenv("GARAGE_API_DB_HOST")
	PORT     = os.Getenv("GARAGE_API_DB_PORT")
	NAME     = os.Getenv("GARAGE_API_DB_NAME")
	USER     = os.Getenv("GARAGE_API_DB_USER")
	PASSWORD = os.Getenv("GARAGE_API_DB_PASSWORD")
)

func GetConnection() *sql.DB {
	var DB_CONNECTION = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		HOST, USER, NAME, PASSWORD, PORT)

	db, err := sql.Open("postgres", DB_CONNECTION)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec(CAR_TABLE)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(CAR_TABLE_INDEX)
	if err != nil {
		panic(err.Error())
	}

	return db
}

