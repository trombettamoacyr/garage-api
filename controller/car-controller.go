package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/trombettamoacyr/garage-api/entity"
	"github.com/trombettamoacyr/garage-api/error"
	"github.com/trombettamoacyr/garage-api/service"
)

type CarController interface {
	CreateCar(resp http.ResponseWriter, req *http.Request)
	GetCars(resp http.ResponseWriter, req *http.Request)
	GetCarById(resp http.ResponseWriter, req *http.Request)
	GetCarDetailById(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

func NewCarController(service service.CarService) CarController {
	carService = service
	return &controller{}
}

var (
	carService service.CarService
)

func (*controller) CreateCar(resp http.ResponseWriter, req *http.Request) {
	var car entity.Car
	err := json.NewDecoder(req.Body).Decode(&car)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the request"`))
		return
	}
	err1 := carService.Validate(&car)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(error.ServiceError{Message: err1.Error()})
		return
	}
	newCar, err2 := carService.Create(&car)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(error.ServiceError{Message: err2.Error()})
		return
	}
	result, _ := json.Marshal(newCar)
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusCreated)
	resp.Write(result)
}

func (*controller) GetCars(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	cars, err := carService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the cars"`))
		return
	}
	result, err := json.Marshal(cars)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the cars array"`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func (*controller) GetCarById(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	vars := mux.Vars(req)
	id, isPresent := vars["id"]
	if !isPresent {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"error": "Id is required"`))
		return
	}
	car, err := carService.FindById(id)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"error": "Car not found"`))
		return
	}
	result, err := json.Marshal(car)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the car entity"`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func (*controller) GetCarDetailById(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	vars := mux.Vars(req)
	id, isPresent := vars["id"]
	if !isPresent {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"error": "Id is required"`))
		return
	}
	car, err := carService.FindDetailById(id)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"error": "Car not found"`))
		return
	}
	result, err := json.Marshal(car)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the car entity"`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}
