package main

import (
	"github.com/trombettamoacyr/garage-api/entity"
	"github.com/trombettamoacyr/garage-api/repository"
	"log"

	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

var (
	repo = repository.NewCarRepository()
)

func getCars(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	cars, err := repo.FindAll()
	handleError(err)
	result, err := json.Marshal(cars)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the cars array"`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func createCar(resp http.ResponseWriter, req *http.Request) {
	var car entity.Car
	err := json.NewDecoder(req.Body).Decode(&car)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the request"`))
		return
	}
	car.Id = generateUUID()
	newCar, err := repo.Save(&car)
	handleError(err)

	result, _ := json.Marshal(newCar)
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusCreated)
	resp.Write(result)
}

func generateUUID() uuid.UUID {
	newUUID, err := uuid.NewRandom()
	handleError(err)
	return newUUID
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("Failed to iterate the list of cars: %v", err)
	}
}
