package controller

import (
	"encoding/json"
	"net/http"

	"github.com/trombettamoacyr/garage-api/entity"
	"github.com/trombettamoacyr/garage-api/error"
	"github.com/trombettamoacyr/garage-api/service"
)

type CarController interface {
	CreateCar(resp http.ResponseWriter, req *http.Request)
	GetCars(resp http.ResponseWriter, req *http.Request)
}

var (
	carService = service.NewCarService()
)

func CreateCar(resp http.ResponseWriter, req *http.Request) {
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

func GetCars(resp http.ResponseWriter, req *http.Request) {
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
