package model

import (
	"github.com/trombettamoacyr/garage-api/entity"
)

type CarDetail struct {
	Car   entity.Car
	Owner Owner
	Image Image
}
