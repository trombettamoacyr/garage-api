package entity

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"first_name"`
	LastName string    `json:"last_name"`
	Email    string    `json:"email"`
}
