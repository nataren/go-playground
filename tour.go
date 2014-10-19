package main

import "fmt"

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	} else {
		return fib(i - 1) + fib(i - 2)
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibIter := 0
	return func() int {
		fibIter += 1
		return fib(fibIter)
	}
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}