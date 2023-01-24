package main

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	iPhone13Pro := NewIPhone("13 pro") // 刚创建的手机有部分电

	fmt.Println(iPhone13Pro.BatteryState()) // 打印部分电状态
	fmt.Println(iPhone13Pro.ConnectPlug())  // 插上电源插头，继续充满电
	fmt.Println(iPhone13Pro.ConnectPlug())  // 满电后再充电，会触发满电保护

	fmt.Println(iPhone13Pro.DisconnectPlug()) // 拔掉电源，使用手机消耗电量，变为有部分电
	fmt.Println(iPhone13Pro.DisconnectPlug()) // 一直使用手机，直到没电
	fmt.Println(iPhone13Pro.DisconnectPlug()) // 没电后会关机

	fmt.Println(iPhone13Pro.ConnectPlug()) // 再次插上电源一会，变为有电状态
}
