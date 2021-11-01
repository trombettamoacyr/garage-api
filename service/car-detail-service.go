package service

import (
	"github.com/trombettamoacyr/garage-api/model"
)

type CarDetails interface {
	FetchCarData(id string) model.CarDetails
}

type carDetailsService struct{}

func NewCarDetailsService() CarDetails {
	return &carDetailsService{}
}

var (
	ownerService     = NewOwnerService()
	imageService     = NewImageService()
	ownerDataChannel = make(chan model.Owner)
	imageDataChannel = make(chan model.Image)
)

func (*carDetailsService) FetchCarData(id string) model.CarDetails {
	go ownerService.FetchDate(id)
	go imageService.FetchData()

	var car = model.CarDetails{}
	car.Owner = getOwnerData()
	car.Image = getImageData()

	return car
}

func getOwnerData() model.Owner {
	owner := <-ownerDataChannel
	return owner
}

func getImageData() model.Image {
	image := <-imageDataChannel
	return image
}
