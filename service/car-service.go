package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/trombettamoacyr/garage-api/model"

	"github.com/trombettamoacyr/garage-api/entity"
	"github.com/trombettamoacyr/garage-api/repository"
)

type CarService interface {
	FindAll() ([]entity.Car, error)
	FindById(id string) (*entity.Car, error)
	FindDetailById(id string) (*model.CarDetail, error)
	Validate(car *entity.Car) error
	Create(car *entity.Car) (*entity.Car, error)
}

type service struct{}

func NewCarService(carDetailService CarDetail, carInsuranceService CarInsurance, repository repository.CarRepository) CarService {
	repo = repository
	detailService = carDetailService
	insuranceService = carInsuranceService
	return &service{}
}

var (
	repo             repository.CarRepository
	detailService    CarDetail
	insuranceService CarInsurance
)

func (*service) FindAll() ([]entity.Car, error) {
	return repo.FindAll()
}

func (*service) FindById(id string) (*entity.Car, error) {
	car, err := repo.FindById(uuid.MustParse(id))
	if err != nil {
		err := errors.New("Car not found.")
		return nil, err
	}
	return car, nil
}

func (*service) FindDetailById(id string) (*model.CarDetail, error) {
	car, err := repo.FindById(uuid.MustParse(id))
	if err != nil {
		err := errors.New("Car not found.")
		return nil, err
	}
	ownerId := car.OwnerId
	details := detailService.FetchCarData(ownerId)

	var carDetail model.CarDetail
	carDetail.Car = *car
	carDetail.Owner = details.Owner
	carDetail.Image = details.Image

	return &carDetail, nil
}

func (*service) Validate(car *entity.Car) error {
	if car == nil {
		err := errors.New("The car is null.")
		return err
	}
	if car.Model == "" {
		err := errors.New("The model is empty.")
		return err
	}
	if car.Brand == "" {
		err := errors.New("The brand is empty.")
		return err
	}
	if car.Hp == 0 {
		err := errors.New("The horse power is empty.")
		return err
	}
	if car.License == "" {
		err := errors.New("The license is empty.")
		return err
	}
	if car.OwnerId == "" {
		err := errors.New("The owner-id is empty.")
		return err
	}
	return nil
}

func (*service) Create(car *entity.Car) (*entity.Car, error) {
	newUUID, _ := uuid.NewRandom()
	insuranceValue := insuranceService.FetchValue()

	car.Id = newUUID
	car.InsuranceValue = insuranceValue

	return repo.Save(car)
}
