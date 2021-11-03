package service

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/trombettamoacyr/garage-api/entity"
)

var (
	testServiceNoRepo = NewCarService(nil, nil, nil)
)

type MockRepository struct {
	mock.Mock
}

type MockInsuranceService struct {
	mock.Mock
}

func TestValidateSuccess(t *testing.T) {
	car := getCarMock()

	err := testServiceNoRepo.Validate(&car)

	assert.Nil(t, err)
}

func TestValidateNilCar(t *testing.T) {
	errMessageExpected := "car is null"

	err := testServiceNoRepo.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func TestValidateModelCar(t *testing.T) {
	car := getCarMock()
	car.Model = ""
	errMessageExpected := "model is empty"

	err := testServiceNoRepo.Validate(&car)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func TestValidateBrandCar(t *testing.T) {
	car := getCarMock()
	car.Brand = ""
	errMessageExpected := "brand is empty"

	err := testServiceNoRepo.Validate(&car)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func TestValidateHpCar(t *testing.T) {
	car := getCarMock()
	car.Hp = 0
	errMessageExpected := "horse power is empty"

	err := testServiceNoRepo.Validate(&car)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func TestValidateLicenseCar(t *testing.T) {
	car := getCarMock()
	car.License = ""
	errMessageExpected := "license is empty"

	err := testServiceNoRepo.Validate(&car)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func TestValidateOwnerId(t *testing.T) {
	car := getCarMock()
	car.OwnerId = ""
	errMessageExpected := "owner id is empty"

	err := testServiceNoRepo.Validate(&car)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func (mock *MockRepository) FindAll() (*[]entity.Car, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*[]entity.Car), args.Error(1)
}

func (mock *MockRepository) FindById(id uuid.UUID) (*entity.Car, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Car), args.Error(1)
}

func (mock *MockRepository) Save(car *entity.Car) (*entity.Car, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Car), args.Error(1)
}

func (mock *MockInsuranceService) FetchValue() string {
	args := mock.Called()
	result := args.Get(0)
	return result.(string)
}

func TestFindAll(t *testing.T) {
	car := getCarMock()
	var cars []entity.Car
	cars = append(cars, car)

	mockRepo := new(MockRepository)
	mockRepo.On("FindAll").Return(&cars, nil)

	carsPointer, err := NewCarService(nil, nil, mockRepo).FindAll()
	cars = *carsPointer

	assert.NotNil(t, cars)
	assert.Equal(t, car.Id, cars[0].Id)
	assert.Equal(t, car.Model, cars[0].Model)
	assert.Equal(t, car.Brand, cars[0].Brand)
	assert.Equal(t, car.Hp, cars[0].Hp)
	assert.Equal(t, car.License, cars[0].License)
	assert.Equal(t, car.OwnerId, cars[0].OwnerId)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	car := getCarMock()
	insuranceValue := "$9999.00"

	mockInsurance := new(MockInsuranceService)
	mockInsurance.On("FetchValue").Return(insuranceValue)

	mockRepo := new(MockRepository)
	mockRepo.On("Save").Return(&car, nil)

	carReturned, err := NewCarService(nil, mockInsurance, mockRepo).Create(&car)

	assert.NotNil(t, carReturned)
	assert.Equal(t, car.Id, carReturned.Id)
	assert.Equal(t, car.Model, carReturned.Model)
	assert.Equal(t, car.Brand, carReturned.Brand)
	assert.Equal(t, car.Hp, carReturned.Hp)
	assert.Equal(t, car.License, carReturned.License)
	assert.Equal(t, car.OwnerId, carReturned.OwnerId)
	assert.Equal(t, insuranceValue, carReturned.InsuranceValue)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func getCarMock() entity.Car {
	carId, _ := uuid.NewRandom()
	return entity.Car{
		Id:      carId,
		Model:   "Model",
		Brand:   "Brand",
		Hp:      999,
		License: "License",
		OwnerId: "id",
	}
}
