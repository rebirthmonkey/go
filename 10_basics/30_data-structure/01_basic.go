package main

import (
	"fmt"
	"math"
)

func main() {
	maxInt8 := math.MaxInt8
	minInt8 := math.MinInt8
	maxInt16 := math.MaxInt16
	minInt16 := math.MinInt16
	maxInt32 := math.MaxInt32
	minInt32 := math.MinInt32
	maxInt64 := math.MaxInt64
	minInt64 := math.MinInt64
	maxUint8 := math.MaxUint8
	maxUint16 := math.MaxUint16
	maxUint32 := math.MaxUint32
	//maxUint64 := math.MaxUint64
	maxFloat32 := math.MaxFloat32
	maxFloat64 := math.MaxFloat64
	fmt.Println("Range of Int8 :: ", minInt8, " to ", maxInt8)
	fmt.Println("Range of Int16 :: ", minInt16, " to ", maxInt16)
	fmt.Println("Range of Int32 :: ", minInt32, " to ", maxInt32)
	fmt.Println("Range of Int64 :: ", minInt64, " to ", maxInt64)
	fmt.Println("Max Uint8 :: ", maxUint8)
	fmt.Println("Max Uint16 :: ", maxUint16)
	fmt.Println("Max Uint32 :: ", maxUint32)
	//fmt.Println("Max Uint64 :: ", maxUint64)
	fmt.Println("Max Float32 :: ", maxFloat32)
	fmt.Println("Max Float64 :: ", maxFloat64)
}
