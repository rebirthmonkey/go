package main

import "fmt"

// ElectricCooker 电饭煲
type ElectricCooker struct {
	fire     string // 火力
	pressure string // 压力
}

// SetFire 设置火力
func (e *ElectricCooker) SetFire(fire string) {
	e.fire = fire
}

// SetPressure 设置压力
func (e *ElectricCooker) SetPressure(pressure string) {
	e.pressure = pressure
}

// Run 持续运行指定时间
func (e *ElectricCooker) Run(duration string) string {
	return fmt.Sprintf("电饭煲设置火力为%s,压力为%s,持续运行%s;", e.fire, e.pressure, duration)
}

// Shutdown 停止
func (e *ElectricCooker) Shutdown() string {
	return "电饭煲停止运行。"
}

// CookCommand 做饭指令接口
type CookCommand interface {
	Execute() string // 指令执行方法
}

// steamRiceCommand 蒸饭指令
type steamRiceCommand struct {
	electricCooker *ElectricCooker // 电饭煲
}

func NewSteamRiceCommand(electricCooker *ElectricCooker) *steamRiceCommand {
	return &steamRiceCommand{
		electricCooker: electricCooker,
	}
}

func (s *steamRiceCommand) Execute() string {
	s.electricCooker.SetFire("中")
	s.electricCooker.SetPressure("正常")
	return "蒸饭:" + s.electricCooker.Run("30分钟")
}

// cookCongeeCommand 煮粥指令
type cookCongeeCommand struct {
	electricCooker *ElectricCooker
}

func NewCookCongeeCommand(electricCooker *ElectricCooker) *cookCongeeCommand {
	return &cookCongeeCommand{
		electricCooker: electricCooker,
	}
}

func (c *cookCongeeCommand) Execute() string {
	c.electricCooker.SetFire("大")
	c.electricCooker.SetPressure("强")
	return "煮粥:" + c.electricCooker.Run("45分钟")
}

// shutdownCommand 停止指令
type shutdownCommand struct {
	electricCooker *ElectricCooker
}

func NewShutdownCommand(electricCooker *ElectricCooker) *shutdownCommand {
	return &shutdownCommand{
		electricCooker: electricCooker,
	}
}

func (s *shutdownCommand) Execute() string {
	return s.electricCooker.Shutdown()
}

// ElectricCookerInvoker 电饭煲指令触发器
type ElectricCookerInvoker struct {
	cookCommand CookCommand
}

// SetCookCommand 设置指令
func (e *ElectricCookerInvoker) SetCookCommand(cookCommand CookCommand) {
	e.cookCommand = cookCommand
}

// ExecuteCookCommand 执行指令
func (e *ElectricCookerInvoker) ExecuteCookCommand() string {
	return e.cookCommand.Execute()
}
