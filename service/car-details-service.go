package service
//
//import (
//	"fmt"
//	"net/http"
//)
//
//type CarDetails interface {
//	FetchCarData(brand string) CarDetails
//	//FetchOwnerData(id string)
//}
//
//type fetchCarDateService struct{}
//
//func NewFetchCarDateService() CarDetails {
//	return &fetchCarDateService{}
//}
//
//var (
//	ownerService     = NewOwnerService()
//	imageService     = NewImageService()
//	ownerDataChannel = make(chan *http.Response)
//	imageDataChannel = make(chan *http.Response)
//)
//
//func (*fetchCarDateService) FetchCarData(brand string) CarDetails {
//
//
//	pathSegment := "name/"
//	urlApi := carDetailsUrl + pathSegment + brand
//
//	client := http.Client{}
//	resp, _ := client.Get(urlApi)
//	fmt.Println(resp.Body)
//}
