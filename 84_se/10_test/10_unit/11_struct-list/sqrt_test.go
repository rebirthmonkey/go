package demo

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	var tests = []struct {
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

	for _, ts := range tests {
		actual := Sqrt(ts.input)
		assert.Equal(t, actual, ts.expected, "Sqrt(%f) should be %f, but got %f", ts.input, ts.expected, actual)
	}
}
