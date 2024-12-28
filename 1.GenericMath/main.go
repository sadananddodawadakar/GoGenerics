package main

import (
	"fmt"
	"genericmath/math"
)

func main() {
	fmt.Println("Demo: Exploring generics with maths package")
	res := math.Add(10, 20)
	fmt.Println("Addition result of integers:", res)
	resf := math.Add(1.5, 2.5)
	fmt.Println("Addition result of float:", resf)
	res = math.Subtract(10, 5)
	fmt.Println("Subtract result of integers:", res)
	resf = math.Subtract(10.5, 5.0)
	fmt.Println("Subtract result of float:", resf)
	res = math.Multiply(5, 10)
	fmt.Println("Multiply result of integers:", res)
	resf = math.Multiply(10.5, 4.5)
	fmt.Println("Multiply result of float:", resf)
	res, _ = math.Divide(100, 10)
	fmt.Println("Division  result of integers:", res)
	resf, _ = math.Divide(100.0, 5.5)
	fmt.Println("Division result of float:", resf)
}
