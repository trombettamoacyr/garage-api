package repository

import (
	"github.com/google/uuid"

	"github.com/trombettamoacyr/garage-api/config"
	"github.com/trombettamoacyr/garage-api/entity"
)

type postgresRepo struct{}

func NewPostgresRepo() CarRepository {
	return &postgresRepo{}
}

const (
	QUERY_SAVE       = "insert into car (id, model, brand, hp, license, insurance_value, owner_id) values ($1, $2, $3, $4, $5, $6, $7)"
	QUERY_FIND_ALL   = "select * from car"
	QUERY_FIND_BY_ID = "select * from car where id = $1"
)

func (*postgresRepo) Save(car *entity.Car) (*entity.Car, error) {
	dbConnection := config.GetConnection()
	_, err := dbConnection.Exec(QUERY_SAVE, car.Id, car.Model, car.Brand, car.Hp, car.License, car.InsuranceValue, car.OwnerId)
	if err != nil {
		return nil, err
	}
	defer dbConnection.Close()
	return car, nil
}

func (*postgresRepo) FindAll() (*[]entity.Car, error) {
	dbConnection := config.GetConnection()
	query, err := dbConnection.Query(QUERY_FIND_ALL)
	if err != nil {
		return nil, err
	}
	var cars []entity.Car
	var car entity.Car
	{
	}

	for query.Next() {
		var carId uuid.UUID
		var model string
		var brand string
		var hp int
		var license string
		var insuranceValue string
		var ownerId string

		err = query.Scan(&carId, &model, &brand, &hp, &license, &insuranceValue, &ownerId)
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
	defer dbConnection.Close()
	return &cars, nil
}

func (*postgresRepo) FindById(id uuid.UUID) (*entity.Car, error) {
	dbConnection := config.GetConnection()
	query, err := dbConnection.Query(QUERY_FIND_BY_ID, id.String())
	if err != nil {
		return nil, err
	}
	car := entity.Car{}
	for query.Next() {
		var carId uuid.UUID
		var model string
		var brand string
		var hp int
		var license string
		var insuranceValue string
		var ownerId string

		err = query.Scan(&carId, &model, &brand, &hp, &license, &insuranceValue, &ownerId)
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
	defer dbConnection.Close()
	return &car, nil
}
