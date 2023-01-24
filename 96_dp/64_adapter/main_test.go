package main

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	huaweiMate40Pro := NewHuaweiPhone("华为 mate40 pro")
	iphone13MaxPro := NewIPhone("苹果 iphone13 pro max")

	powerBank := &PowerBank{"飞利浦"}
	fmt.Println(powerBank.Charge(NewHuaweiPhonePlugAdapter(huaweiMate40Pro)))
	fmt.Println(powerBank.Charge(NewApplePhonePlugAdapter(iphone13MaxPro)))
}
