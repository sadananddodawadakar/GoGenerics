# Generic Maths
This package implements basic mathematical functions (Add, Subtract, Multiply & Devide).
Package implements functionalities using generics with int | float64 constraints.

func Add[T int | float64](a, b T) T

func Subtract[T int | float64](a, b T) T 

func Multiply[T int | float64](a, b T) T 

func Divide[T int | float64](a, b T) (T, error)

# Build instruction
> go build main.go

# Execution
> ./main

