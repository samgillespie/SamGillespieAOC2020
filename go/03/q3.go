package main

import (
	"fmt"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(3)
	q3part1(data)
	q3part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func walk(data []string, x_step int, y_step int) int {
	var cur_position [2]int = [2]int{0, 0}
	x_size := len(data[0])
	y_size := len(data)
	trees := 0
	y := 0
	for y*y_step <= y_size-1 {
		y++
		// Are we at a tree
		if data[cur_position[1]][cur_position[0]] == '#' {
			trees += 1
		}

		// Take a step
		cur_position[1] = y * y_step
		if cur_position[0]+x_step >= x_size {
			cur_position[0] = cur_position[0] - x_size + x_step
		} else {
			cur_position[0] = cur_position[0] + x_step
		}
	}
	return trees
}

func q3part1(data []string) {

	trees := walk(data, 3, 1)

	fmt.Println("Question 3 Part 1 Solution:")
	fmt.Println(trees)
}

func q3part2(data []string) {
	trees_1 := walk(data, 1, 1)
	trees_2 := walk(data, 3, 1)
	trees_3 := walk(data, 5, 1)
	trees_4 := walk(data, 7, 1)
	trees_5 := walk(data, 1, 2)
	fmt.Println("Question 3 Part 2 Solution:")
	fmt.Println(trees_1 * trees_2 * trees_3 * trees_4 * trees_5)
}
