package main

import "fmt"

// Season 季节的策略接口，不同季节表现得天气不同
type Season interface {
	ShowWeather(city string) string // 显示指定城市的天气情况
}

type spring struct {
	weathers map[string]string // 存储不同城市春天气候
}

func NewSpring() *spring {
	return &spring{
		weathers: map[string]string{"北京": "干燥多风", "昆明": "清凉舒适"},
	}
}

func (s *spring) ShowWeather(city string) string {
	return fmt.Sprintf("%s的春天，%s;", city, s.weathers[city])
}

type summer struct {
	weathers map[string]string // 存储不同城市夏天气候
}

func NewSummer() *summer {
	return &summer{
		weathers: map[string]string{"北京": "高温多雨", "昆明": "清凉舒适"},
	}
}

func (s *summer) ShowWeather(city string) string {
	return fmt.Sprintf("%s的夏天，%s;", city, s.weathers[city])
}

type autumn struct {
	weathers map[string]string // 存储不同城市秋天气候
}

func NewAutumn() *autumn {
	return &autumn{
		weathers: map[string]string{"北京": "凉爽舒适", "昆明": "清凉舒适"},
	}
}

func (a *autumn) ShowWeather(city string) string {
	return fmt.Sprintf("%s的秋天，%s;", city, a.weathers[city])
}

type winter struct {
	weathers map[string]string // 存储不同城市冬天气候
}

func NewWinter() *winter {
	return &winter{
		weathers: map[string]string{"北京": "干燥寒冷", "昆明": "清凉舒适"},
	}
}

func (w *winter) ShowWeather(city string) string {
	return fmt.Sprintf("%s的冬天，%s;", city, w.weathers[city])
}

// City 城市
type City struct {
	name    string
	feature string
	season  Season
}

// NewCity 根据名称及季候特征创建城市
func NewCity(name, feature string) *City {
	return &City{
		name:    name,
		feature: feature,
	}
}

// SetSeason 设置不同季节，类似天气在不同季节的不同策略
func (c *City) SetSeason(season Season) {
	c.season = season
}

// String 显示城市的气候信息
func (c *City) String() string {
	return fmt.Sprintf("%s%s，%s", c.name, c.feature, c.season.ShowWeather(c.name))
}
