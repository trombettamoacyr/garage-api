package util

import (
	"math/rand"
	"strconv"
)

func NewRandomNumber(min int, max int) string {
	num := rand.Intn(max-min) + min
	return strconv.Itoa(num)
}
