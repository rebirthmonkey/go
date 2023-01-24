package main

import (
	"bytes"
	"fmt"
)

// Taxi 出租车，享元对象，保存不变的内在属性信息
type Taxi struct {
	licensePlate string // 车牌
	color        string // 颜色
	brand        string // 汽车品牌
	company      string // 所属公司
}

// LocateFor 获取定位信息
func (t *Taxi) LocateFor(monitorMap string, x, y int) string {
	return fmt.Sprintf("%s,对于车牌号%s,%s,%s品牌,所属%s公司,定位(%d,%d)", monitorMap,
		t.licensePlate, t.color, t.brand, t.company, x, y)
}

// taxiFactoryInstance 出租车工厂单例
var taxiFactoryInstance = &TaxiFactory{
	taxis: make(map[string]*Taxi),
}

// GetTaxiFactory 获取出租车工厂单例
func GetTaxiFactory() *TaxiFactory {
	return taxiFactoryInstance
}

// TaxiFactory 出租车工厂类
type TaxiFactory struct {
	taxis map[string]*Taxi // key为车牌号
}

// getTaxi 获取出租车
func (f *TaxiFactory) getTaxi(licensePlate, color, brand, company string) *Taxi {
	if _, ok := f.taxis[licensePlate]; !ok {
		f.taxis[licensePlate] = &Taxi{
			licensePlate: licensePlate,
			color:        color,
			brand:        brand,
			company:      company,
		}
	}
	return f.taxis[licensePlate]
}

// TaxiPosition 出租车位置信息 x,y为外在数据信息，taxi为内在数据信息（享元对象）
type TaxiPosition struct {
	x    int
	y    int
	taxi *Taxi
}

func NewTaxiPosition(taxi *Taxi, x, y int) *TaxiPosition {
	return &TaxiPosition{
		taxi: taxi,
		x:    x,
		y:    y,
	}
}

// LocateFor 定位信息
func (p *TaxiPosition) LocateFor(monitorMap string) string {
	return p.taxi.LocateFor(monitorMap, p.x, p.y)
}

// TaxiDispatcher 出租车调度系统
type TaxiDispatcher struct {
	name   string
	traces map[string][]*TaxiPosition // 存储出租车当天轨迹信息，key为车牌号
}

func NewTaxiDispatcher(name string) *TaxiDispatcher {
	return &TaxiDispatcher{
		name:   name,
		traces: make(map[string][]*TaxiPosition),
	}
}

// AddTrace 添加轨迹
func (t *TaxiDispatcher) AddTrace(licensePlate, color, brand, company string, x, y int) {
	taxi := GetTaxiFactory().getTaxi(licensePlate, color, brand, company)
	t.traces[licensePlate] = append(t.traces[licensePlate], NewTaxiPosition(taxi, x, y))
}

// ShowTraces 显示轨迹
func (t *TaxiDispatcher) ShowTraces(licensePlate string) string {
	bytesBuf := bytes.Buffer{}
	for _, trace := range t.traces[licensePlate] {
		bytesBuf.WriteString(trace.LocateFor(t.name))
		bytesBuf.WriteByte('\n')
	}
	return bytesBuf.String()
}
