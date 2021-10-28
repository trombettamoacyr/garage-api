package model

import "github.com/trombettamoacyr/garage-api/entity"

type CarDetails struct {
	car   entity.Car
	owner struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}
}
