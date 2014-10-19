package main

import (
	"fmt"
	"math"
)

func MySqrt(x float64) float64 {
	return SqrtIter(1.0, x)
}

func SqrtIter(guess, x float64) float64 {
	for GoodEnough(guess, x) != true {
		guess = Improve(guess, x)
	}
	return guess
}

func GoodEnough(guess, x float64) bool {
	return math.Abs(math.Pow(guess, 2) - x) < 0.001
}

func Improve(guess, x float64) float64 {
	return Average(guess, x / guess)
}

func Average(x, y float64) float64 {
	return (x + y) / 2
}

func main() {
	fmt.Println("MySqrt: ", MySqrt(1))
	fmt.Println("math.Sqrt: ", math.Sqrt(1))

	fmt.Println("MySqrt: ", MySqrt(2))
	fmt.Println("math.Sqrt: ", math.Sqrt(2))

	fmt.Println("MySqrt: ", MySqrt(3))
	fmt.Println("math.Sqrt: ", math.Sqrt(3))

	fmt.Println("MySqrt: ", MySqrt(4))
	fmt.Println("math.Sqrt: ", math.Sqrt(4))

	fmt.Println("MySqrt: ", MySqrt(5))
	fmt.Println("math.Sqrt: ", math.Sqrt(5))

}