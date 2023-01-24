package main

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	dispatcher := NewTaxiDispatcher("北京市出租车调度系统")
	dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 10, 20)
	dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 20, 30)
	dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 30, 40)
	dispatcher.AddTrace("京B.123456", "黄色", "北京现代", "北汽", 40, 50)

	dispatcher.AddTrace("京B.567890", "红色", "一汽大众", "首汽", 20, 40)
	dispatcher.AddTrace("京B.567890", "红色", "一汽大众", "首汽", 50, 50)

	fmt.Println(dispatcher.ShowTraces("京B.123456"))
	fmt.Println(dispatcher.ShowTraces("京B.567890"))
}
