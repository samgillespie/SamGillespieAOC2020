package main

import (
	"fmt"

	lib "./lib"
)

func main() {
	data := lib.ReadInputAsInt(1)
	part1(data)
	part2(data)
}

func part1(data []int) {
	// Part 1
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			sum := data[i] + data[j]
			if sum == 2020 {
				fmt.Println("Part 1 Solution")
				fmt.Println(data[i] * data[j])
				break
			}
		}
	}
}

func part2(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for k := j + 1; k < len(data); k++ {
				if k <= j {
					continue
				}

				sum := data[i] + data[j] + data[k]
				if sum == 2020 {
					fmt.Println("Part 2 Solution")
					fmt.Println(data[i] * data[j] * data[k])
					return
				}
			}
		}
	}
}
