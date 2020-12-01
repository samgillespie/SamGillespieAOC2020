package main

import (
	"fmt"
	"log"
	"time"

	lib "./lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsInt(1)
	part1(data)
	part2(data)
	elapsed := time.Since(start)

	log.Printf("Main took %s", elapsed)
}

func part1(data []int) {
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

func part2(data []int) {
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
