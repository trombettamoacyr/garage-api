package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"

	"github.com/trombettamoacyr/garage-api/entity"
)

type sqliteRepo struct{}

func NewSqliteRepo() CarRepository {
	os.Remove("./car.db")

	db, err := sql.Open("sqlite3", "./car.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE car (id uuid NOT NULL, model varchar NOT NULL, brand varchar NOT NULL, hp int NOT NULL, license varchar NOT NULL, insurance_value varchar NULL, owner_id varchar NULL);
	DELETE from car;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	defer db.Close()
	return &sqliteRepo{}
}

const (
	QUERY_SAVE_SQLITE       = "insert into car (id, model, brand, hp, license, insurance_value, owner_id) values (?, ?, ?, ?, ?, ?, ?)"
	QUERY_FIND_ALL_SQLITE   = "select * from car"
	QUERY_FIND_BY_ID_SQLITE = "select * from car where id = $1"
)

func (*sqliteRepo) Save(car *entity.Car) (*entity.Car, error) {
	db, err := sql.Open("sqlite3", "./car.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stmt, err := tx.Prepare(QUERY_SAVE_SQLITE)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	_, err = stmt.Exec(car.Id, car.Model, car.Brand, car.Hp, car.License, car.InsuranceValue, car.OwnerId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tx.Commit()

	defer stmt.Close()
	return car, nil
}

func (*sqliteRepo) FindAll() (*[]entity.Car, error) {
	db, err := sql.Open("sqlite3", "./car.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := db.Query(QUERY_FIND_ALL_SQLITE)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var cars []entity.Car
	var car entity.Car

	for rows.Next() {
		var carId uuid.UUID
		var model string
		var brand string
		var hp int
		var license string
		var insuranceValue string
		var ownerId string

		err = rows.Scan(&carId, &model, &brand, &hp, &license, &insuranceValue, &ownerId)
		if err != nil {
			return nil, err
		}
		car.Id = carId
		car.Model = model
		car.Brand = brand
		car.Hp = hp
		car.License = license
		car.InsuranceValue = insuranceValue
		car.OwnerId = ownerId

		cars = append(cars, car)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()
	return &cars, nil
}

func (*sqliteRepo) FindById(id uuid.UUID) (*entity.Car, error) {
	db, err := sql.Open("sqlite3", "./car.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	row, err := db.Query(QUERY_FIND_BY_ID_SQLITE, id.String())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	car := entity.Car{}
	for row.Next() {
		var carId uuid.UUID
		var model string
		var brand string
		var hp int
		var license string
		var insuranceValue string
		var ownerId string

		err = row.Scan(&carId, &model, &brand, &hp, &license, &insuranceValue, &ownerId)
		if err != nil {
			return nil, err
		}
		car.Id = carId
		car.Model = model
		car.Brand = brand
		car.Hp = hp
		car.License = license
		car.InsuranceValue = insuranceValue
		car.OwnerId = ownerId
	}
	defer row.Close()
	return &car, nil
}
