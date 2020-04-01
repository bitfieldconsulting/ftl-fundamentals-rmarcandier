package calculator_test

import (
	"calculator"
	"testing"
)

func TestSubtract(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 2, b: 1, want: 1},
		{a: 4, b: 2, want: 2},
		{a: 6, b: 6, want: 0},
		{a: 1, b: -2, want: 3},
		{a: -10, b: -2, want: -8},
		{a: 0, b: 7, want: -7},
		{a: 8, b: 0, want: 8},
	}
	for _, testCase := range testCases {
		got := calculator.Subtract(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Subtract(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 2, b: 1, want: 2},
		{a: 4, b: 2, want: 8},
		{a: 6, b: 6, want: 36},
		{a: 1, b: -2, want: -2},
		{a: -10, b: -2, want: 20},
		{a: 0, b: 7, want: 0},
		{a: 8, b: 0, want: 0},
	}
	for _, testCase := range testCases {
		got := calculator.Multiply(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Multiply(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 2, b: 1, want: 2},
		{a: 4, b: 2, want: 2},
		{a: 6, b: 6, want: 1},
		{a: 10, b: 5, want: 2},
		{a: 20, b: 2, want: 10},
		{a: 0, b: 7, want: 0},
		{a: -8, b: 2, want: -4},
	}
	for _, testCase := range testCases {
		got := calculator.Divide(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Divide(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}
