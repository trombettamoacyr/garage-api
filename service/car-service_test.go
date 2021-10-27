package service

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/trombettamoacyr/garage-api/entity"
)

var (
	testServiceNoRepo = NewCarService(nil)
)

type MockRepository struct {
	mock.Mock
}

func TestValidateSuccess(t *testing.T) {
	car := getCarMock()

	err := testServiceNoRepo.Validate(&car)

	assert.Nil(t, err)
}

func TestValidateNilCar(t *testing.T) {
	errMessageExpected := "The car is null."

	err := testServiceNoRepo.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func TestValidateModelCar(t *testing.T) {
	car := getCarMock()
	car.Model = ""
	errMessageExpected := "The model is empty."

	err := testServiceNoRepo.Validate(&car)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func TestValidateBrandCar(t *testing.T) {
	car := getCarMock()
	car.Brand = ""
	errMessageExpected := "The brand is empty."

	err := testServiceNoRepo.Validate(&car)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func TestValidateHpCar(t *testing.T) {
	car := getCarMock()
	car.Hp = 0
	errMessageExpected := "The horse power is empty."

	err := testServiceNoRepo.Validate(&car)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func TestValidateLicenseCar(t *testing.T) {
	car := getCarMock()
	car.License = ""
	errMessageExpected := "The license is empty."

	err := testServiceNoRepo.Validate(&car)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func TestValidateOwnerId(t *testing.T) {
	car := getCarMock()
	car.OwnerId = ""
	errMessageExpected := "The owner-id is empty."

	err := testServiceNoRepo.Validate(&car)

	assert.NotNil(t, err)
	assert.Equal(t, errMessageExpected, err.Error())
}

func (mock *MockRepository) FindById(id uuid.UUID) (*entity.Car, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Car), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Car, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Car), args.Error(1)
}

func TestFindAll(t *testing.T) {
	car := getCarMock()
	mockRepo := new(MockRepository)
	mockRepo.On("FindAll").Return([]entity.Car{car}, nil)

	carReturned, err := NewCarService(mockRepo).FindAll()

	assert.NotNil(t, carReturned)
	assert.Equal(t, car.Id, carReturned[0].Id)
	assert.Equal(t, car.Model, carReturned[0].Model)
	assert.Equal(t, car.Brand, carReturned[0].Brand)
	assert.Equal(t, car.Hp, carReturned[0].Hp)
	assert.Equal(t, car.License, carReturned[0].License)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func (mock *MockRepository) Save(car *entity.Car) (*entity.Car, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Car), args.Error(1)
}

func TestCreate(t *testing.T) {
	car := getCarMock()
	mockRepo := new(MockRepository)
	mockRepo.On("Save").Return(&car, nil)

	carReturned, err := NewCarService(mockRepo).Create(&car)

	assert.NotNil(t, carReturned)
	assert.Equal(t, car.Id, carReturned.Id)
	assert.Equal(t, car.Model, carReturned.Model)
	assert.Equal(t, car.Brand, carReturned.Brand)
	assert.Equal(t, car.Hp, carReturned.Hp)
	assert.Equal(t, car.License, carReturned.License)
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
