package main

import (
	"fmt"
	"sort"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsInt(10)
	q10part1(data)
	q10part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func q10part1(data []int) {
	// Sort the list
	sort.Ints(data)
	data = append(data, data[len(data)-1]+3)
	current_jolts := 0
	jumps := make(map[int]int)
	for _, value := range data {
		diff := value - current_jolts
		jumps[diff] += 1
		current_jolts = value
	}
	fmt.Println("Day 10 Part 1 Answer")
	fmt.Println(jumps[1] * jumps[3])
}

func tribonacci(n int) map[int]int {
	trib := make(map[int]int, n)
	trib[1] = 1
	trib[2] = 1
	trib[3] = 2
	for i := 4; i < n; i++ {
		trib[i] = trib[i-1] + trib[i-2] + trib[i-3]
	}
	return trib
}

func q10part2(data []int) {
	data = append(data, 0)

	sort.Ints(data)
	data = append(data, data[len(data)-1]+3)
	permutations := 1
	tribonacciMap := tribonacci(50)
	index := 0
	for index+1 < len(data) {
		run_length := 1
		// Break the problem into a number of runs of 1
		i := index
		for {
			if data[i+1]-data[i] == 1 {
				run_length += 1
			} else {
				break
			}
			i += 1
		}
		permutations = permutations * tribonacciMap[run_length]
		index += run_length
	}

	fmt.Println("Day 10 Part 2 Answer")
	fmt.Println(permutations)
}
