package service

import (
	"fmt"
	"net/http"
)

type CarDetails interface {
	FetchCarData(brand string)
	//FetchOwnerData(id string)
}

type fetchCarDateService struct {}

func NewFetchCarDateService() CarDetails {
	return &fetchCarDateService{}
}

const (
	carDetailsUrl = "https://myfakeapi.com/api/cars/"
)

func (*fetchCarDateService) FetchCarData(brand string) {
	pathSegment := "name/"
	urlApi := carDetailsUrl + pathSegment + brand

	client := http.Client{}
	resp, _ := client.Get(urlApi)
	fmt.Println(resp.Body)
}