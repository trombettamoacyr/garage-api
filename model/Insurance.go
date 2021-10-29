package model

type Insurance struct {
	Car CarInsurance `json:"Car"`
}

type CarInsurance struct {
	InsuranceValue string `json:"price"`
}
