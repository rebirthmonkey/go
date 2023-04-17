package demo

import (
	"testing"
)

// 为以上函数编写单元测试代码
// 1. 测试两个相等的数字
// 2. 测试两个不相等的数字

// 1. 测试两个相等的数字
func TestEqualTrue(t *testing.T) {
	a := 1
	b := 1
	expected := true
	actual := equal(a, b)
	if actual != expected {
		t.Errorf("equal(%d, %d) should be %v, but is %v\n", a, b, expected, actual)
	}
}

// 2. 测试两个不相等的数字
func TestEqualFalse(t *testing.T) {
	a := 1
	b := 2
	expected := true
	actual := equal(a, b)
	if actual != expected {
		t.Errorf("equal(%d, %d) should be %v, but is %v\n", a, b, expected, actual)
	}
}
