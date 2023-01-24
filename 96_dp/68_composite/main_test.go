package main

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	gusu := NewTown("姑苏区", 100, 2000.00)
	fmt.Println(ShowRegionInfo(gusu))
	wuzhong := NewTown("吴中区", 150, 2600.00)
	fmt.Println(ShowRegionInfo(wuzhong))
	huqiu := NewTown("虎丘区", 80, 1800.00)
	fmt.Println(ShowRegionInfo(huqiu))

	kunshan := NewCities("昆山市")
	kunshan.Add(NewTown("玉山镇", 60, 1200.00),
		NewTown("周庄镇", 68, 1900.00),
		NewTown("花桥镇", 78, 2200.00))
	fmt.Println(ShowRegionInfo(kunshan))

	changshu := NewCities("常熟市")
	changshu.Add(NewTown("沙家浜镇", 55, 1100.00),
		NewTown("古里镇", 59, 1300.00),
		NewTown("辛庄镇", 68, 2100.00))
	fmt.Println(ShowRegionInfo(changshu))

	suzhou := NewCities("苏州市")
	suzhou.Add(gusu, wuzhong, huqiu, kunshan, changshu)
	fmt.Println(ShowRegionInfo(suzhou))

}

func ShowRegionInfo(region Region) string {
	return fmt.Sprintf("%s, 人口:%d万, GDP:%.2f亿", region.Name(), region.Population(), region.GDP())
}
