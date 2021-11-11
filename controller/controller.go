package controller

import "net/http"

type CarController interface {

	// CreateCar
	// @Summary Create a car
	// @Description Create a new car to the garage
	// @Tags Cars
	// @Accept json
	// @Param car body entity.Car true "Car data"
	// @Success 200 {object} entity.Car
	// @Router / [post]
	CreateCar(resp http.ResponseWriter, req *http.Request)

	// GetCars
	// @Summary Get all cars
	// @Description Get all cars from garage
	// @Tags Cars
	// @Success 200 {array} entity.Car
	// @Failure 404 {object} object
	// @Router / [get]
	GetCars(resp http.ResponseWriter, req *http.Request)

	// GetCarById
	// @Summary Get a car information
	// @Description Get a car and insurance value
	// @Tags Cars
	// @Param id path string true "Car id"
	// @Success 200 {object} entity.Car
	// @Failure 404 {object} object
	// @Router / [get]
	GetCarById(resp http.ResponseWriter, req *http.Request)

	// GetCarDetailById
	// @Summary Get a car details
	// @Description Get a car, owner information and pictures
	// @Tags Cars
	// @Param id path string true "Car id"
	// @Success 200 {object} model.CarDetail
	// @Failure 404 {object} object
	// @Router / [get]
	GetCarDetailById(resp http.ResponseWriter, req *http.Request)
}
