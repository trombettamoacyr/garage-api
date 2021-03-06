package entity

import "github.com/google/uuid"

type Car struct {
	Id             uuid.UUID `json:"id"`
	Model          string    `json:"model"`
	Brand          string    `json:"brand"`
	Hp             int       `json:"hp"`
	License        string    `json:"license"`
	InsuranceValue string    `json:"insurance_value"`
	OwnerId        string    `json:"owner_id"`
}
