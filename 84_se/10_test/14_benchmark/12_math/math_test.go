package test

import (
	"testing"
)

// 编写Abs函数的Benchmark测试
func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(10)
	}
}

// 编写Abs函数的Benchmark测试
func BenchmarkAbs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(20)
	}
}

// 编写Max函数的Benchmark测试
func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(10, 20)
	}
}

// 编写Max函数的Benchmark测试
func BenchmarkMax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(30, 20)
	}
}

// 编写Min函数的Benchmark测试
func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(10, 20)
	}
}

// 编写Min函数的Benchmark测试
func BenchmarkMin2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(30, 20)
	}
}

// 编写RandInt函数的Benchmark测试
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt()
	}
}
