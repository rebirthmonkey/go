package main

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	xierqiStation := NewSubwayStation("西二旗")
	fmt.Println(EnhanceEnterStationProcess(xierqiStation, false, false).Enter())
	fmt.Println(EnhanceEnterStationProcess(xierqiStation, true, false).Enter())
	fmt.Println(EnhanceEnterStationProcess(xierqiStation, true, true).Enter())
}

// EnhanceEnterStationProcess 根据是否有行李，是否处于疫情，增加进站流程
func EnhanceEnterStationProcess(station Station, hasLuggage bool, hasEpidemic bool) Station {
	if hasLuggage {
		station = NewSecurityCheckDecorator(station)
	}
	if hasEpidemic {
		station = NewEpidemicProtectionDecorator(station)
	}
	return station
}
