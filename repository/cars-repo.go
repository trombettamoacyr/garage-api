package repository

import (
	"github.com/google/uuid"
	"github.com/trombettamoacyr/garage-api/entity"
)

type CarRepository interface {
	Save(car *entity.Car) (*entity.Car, error)
	FindAll() (*[]entity.Car, error)
	FindById(id uuid.UUID) (*entity.Car, error)
}
