package main

import "testing"

func TestVisitor(t *testing.T) {
	allEmployees := AllEmployees() // 获取所有员工
	kpiTop := new(kpiTopVisitor)   // 创建KPI排行访问者
	VisitAllEmployees(kpiTop, allEmployees)
	kpiTop.Publish() // 发布排行榜

	salary := new(salaryVisitor) // 创建薪酬访问者
	VisitAllEmployees(salary, allEmployees)
}

// VisitAllEmployees 遍历所有员工调用访问者
func VisitAllEmployees(visitor EmployeeVisitor, allEmployees []Employee) {
	for _, employee := range allEmployees {
		employee.Accept(visitor)
	}
}

// AllEmployees 获得所有公司员工
func AllEmployees() []Employee {
	var employees []Employee
	employees = append(employees, NewHR("小明", 10))
	employees = append(employees, NewProductManager("小红", 4, 7))
	employees = append(employees, NewSoftwareEngineer("张三", 10, 5))
	employees = append(employees, NewSoftwareEngineer("李四", 3, 6))
	employees = append(employees, NewSoftwareEngineer("王五", 7, 1))
	return employees
}
