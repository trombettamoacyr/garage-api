package service

import (
	"encoding/json"
	"net/http"

	"github.com/trombettamoacyr/garage-api/model"
	"github.com/trombettamoacyr/garage-api/util"
)

type ImageService interface {
	FetchData()
}

type fetchImageService struct{}

func NewImageService() ImageService {
	return &fetchImageService{}
}

const (
	imageApiUrl = "https://jsonplaceholder.typicode.com/photos/"
)

func (*fetchImageService) FetchData() {
	randomNum := util.NewRandomNumber(1, 2000)
	urlApi := imageApiUrl + randomNum

	client := http.Client{}
	resp, _ := client.Get(urlApi)

	var image = model.Image{}
	json.NewDecoder(resp.Body).Decode(&image)

	imageDataChannel <- image
}
