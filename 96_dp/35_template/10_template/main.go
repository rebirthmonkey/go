package main

import (
	"bytes"
	"fmt"
)

// IActor 演员接口
type IActor interface {
	DressUp() string // 装扮
}

// dressBehavior 装扮的多个行为，这里多个行为是私有的，通过DressUp模版方法调用
type dressBehavior interface {
	makeUp() string // 化妆
	clothe() string // 穿衣
	wear() string   // 配饰
}

// BaseActor 演员基类
type BaseActor struct {
	roleName      string // 扮演角色
	dressBehavior        // 装扮行为
}

// DressUp 统一实现演员接口的DressUp模版方法，装扮过程通过不同装扮行为进行扩展
func (b *BaseActor) DressUp() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("扮演%s的", b.roleName))
	buf.WriteString(b.makeUp())
	buf.WriteString(b.clothe())
	buf.WriteString(b.wear())
	return buf.String()
}

// womanActor 扩展装扮行为的女演员
type womanActor struct {
	BaseActor
}

// NewWomanActor 指定角色创建女演员
func NewWomanActor(roleName string) *womanActor {
	actor := new(womanActor)    // 创建女演员
	actor.roleName = roleName   // 设置角色
	actor.dressBehavior = actor // 将女演员实现的扩展装扮行为，设置给自己的装扮行为接口
	return actor
}

// 化妆
func (w *womanActor) makeUp() string {
	return "女演员涂着口红，画着眉毛；"
}

// 穿衣
func (w *womanActor) clothe() string {
	return "穿着连衣裙；"
}

// 配饰
func (w *womanActor) wear() string {
	return "带着耳环，手拎着包；"
}

// manActor 扩展装扮行为的男演员
type manActor struct {
	BaseActor
}

func NewManActor(roleName string) *manActor {
	actor := new(manActor)
	actor.roleName = roleName
	actor.dressBehavior = actor // 将男演员实现的扩展装扮行为，设置给自己的装扮行为接口
	return actor
}

func (m *manActor) makeUp() string {
	return "男演员刮净胡子，抹上发胶；"
}

func (m *manActor) clothe() string {
	return "穿着一身西装；"
}

func (m *manActor) wear() string {
	return "带上手表，抽着烟；"
}

// NewChildActor 扩展装扮行为的儿童演员
type childActor struct {
	BaseActor
}

func NewChildActor(roleName string) *childActor {
	actor := new(childActor)
	actor.roleName = roleName
	actor.dressBehavior = actor // 将儿童演员实现的扩展装扮行为，设置给自己的装扮行为接口
	return actor
}

func (c *childActor) makeUp() string {
	return "儿童演员抹上红脸蛋；"
}

func (c *childActor) clothe() string {
	return "穿着一身童装；"
}

func (c *childActor) wear() string {
	return "手里拿着一串糖葫芦；"
}
