package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	// Cases
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")

	case "linux":
		fmt.Println("Linux.")

	default:
		fmt.Printf("%s.", "")
	}
}