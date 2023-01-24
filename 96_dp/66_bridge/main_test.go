package main

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	// 坐飞机去三亚度蜜月
	honeymoonTravel := NewTravelExperience("honeymoon", new(airplane), NewSeaside("SanyaYalongBay"))
	fmt.Println(honeymoonTravel.Describe())
	// 坐车去泰山毕业旅游
	graduationTrip := NewTravelExperience("graduationTrip", new(car), NewMountain("Tarzan"))
	fmt.Println(graduationTrip.Describe())

	// 野外生存培训后，坐车去罗布泊，徒步穿越
	desertAdventure := NewAdventureExperience("wilderness survival training", "adventure", new(car), NewDesert("Lop Nor"))
	fmt.Println(desertAdventure.Describe())
}
