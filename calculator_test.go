package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
		{a: -5, b: 3, want: -2},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: 1, b: 1, want: 0},
		{a: 5, b: 0, want: 5},
		{a: 5, b: -3, want: 8},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 4, b: 2, want: 8},
		{a: 1, b: 1, want: 1},
		{a: 5, b: 0, want: 0},
		{a: 5, b: -3, want: -15},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b        float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{a: 4, b: 2, want: 2, errExpected: false},
		{a: 1, b: 1, want: 1, errExpected: false},
		{a: 5, b: 0, want: 0, errExpected: true},
		{a: -15, b: -3, want: 5, errExpected: false},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			if !tc.errExpected {
				t.Fatalf(err.Error())
			}
		} else if tc.want != got {
			t.Fatalf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a           float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{a: 4, want: 2, errExpected: false},
		{a: 1, want: 1, errExpected: false},
		{a: 25, want: 5, errExpected: false},
		{a: -25, want: 0, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			if !tc.errExpected {
				t.Fatalf(err.Error())
			}
		} else if tc.want != got {
			t.Fatalf("Sqrt(%f): want %f, got %f", tc.a, tc.want, got)
		}
	}
}

func TestEvaluate(t *testing.T) {
	t.Parallel()
	type testCase struct {
		exp         string
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{exp: "2 + 2", want: 4, errExpected: false},
		{exp: "25 - 5", want: 20, errExpected: false},
		{exp: "25 * 4", want: 100, errExpected: false},
		{exp: "25 / 5", want: 5, errExpected: false},
		{exp: "5 / 0", want: 0, errExpected: true},
		{exp: "5*0", want: 0, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Evaluate(tc.exp)
		if err != nil {
			if !tc.errExpected {
				t.Fatalf(err.Error())
			}
		} else if tc.want != got {
			t.Fatalf("Evaluate(\"%s\"): want %f, got %f", tc.exp, tc.want, got)
		}
	}
}
