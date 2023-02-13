package main

import (
	"strconv"
)

func main() {
	// Int to String
	var i1 int = 5
	var s1 string
	s1 = strconv.Itoa(i1)
	println("the string is: ", s1)

	// String to Int
	var s2 string = "5"
	var i2 int
	i2, _ = strconv.Atoi(s2)
	println("the int is: ", i2)

	// String to Float
	var s3 string = "5"
	var f float64
	f, _ = strconv.ParseFloat(s3, 32)
	println("the float is: ", f)
}
