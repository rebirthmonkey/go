package main

import (
	"fmt"
)

// Member 成员接口
type Member interface {
	Desc() string // 输出成员描述信息
}

// Teacher 老师
type Teacher struct {
	name    string // 名称
	subject string // 所教课程
}

// NewTeacher 根据姓名、课程创建老师对象
func NewTeacher(name, subject string) *Teacher {
	return &Teacher{
		name:    name,
		subject: subject,
	}
}

func (t *Teacher) Desc() string {
	return fmt.Sprintf("%s班主任老师负责教%s", t.name, t.subject)
}

// Student 学生
type Student struct {
	name     string // 姓名
	sumScore int    // 考试总分数
}

// NewStudent 创建学生对象
func NewStudent(name string, sumScore int) *Student {
	return &Student{
		name:     name,
		sumScore: sumScore,
	}
}

func (t *Student) Desc() string {
	return fmt.Sprintf("%s同学考试总分为%d", t.name, t.sumScore)
}

// Iterator 迭代器接口
type Iterator interface {
	Next() Member  // 迭代下一个成员
	HasMore() bool // 是否还有
}

// memberIterator 班级成员迭代器实现
type memberIterator struct {
	class *Class // 需迭代的班级
	index int    // 迭代索引
}

func (m *memberIterator) Next() Member {
	// 迭代索引为-1时，返回老师成员，否则遍历学生slice
	if m.index == -1 {
		m.index++
		return m.class.teacher
	}
	student := m.class.students[m.index]
	m.index++
	return student
}

func (m *memberIterator) HasMore() bool {
	return m.index < len(m.class.students)
}

// Iterable 可迭代集合接口，实现此接口返回迭代器
type Iterable interface {
	CreateIterator() Iterator
}

// Class 班级，包括老师和同学
type Class struct {
	name     string
	teacher  *Teacher
	students []*Student
}

// NewClass 根据班主任老师名称，授课创建班级
func NewClass(name, teacherName, teacherSubject string) *Class {
	return &Class{
		name:    name,
		teacher: NewTeacher(teacherName, teacherSubject),
	}
}

// CreateIterator 创建班级迭代器
func (c *Class) CreateIterator() Iterator {
	return &memberIterator{
		class: c,
		index: -1, // 迭代索引初始化为-1，从老师开始迭代
	}
}

func (c *Class) Name() string {
	return c.name
}

// AddStudent 班级添加同学
func (c *Class) AddStudent(students ...*Student) {
	c.students = append(c.students, students...)
}
