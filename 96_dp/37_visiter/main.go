package main

import (
	"fmt"
	"sort"
)

// Employee 员工接口
type Employee interface {
	KPI() string                    // 完成kpi信息
	Accept(visitor EmployeeVisitor) // 接受访问者对象
}

// productManager 产品经理
type productManager struct {
	name         string // 名称
	productNum   int    // 上线产品数
	satisfaction int    // 平均满意度
}

func NewProductManager(name string, productNum int, satisfaction int) *productManager {
	return &productManager{
		name:         name,
		productNum:   productNum,
		satisfaction: satisfaction,
	}
}

func (p *productManager) KPI() string {
	return fmt.Sprintf("产品经理%s，上线%d个产品，平均满意度为%d", p.name, p.productNum, p.satisfaction)
}

func (p *productManager) Accept(visitor EmployeeVisitor) {
	visitor.VisitProductManager(p)
}

// softwareEngineer 软件工程师
type softwareEngineer struct {
	name           string // 姓名
	requirementNum int    // 完成需求数
	bugNum         int    // 修复问题数
}

func NewSoftwareEngineer(name string, requirementNum int, bugNum int) *softwareEngineer {
	return &softwareEngineer{
		name:           name,
		requirementNum: requirementNum,
		bugNum:         bugNum,
	}
}

func (s *softwareEngineer) KPI() string {
	return fmt.Sprintf("软件工程师%s，完成%d个需求，修复%d个问题", s.name, s.requirementNum, s.bugNum)
}

func (s *softwareEngineer) Accept(visitor EmployeeVisitor) {
	visitor.VisitSoftwareEngineer(s)
}

// hr 人力资源
type hr struct {
	name       string // 姓名
	recruitNum int    // 招聘人数
}

func NewHR(name string, recruitNum int) *hr {
	return &hr{
		name:       name,
		recruitNum: recruitNum,
	}
}

func (h *hr) KPI() string {
	return fmt.Sprintf("人力资源%s，招聘%d名员工", h.name, h.recruitNum)
}

func (h *hr) Accept(visitor EmployeeVisitor) {
	visitor.VisitHR(h)
}

// EmployeeVisitor 员工访问者接口
type EmployeeVisitor interface {
	VisitProductManager(pm *productManager)     // 访问产品经理
	VisitSoftwareEngineer(se *softwareEngineer) // 访问软件工程师
	VisitHR(hr *hr)                             // 访问人力资源
}

// kpi kpi对象
type kpi struct {
	name string // 完成kpi姓名
	sum  int    // 完成kpi总数量
}

// kpiTopVisitor 员工kpi排名访问者
type kpiTopVisitor struct {
	top []*kpi
}

func (k *kpiTopVisitor) VisitProductManager(pm *productManager) {
	k.top = append(k.top, &kpi{
		name: pm.name,
		sum:  pm.productNum + pm.satisfaction,
	})
}

func (k *kpiTopVisitor) VisitSoftwareEngineer(se *softwareEngineer) {
	k.top = append(k.top, &kpi{
		name: se.name,
		sum:  se.requirementNum + se.bugNum,
	})
}

func (k *kpiTopVisitor) VisitHR(hr *hr) {
	k.top = append(k.top, &kpi{
		name: hr.name,
		sum:  hr.recruitNum,
	})
}

// Publish 发布KPI排行榜
func (k *kpiTopVisitor) Publish() {
	sort.Slice(k.top, func(i, j int) bool {
		return k.top[i].sum > k.top[j].sum
	})
	for i, curKPI := range k.top {
		fmt.Printf("第%d名%s：完成KPI总数%d\n", i+1, curKPI.name, curKPI.sum)
	}
}

// salaryVisitor 薪酬访问者
type salaryVisitor struct{}

func (s *salaryVisitor) VisitProductManager(pm *productManager) {
	fmt.Printf("产品经理基本薪资：1000元，KPI单位薪资：100元，")
	fmt.Printf("%s，总工资为%d元\n", pm.KPI(), (pm.productNum+pm.satisfaction)*100+1000)
}

func (s *salaryVisitor) VisitSoftwareEngineer(se *softwareEngineer) {
	fmt.Printf("软件工程师基本薪资：1500元，KPI单位薪资：80元，")
	fmt.Printf("%s，总工资为%d元\n", se.KPI(), (se.requirementNum+se.bugNum)*80+1500)
}

func (s *salaryVisitor) VisitHR(hr *hr) {
	fmt.Printf("人力资源基本薪资：800元，KPI单位薪资：120元，")
	fmt.Printf("%s，总工资为%d元\n", hr.KPI(), hr.recruitNum*120+800)
}
