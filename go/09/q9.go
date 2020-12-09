package main

import (
	"fmt"
	"math"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsInt(9)
	invalid_entry := q9part1(data)

	fmt.Println("Question 9 Part 1 Answer")
	fmt.Println(invalid_entry)

	q9part2(data, invalid_entry)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func calculate_valid_entries(data []int, index int, preamble int) []int {
	var validEntries []int

	for index1 := int(math.Max(float64(index-preamble), 0)); index1 < index; index1++ {
		for index2 := int(math.Max(float64(index-preamble), 1)); index2 < index; index2++ {
			if index1 == index2 {
				continue
			}
			validEntries = append(validEntries, data[index1]+data[index2])
		}
	}
	return validEntries
}

func q9part1(data []int) int {
	// Part 1
	// {"position in list": [Valid sums of previous 25]}
	preamble := 25
	for index, entry := range data {
		if index < 25 {
			continue
		}
		validEntries := calculate_valid_entries(data, index, preamble)
		isValid := false
		for validEntry := range validEntries {
			if validEntries[validEntry] == entry {
				isValid = true
				break
			}
		}
		if isValid == false {
			return entry
		}
	}
	return -1
}

func q9part2(data []int, invalid_entry int) {
	// Part 2
	for startIndex := 0; startIndex < len(data); startIndex++ {
		sum := data[startIndex]
		cursor := startIndex
		largest := float64(0)
		smallest := float64(9999999999)
		for {
			cursor++
			sum += data[cursor]
			largest = math.Max(largest, float64(data[cursor]))
			smallest = math.Min(smallest, float64(data[cursor]))
			if sum == invalid_entry {
				fmt.Println("Question 9 Part 2 Answer")
				fmt.Println(int(smallest + largest))
				return
			} else if sum > invalid_entry {
				break
			}
		}
	}
}
