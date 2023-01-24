package main

import "fmt"

// Station 车站，修饰器模式统一接口
type Station interface {
	Enter() string // 进站
}

// subwayStation 地铁站
type subwayStation struct {
	name string
}

// NewSubwayStation 创建指定站名地铁站
func NewSubwayStation(name string) *subwayStation {
	return &subwayStation{
		name: name,
	}
}

// Enter 进地铁站
func (s *subwayStation) Enter() string {
	return fmt.Sprintf("买票进入%s地铁站。", s.name)
}

// securityCheckDecorator 进站安检修饰器
type securityCheckDecorator struct {
	station Station
}

func NewSecurityCheckDecorator(station Station) *securityCheckDecorator {
	return &securityCheckDecorator{
		station: station,
	}
}

func (s *securityCheckDecorator) Enter() string {
	return "行李通过安检；" + s.station.Enter()
}

// epidemicProtectionDecorator 进站疫情防护修饰器
type epidemicProtectionDecorator struct {
	station Station
}

func NewEpidemicProtectionDecorator(station Station) *epidemicProtectionDecorator {
	return &epidemicProtectionDecorator{
		station: station,
	}
}

func (e *epidemicProtectionDecorator) Enter() string {
	return "测量体温，佩戴口罩；" + e.station.Enter()
}
