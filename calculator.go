// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Add takes two numbers and returns the result of adding them together.
func Add(inputs ...float64) float64 {
	sum := 0.0
	for _, input := range inputs {
		sum += input
	}
	return sum
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(inputs ...float64) float64 {
	difference := 0.0
	for idx, input := range inputs {
		if idx == 0 {
			difference = input
		} else {
			difference -= input
		}
	}
	return difference
}

// Multiply takes two numbers and returns the result of multiplying the second
// by the first.
func Multiply(inputs ...float64) float64 {
	product := 1.0
	for _, input := range inputs {
		product *= input
	}
	return product
}

// Divide takes two numbers and returns the result of dividing the first
// by the second.
func Divide(inputs ...float64) (float64, error) {
	quotient := 1.0
	for idx, input := range inputs {
		if idx == 0 {
			quotient = input
		} else {
			if input == 0 {
				return 0, fmt.Errorf("bad input: %f (division by zero is undefined)", input)
			}
			quotient /= input
		}
	}
	return quotient, nil
}

// Sqrt takes a number and returns the square root of that number.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input: %f (Square Roots of negative numbers is undefined)", a)
	}
	return math.Sqrt(a), nil
}

// Evaluate takes a string representation of a math expression, then figures out which function to call
// based on the operator
func Evaluate(exp string) (float64, error) {
	result := 1.0
	// parse the string - split by spaces
	strArr := strings.Split(exp, " ")
	firstVal, err := strconv.ParseFloat(strArr[0], 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid string to float conversion for input: %s", exp)
	}
	operator := strArr[1]
	secondVal, err := strconv.ParseFloat(strArr[2], 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid string to float conversion for input: %s", exp)
	}

	switch operator {
	case "+":
		result = Add(firstVal, secondVal)
	case "-":
		result = Subtract(firstVal, secondVal)
	case "*":
		result = Multiply(firstVal, secondVal)
	case "/":
		result, err = Divide(firstVal, secondVal)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}
