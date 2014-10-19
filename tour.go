package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	finalSlice := make([][]uint8, dy)
	for i := range finalSlice {
		innerSlice := make([]uint8, dx)
		for j := range innerSlice {
			innerSlice[j] = uint8(i) ^ uint8(j)
		}
		finalSlice[i] = innerSlice
	}
	return finalSlice
}

func main() {
    pic.Show(Pic)
}