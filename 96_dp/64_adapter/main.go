package main

import "fmt"

// HuaweiPlug 华为手机充电插槽接口
type HuaweiPlug interface {
	ConnectTypeC() string
}

// HuaweiPhone 华为系列手机
type HuaweiPhone struct {
	model string
}

// NewHuaweiPhone 华为手机创建方法
func NewHuaweiPhone(model string) *HuaweiPhone {
	return &HuaweiPhone{
		model: model,
	}
}

// ConnectTypeC 华为手机TypeC充电插槽
func (h *HuaweiPhone) ConnectTypeC() string {
	return fmt.Sprintf("%v connect typeC plug", h.model)
}

// ApplePlug 苹果手机充电插槽
type ApplePlug interface {
	ConnectLightning() string
}

// IPhone 苹果系列手机
type IPhone struct {
	model string
}

// NewIPhone 苹果手机创建方法
func NewIPhone(model string) *IPhone {
	return &IPhone{
		model: model,
	}
}

// ConnectLightning 苹果手机Lightning充电插槽
func (i *IPhone) ConnectLightning() string {
	return fmt.Sprintf("%v connect lightning plug", i.model)
}

// CommonPlug 通用的USB电源插槽
type CommonPlug interface {
	ConnectUSB() string
}

// HuaweiPhonePlugAdapter 华为TypeC充电插槽适配通用USB充电插槽
type HuaweiPhonePlugAdapter struct {
	huaweiPhone HuaweiPlug
}

// NewHuaweiPhonePlugAdapter 创建华为手机适配USB充电插槽适配器
func NewHuaweiPhonePlugAdapter(huaweiPhone HuaweiPlug) *HuaweiPhonePlugAdapter {
	return &HuaweiPhonePlugAdapter{
		huaweiPhone: huaweiPhone,
	}
}

// ConnectUSB 链接USB
func (h *HuaweiPhonePlugAdapter) ConnectUSB() string {
	return fmt.Sprintf("%v adapt to usb ", h.huaweiPhone.ConnectTypeC())
}

// ApplePhonePlugAdapter 苹果Lightning充电插槽适配通用USB充电插槽
type ApplePhonePlugAdapter struct {
	iPhone ApplePlug
}

// NewApplePhonePlugAdapter 创建苹果手机适配USB充电插槽适配器
func NewApplePhonePlugAdapter(iPhone ApplePlug) *ApplePhonePlugAdapter {
	return &ApplePhonePlugAdapter{
		iPhone: iPhone,
	}
}

// ConnectUSB 链接USB
func (a *ApplePhonePlugAdapter) ConnectUSB() string {
	return fmt.Sprintf("%v adapt to usb ", a.iPhone.ConnectLightning())
}

// PowerBank 充电宝
type PowerBank struct {
	brand string
}

// Charge 支持通用USB接口充电
func (p *PowerBank) Charge(plug CommonPlug) string {
	return fmt.Sprintf("%v power bank connect usb plug, start charge for %v", p.brand, plug.ConnectUSB())
}
