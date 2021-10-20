package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Car struct {
	Id      uuid.UUID `json:"id"`
	Model   string    `json:"model"`
	Brand   string    `json:"brand"`
	Hp      int       `json:"hp"`
	License string    `json:"license"`
}

var (
	cars []Car
)

func init() {
	cars = []Car{
		Car{
			Id: generateUUID(), Model: "UP", Brand: "Volkswagen", Hp: 100, License: "ABC1020",
		},
	}
}

func getCars(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
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
	var car Car
	err := json.NewDecoder(req.Body).Decode(&car)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the request"`))
		return
	}
	car.Id = generateUUID()
	cars = append(cars, car)
	result, _ := json.Marshal(car)
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
		panic(err.Error())
	}
}
