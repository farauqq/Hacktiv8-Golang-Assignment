package service

import (
	"math/rand"
)

func RandomNumberWater() int {
	num := rand.Intn(100-1) + 1
	return num
}

func RandomNumberWind() int {
	num := rand.Intn(100-1) + 1
	return num
}
