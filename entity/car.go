package entity

import "github.com/google/uuid"

type Car struct {
	Id             uuid.UUID `json:"id"`
	Model          string    `json:"model"`
	Brand          string    `json:"brand"`
	Hp             int       `json:"hp"`
	License        string    `json:"license"`
	InsurancePrice int64     `json:"insurance_price"`
	OwnerId        string    `json:"owner_id"`
}
