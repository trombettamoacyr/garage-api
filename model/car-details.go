package model

import (
	"github.com/trombettamoacyr/garage-api/entity"
)

type CarDetails struct {
	Car   entity.Car
	Owner Owner
	Image Image
}
