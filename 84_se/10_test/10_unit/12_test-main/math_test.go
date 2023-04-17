package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("do some setup before starting the test")
	m.Run()
	fmt.Println("do some cleanup after finishing the test")
}

func TestAbs(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{-0.3, 0.3},
		{-2, 2},
		{-3.1, 3.1},
		{5, 5},
	}

	for _, tt := range tests {
		actual := Abs(tt.input)
		assert.Equal(t, actual, tt.expected, "Abs test")
	}
}

// 编写Max函数的单元测试，使用表驱动测试，表为{{1, 2} {2, 3} {3, 4}}，期望结果为{2, 3, 4}，并且使用assert包进行断言。
func TestMax(t *testing.T) {
	tests := []struct {
		x, y float64
		want float64
	}{
		{1, 2, 2},
		{2, 3, 3},
		{3, 4, 4},
	}

	for _, tt := range tests {
		got := Max(tt.x, tt.y)
		assert.Equal(t, got, tt.want, "Max test")
	}
}

// 编写Min函数的单元测试，使用表驱动测试，表为{{1, 2} {2, 3} {3, 4}}，期望结果为{1, 3, 4}，并且使用assert包进行断言。
func TestMin(t *testing.T) {
	tests := []struct {
		x, y float64
		want float64
	}{
		{1, 2, 1},
		{2, 3, 2},
		{3, 4, 3},
	}

	for _, tt := range tests {
		got := Min(tt.x, tt.y)
		assert.Equal(t, got, tt.want, "Min test")
	}
}
