package service

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/trombettamoacyr/garage-api/model"
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

type MockDetailService struct {
	mock.Mock
}

func (mock *MockRepository) FindAll() (*[]entity.Car, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*[]entity.Car), args.Error(1)
}

func (mock *MockRepository) FindById(id uuid.UUID) (*entity.Car, error) {
	args := mock.Called(id)
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

func (mock *MockDetailService) FetchCarData(ownerId string) *model.CarDetail {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.CarDetail)
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

func TestFindById(t *testing.T) {
	car := getCarMock()
	carId := car.Id
	carIdString := carId.String()

	mockRepo := new(MockRepository)
	mockRepo.On("FindById", carId).Return(&car, nil)

	carPointer, err := NewCarService(nil, nil, mockRepo).FindById(carIdString)
	carReturned := *carPointer

	assert.NotNil(t, carReturned)
	assert.Equal(t, car.Id, carReturned.Id)
	assert.Equal(t, car.Model, carReturned.Model)
	assert.Equal(t, car.Brand, carReturned.Brand)
	assert.Equal(t, car.Hp, carReturned.Hp)
	assert.Equal(t, car.License, carReturned.License)
	assert.Equal(t, car.OwnerId, carReturned.OwnerId)
	assert.Equal(t, car.InsuranceValue, carReturned.InsuranceValue)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

//func TestFindDetailById(t *testing.T) {
//	car := getCarMock()
//	carDetail := getCarDetailsMock()
//	carId := car.Id.String()
//
//	mockRepo := new(MockRepository)
//	mockRepo.On("FindById", carId).Return(&car, nil)
//
//	mockDetail := new(MockDetailService)
//	mockDetail.On("FetchCarData", car.OwnerId).Return(carDetail)
//
//	carPointer, err := NewCarService(mockDetail, nil, mockRepo).FindDetailById(carId)
//	carReturned := *carPointer
//
//	assert.NotNil(t, carReturned)
//	assert.Equal(t, car.Id, carReturned.Car.Id)
//	assert.Equal(t, car.Model, carReturned.Car.Model)
//	assert.Equal(t, car.Brand, carReturned.Car.Brand)
//	assert.Equal(t, car.Hp, carReturned.Car.Hp)
//	assert.Equal(t, car.License, carReturned.Car.License)
//	assert.Equal(t, car.OwnerId, carReturned.Car.OwnerId)
//	assert.Equal(t, carDetail.Owner.Owner.FirstName, carReturned.Owner.Owner.FirstName)
//	assert.Equal(t, carDetail.Owner.Owner.LastName, carReturned.Owner.Owner.LastName)
//	assert.Equal(t, carDetail.Owner.Owner.Email, carReturned.Owner.Owner.Email)
//	assert.Equal(t, carDetail.Owner.Owner.Phone, carReturned.Owner.Owner.Phone)
//	assert.Equal(t, carDetail.Image.Url, carReturned.Image.Url)
//	assert.Equal(t, carDetail.Image.ThumbnailUrl, carReturned.Image.ThumbnailUrl)
//	assert.Nil(t, err)
//
//	mockRepo.AssertExpectations(t)
//	mockDetail.AssertExpectations(t)
//}

func getCarMock() entity.Car {
	carId, _ := uuid.NewRandom()
	return entity.Car{
		Id:      carId,
		Model:   "Model",
		Brand:   "Brand",
		Hp:      999,
		License: "License",
		OwnerId: "ownerId",
		InsuranceValue: "R$999",
	}
}

func getCarDetailsMock() model.CarDetail {
	return model.CarDetail{
		Owner: getOwnerMock(),
		Image: model.Image{
			Url:          "www.test.com/image.jpg",
			ThumbnailUrl: "www.test.com/thumb/image.jpg",
		},
	}
}

func getOwnerMock() model.Owner {
	return model.Owner{
		Owner: model.OwnerData{
			FirstName: "Robert",
			LastName:  "Martin",
			Email:     "uncle@bob.com.br",
			Phone:     "112233",
		},
	}
}
