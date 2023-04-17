package test

import (
	"math"
	"math/rand"
)

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func RandInt() int {
	return rand.Int()
}
