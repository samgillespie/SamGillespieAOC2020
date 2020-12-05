package main

import (
	"fmt"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(5)
	max_seat_id := q5part1(data)
	fmt.Println("Question 5 Part 1 Solution:")
	fmt.Println(max_seat_id)

	q5part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func ConvertPassToInt(pass string) []int {
	row_min := 0
	row_max := 127
	for i := 0; i < 7; i++ {
		if pass[i] == 'F' {
			row_max = row_max - (row_max-row_min+2)/2 // Add +2 for rounding
		} else if pass[i] == 'B' {
			row_min = row_min + (row_max-row_min+2)/2 // Add +2 for rounding
		} else {
			panic(new(error))
		}
	}

	// Sanity Check
	if row_min != row_max {
		panic(new(error))
	}

	column_min := 0
	column_max := 7
	for i := 7; i <= 9; i++ {
		if pass[i] == 'L' {
			column_max = column_max - (column_max-column_min+2)/2
		} else if pass[i] == 'R' {
			column_min = column_min + (column_max-column_min+2)/2
		} else {
			panic(new(error))
		}
	}

	// Sanity Check
	if column_min != column_max {
		panic(new(error))
	}

	seat_id := row_min*8 + column_min
	return []int{row_min, column_min, seat_id}
}
func q5part1(data []string) int {
	max_seat_id := 0
	for i := 0; i < len(data); i++ {
		entry := data[i]
		result := ConvertPassToInt(entry)
		if max_seat_id < result[2] {
			max_seat_id = result[2]
		}
	}
	return max_seat_id
}

func q5part2(data []string) {
	seats := make(map[int]bool)
	max_seat_num := q5part1(data)
	for i := 32; i < max_seat_num-1; i++ {
		seats[i] = true
	}

	for i := 0; i < len(data); i++ {
		entry := data[i]
		result := ConvertPassToInt(entry)
		delete(seats, result[2])
	}
	for k := range seats {
		fmt.Println("Question 5 Part 2 Solution:")
		fmt.Println(k)
	}

}
