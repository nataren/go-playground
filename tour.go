package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func MySqrt(x float64) (float64, error) {
	if x <= 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return SqrtIter(1.0, x), nil
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
	sqrtOfThree, err := MySqrt(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("MySqrt: ", sqrtOfThree)
	}
	sqrtOfMinusOne, err2 := MySqrt(-3)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("math.Sqrt: ", sqrtOfMinusOne)
	}
}