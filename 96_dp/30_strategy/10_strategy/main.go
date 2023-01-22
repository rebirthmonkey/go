package main

import (
	"fmt"
)

type IStrategy interface {
	do(int, int) int
}

// Strategy实现：加
type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

// Strategy实现：减
type reduce struct{}

func (*reduce) do(a, b int) int {
	return a - b
}

// Operator 用于在不同的 strategy 间 switch.
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) do(a, b int) int {
	return operator.strategy.do(a, b)
}

func main() {
	operator := Operator{}

	operator.setStrategy(&add{})
	result := operator.do(1, 2)
	fmt.Println("add 1+2=", result)

	operator.setStrategy(&reduce{})
	result = operator.do(2, 1)
	fmt.Println("reduce 2-1=", result)
}
