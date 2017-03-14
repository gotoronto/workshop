package main

import "github.com/gotoronto/workshop/slices/display"

func Pic(dx, dy int) [][]int {
	// implement
	output := make([][]int, dy)
	for y := range output {
		output[y] = make([]int, dx)
		for x := range output[y] {
			output[y][x] = y
		}
	}
	return output
}

func main() {
	display.Show(Pic(80, 25))
}
