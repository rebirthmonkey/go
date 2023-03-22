package main

import (
	"fmt"
	"math"
)

func main() {
	squareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(squareRoot(9))
}
