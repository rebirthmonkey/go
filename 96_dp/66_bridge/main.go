package main

import "fmt"

// Traffic 交通工具
type Traffic interface {
	Transport() string
}

// airplane 飞机
type airplane struct{}

// Transport 坐飞机
func (a *airplane) Transport() string {
	return "by airplane"
}

// car 汽车
type car struct{}

// Transport 坐汽车
func (t *car) Transport() string {
	return "by car"
}

// Location 地点
type Location interface {
	Name() string       // 地点名称
	PlaySports() string // 参与运动
}

// namedLocation 被命名的地点，统一引用此类型，声明名字字段及获取方法
type namedLocation struct {
	name string
}

// Name 获取地点名称
func (n namedLocation) Name() string {
	return n.name
}

// seaside 海边
type seaside struct {
	namedLocation
}

// NewSeaside 创建指定名字的海边，比如三亚湾
func NewSeaside(name string) *seaside {
	return &seaside{
		namedLocation: namedLocation{
			name: name,
		},
	}
}

// PlaySports 海边可以冲浪
func (s *seaside) PlaySports() string {
	return fmt.Sprintf("surfing")
}

// mountain 山
type mountain struct {
	namedLocation
}

// NewMountain 创建指定名字的山，比如泰山
func NewMountain(name string) *mountain {
	return &mountain{
		namedLocation: namedLocation{
			name: name,
		},
	}
}

// PlaySports 可以爬山
func (m *mountain) PlaySports() string {
	return fmt.Sprintf("climbing")
}

// desert 荒漠
type desert struct {
	namedLocation
}

// NewDesert 创建指定名字的荒漠，比如罗布泊
func NewDesert(name string) *desert {
	return &desert{
		namedLocation: namedLocation{
			name: name,
		},
	}
}

// PlaySports 荒漠可以徒步穿越
func (d *desert) PlaySports() string {
	return fmt.Sprintf("trekking")
}

// Experience 经历
type Experience interface {
	Describe() string // 描述经历
}

// travelExperience 旅游经历
type travelExperience struct {
	subject  string
	traffic  Traffic
	location Location
}

// NewTravelExperience 创建旅游经历，包括主题、交通方式、地点
func NewTravelExperience(subject string, traffic Traffic, location Location) *travelExperience {
	return &travelExperience{
		subject:  subject,
		traffic:  traffic,
		location: location,
	}
}

// Describe 描述旅游经历
func (t *travelExperience) Describe() string {
	return fmt.Sprintf("%s is to %s %s and %s", t.subject, t.location.Name(), t.traffic.Transport(), t.location.PlaySports())
}

// adventureExperience 探险经历
type adventureExperience struct {
	survivalTraining string
	travelExperience
}

// NewAdventureExperience 创建探险经历，包括探险需要的培训，其他的与路由参数类似
func NewAdventureExperience(training string, subject string, traffic Traffic, location Location) *adventureExperience {
	return &adventureExperience{
		survivalTraining: training,
		travelExperience: *NewTravelExperience(subject, traffic, location),
	}
}

// Describe 描述探险经历
func (a *adventureExperience) Describe() string {
	return fmt.Sprintf("after %s, %s", a.survivalTraining, a.travelExperience.Describe())
}
