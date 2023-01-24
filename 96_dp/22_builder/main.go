package main

// Quantity 分量
type Quantity int

const (
	Small  Quantity = 1
	Middle Quantity = 5
	Large  Quantity = 10
)

type PancakeBuilder interface {
	// PutPaste 放面糊
	PutPaste(quantity Quantity)
	// PutEgg 放鸡蛋
	PutEgg(num int)
	// PutWafer 放薄脆
	PutWafer()
	// PutFlavour 放调料 Coriander香菜，Shallot葱 Sauce酱
	PutFlavour(hasCoriander, hasShallot, hasSauce bool)
	// Build 摊煎饼
	Build() *Pancake
}

// Pancake  煎饼
type Pancake struct {
	pasteQuantity Quantity // 面糊分量
	eggNum        int      // 鸡蛋数量
	wafer         string   // 薄脆
	hasCoriander  bool     // 是否放香菜
	hasShallot    bool     // 是否放葱
	hasSauce      bool     // 是否放酱
}

type normalPancakeBuilder struct {
	pasteQuantity Quantity // 面糊量
	eggNum        int      // 鸡蛋数量
	friedWafer    string   // 油炸薄脆
	hasCoriander  bool     // 是否放香菜
	hasShallot    bool     // 是否放葱
	hasHotSauce   bool     // 是否放辣味酱
}

func NewNormalPancakeBuilder() *normalPancakeBuilder {
	return &normalPancakeBuilder{}
}

func (n *normalPancakeBuilder) PutPaste(quantity Quantity) {
	n.pasteQuantity = quantity
}

func (n *normalPancakeBuilder) PutEgg(num int) {
	n.eggNum = num
}

func (n *normalPancakeBuilder) PutWafer() {
	n.friedWafer = "油炸的薄脆"
}

func (n *normalPancakeBuilder) PutFlavour(hasCoriander, hasShallot, hasSauce bool) {
	n.hasCoriander = hasCoriander
	n.hasShallot = hasShallot
	n.hasHotSauce = hasSauce
}

func (n *normalPancakeBuilder) Build() *Pancake {
	return &Pancake{
		pasteQuantity: n.pasteQuantity,
		eggNum:        n.eggNum,
		wafer:         n.friedWafer,
		hasCoriander:  n.hasCoriander,
		hasShallot:    n.hasShallot,
		hasSauce:      n.hasHotSauce,
	}
}

type healthyPancakeBuilder struct {
	milletPasteQuantity Quantity // 小米面糊量
	chaiEggNum          int      // 柴鸡蛋数量
	nonFriedWafer       string   // 非油炸薄脆
	hasCoriander        bool     // 是否放香菜
	hasShallot          bool     // 是否放葱
}

func NewHealthyPancakeBuilder() *healthyPancakeBuilder {
	return &healthyPancakeBuilder{}
}

func (n *healthyPancakeBuilder) PutPaste(quantity Quantity) {
	n.milletPasteQuantity = quantity
}

func (n *healthyPancakeBuilder) PutEgg(num int) {
	n.chaiEggNum = num
}

func (n *healthyPancakeBuilder) PutWafer() {
	n.nonFriedWafer = "非油炸的薄脆"
}

func (n *healthyPancakeBuilder) PutFlavour(hasCoriander, hasShallot, _ bool) {
	n.hasCoriander = hasCoriander
	n.hasShallot = hasShallot
}

func (n *healthyPancakeBuilder) Build() *Pancake {
	return &Pancake{
		pasteQuantity: n.milletPasteQuantity,
		eggNum:        n.chaiEggNum,
		wafer:         n.nonFriedWafer,
		hasCoriander:  n.hasCoriander,
		hasShallot:    n.hasShallot,
		hasSauce:      false,
	}
}

// PancakeCook 摊煎饼师傅
type PancakeCook struct {
	builder PancakeBuilder
}

func NewPancakeCook(builder PancakeBuilder) *PancakeCook {
	return &PancakeCook{
		builder: builder,
	}
}

// SetPancakeBuilder 重新设置煎饼构造器
func (p *PancakeCook) SetPancakeBuilder(builder PancakeBuilder) {
	p.builder = builder
}

// MakePancake 摊一个一般煎饼
func (p *PancakeCook) MakePancake() *Pancake {
	p.builder.PutPaste(Middle)
	p.builder.PutEgg(1)
	p.builder.PutWafer()
	p.builder.PutFlavour(true, true, true)
	return p.builder.Build()
}

// MakeBigPancake 摊一个巨无霸煎饼
func (p *PancakeCook) MakeBigPancake() *Pancake {
	p.builder.PutPaste(Large)
	p.builder.PutEgg(3)
	p.builder.PutWafer()
	p.builder.PutFlavour(true, true, true)
	return p.builder.Build()
}

// MakePancakeForFlavour 摊一个自选调料霸煎饼
func (p *PancakeCook) MakePancakeForFlavour(hasCoriander, hasShallot, hasSauce bool) *Pancake {
	p.builder.PutPaste(Large)
	p.builder.PutEgg(3)
	p.builder.PutWafer()
	p.builder.PutFlavour(hasCoriander, hasShallot, hasSauce)
	return p.builder.Build()
}
