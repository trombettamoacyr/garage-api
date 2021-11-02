package service

import (
	"github.com/trombettamoacyr/garage-api/model"
)

type CarDetail interface {
	FetchCarData(id string) *model.CarDetail
}

type carDetailService struct{}

func NewCarDetailService() CarDetail {
	return &carDetailService{}
}

var (
	ownerService     = NewOwnerService()
	imageService     = NewImageService()
	ownerDataChannel = make(chan model.Owner)
	imageDataChannel = make(chan model.Image)
)

func (*carDetailService) FetchCarData(id string) *model.CarDetail {
	go ownerService.FetchDate(id)
	go imageService.FetchData()

	var car = model.CarDetail{}
	car.Owner = getOwnerData()
	car.Image = getImageData()

	return &car
}

func getOwnerData() model.Owner {
	owner := <-ownerDataChannel
	return owner
}

func getImageData() model.Image {
	image := <-imageDataChannel
	return image
}
