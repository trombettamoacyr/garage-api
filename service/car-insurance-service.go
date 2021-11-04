package service

import (
	"encoding/json"
	"net/http"

	"github.com/trombettamoacyr/garage-api/model"
	"github.com/trombettamoacyr/garage-api/util"
)

type CarInsurance interface {
	FetchValue() string
}

type carInsuranceService struct{}

func NewCarInsuranceService() CarInsurance {
	return &carInsuranceService{}
}

const (
	carDetailUrl = "https://myfakeapi.com/api/cars/"
)

func (*carInsuranceService) FetchValue() string {
	randomNum := util.NewRandomNumber(1, 999)
	urlApi := carDetailUrl + randomNum

	client := http.Client{}
	resp, _ := client.Get(urlApi)

	var insurance = model.Insurance{}
	json.NewDecoder(resp.Body).Decode(&insurance)

	return insurance.Car.InsuranceValue
}
