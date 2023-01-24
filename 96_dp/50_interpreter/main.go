package main

import "strings"

// Expression 表达式接口，包含一个解释方法
type Expression interface {
	Interpret(context string) bool
}

// terminalExpression 终结符表达式，判断表达式中是否包含匹配数据
type terminalExpression struct {
	matchData string
}

func NewTerminalExpression(matchData string) *terminalExpression {
	return &terminalExpression{matchData: matchData}
}

// Interpret 判断是否包含匹配字符
func (t *terminalExpression) Interpret(context string) bool {
	if strings.Contains(context, t.matchData) {
		return true
	}
	return false
}

// orExpression 或表达式
type orExpression struct {
	left, right Expression
}

func NewOrExpression(left, right Expression) *orExpression {
	return &orExpression{
		left:  left,
		right: right,
	}
}

func (o *orExpression) Interpret(context string) bool {
	return o.left.Interpret(context) || o.right.Interpret(context)
}

// andExpression 与表达式
type andExpression struct {
	left, right Expression
}

func NewAndExpression(left, right Expression) *andExpression {
	return &andExpression{
		left:  left,
		right: right,
	}
}

func (o *andExpression) Interpret(context string) bool {
	return o.left.Interpret(context) && o.right.Interpret(context)
}
