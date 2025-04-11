package math

import "fmt"

func Add[T int | float64](a, b T) T {
	return a + b
}

func Subtract[T int | float64](a, b T) T {
	return a - b
}

func Multiply[T int | float64](a, b T) T {
	return a * b
}

func Divide[T int | float64](a, b T) (T, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}
	return a / b, nil
}

