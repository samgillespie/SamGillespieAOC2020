package main

import (
	"fmt"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsInt(15)

	solution := PlayElfGame(data, 2020)
	fmt.Println("Day 15 Part 1 Solution:")
	fmt.Println(solution)

	solution = PlayElfGame(data, 30000000)
	fmt.Println("Day 15 Part 2 Solution:")
	fmt.Println(solution)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func PlayElfGame(data []int, goalNumber int) int {
	// Make two maps, last position and current position
	// Bit fiddly, and uses twice as much memory as it should but :shrug:
	position := make(map[int]int)
	prevPosition := make(map[int]int)
	var spoken int
	var inMap bool
	for index := 0; index < goalNumber; index += 1 {
		if index < len(data) {
			spoken = data[index]
			inMap = false
			position[spoken] = index
		} else if inMap == false {
			spoken = 0
			_, inMap = position[spoken]
			if inMap == true {
				prevPosition[spoken] = position[spoken]
			}
			position[spoken] = index
		} else {
			prevPos, _ := prevPosition[spoken]
			spoken = index - 1 - prevPos
			_, inMap = position[spoken]
			prevPosition[spoken] = position[spoken]
			position[spoken] = index
		}
	}
	return spoken
}
