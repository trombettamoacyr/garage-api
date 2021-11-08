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
	QUERY_FIND_BY_ID = "select * from car where id = $1"
)

func (*postgresRepo) Save(car *entity.Car) (*entity.Car, error) {
	return nil, nil
}

func (*postgresRepo) FindAll() (*[]entity.Car, error) {
	return nil, nil
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
