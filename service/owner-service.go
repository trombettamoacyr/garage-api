package service

import (
	"encoding/json"
	"github.com/trombettamoacyr/garage-api/model"
	"net/http"
)

type OwnerService interface {
	FetchDate(id string)
}

type fetchOwnerDataService struct{}

func NewOwnerService() OwnerService {
	return &fetchOwnerDataService{}
}

const (
	ownerApiUrl = "https://myfakeapi.com/api/users/"
)

func (*fetchOwnerDataService) FetchDate(id string) {
	url := ownerApiUrl + id
	client := http.Client{}
	resp, _ := client.Get(url)

	var owner = model.Owner{}
	json.NewDecoder(resp.Body).Decode(&owner)

	ownerDataChannel <- owner
}
