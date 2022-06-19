package demo

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	var shouldSuccess = []struct {
		input    float64 // input
		expected float64 // expected result
	}{
		{math.Inf(1), math.Inf(1)}, // positive infinity
		{math.Inf(-1), math.NaN()}, // negative infinity
		{-1.0, math.NaN()},
		{0.0, 0.0},
		{-0.0, -0.0},
		{1.0, 1.0},
		{4.0, 2.0},
	}

	for _, ts := range shouldSuccess {
		if actual := Sqrt(ts.input); actual != ts.expected {
			t.Fatalf("Sqrt(%f) should be %v, but is: %v\n", ts.input, ts.expected, actual)
		}
	}
}
