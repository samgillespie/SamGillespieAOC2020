package main

import (
	"fmt"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsInt(1)
	q1part1(data)
	q1part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func q1part1(data []int) {
	// Part 1
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			sum := data[i] + data[j]
			if sum == 2020 {
				fmt.Println("Question 1 Part 1 Solution:")
				fmt.Println(data[i] * data[j])
				break
			}
		}
	}
}

func q1part2(data []int) {
	// Part 2
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for k := j + 1; k < len(data); k++ {
				sum := data[i] + data[j] + data[k]
				if sum == 2020 {
					fmt.Println("Question 1 Part 2 Solution:")
					fmt.Println(data[i] * data[j] * data[k])
					return
				}
			}
		}
	}
}
