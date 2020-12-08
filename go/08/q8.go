package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(8)
	q8part1(data)
	q8part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func q8part1(data []string) {
	accumulator := 0

	indexes_hit := make(map[int]int)
	row_index := 0
	for {
		// Check if we've already processed this row
		if indexes_hit[row_index] == 1 {
			fmt.Println("Question 8 Part 1 Solution")
			fmt.Println(accumulator)
			return
		}
		indexes_hit[row_index] += 1

		row := data[row_index]
		str_split := strings.Split(row, " ")
		if str_split[0] == "nop" {
			row_index += 1
		} else if str_split[0] == "acc" {
			change, _ := strconv.Atoi(str_split[1])
			accumulator += change
			row_index += 1
		} else if str_split[0] == "jmp" {
			change, _ := strconv.Atoi(str_split[1])
			row_index += change
		}

	}
}

func q8part2(data []string) {

	for changed_instr_int := 0; changed_instr_int < len(data); changed_instr_int++ {
		temp_data := make([]string, len(data))
		copy(temp_data, data)
		// Change the nth entry to the other, and see if it terminates
		changed_int_split := strings.Split(temp_data[changed_instr_int], " ")
		if changed_int_split[0] == "jmp" {
			temp_data[changed_instr_int] = "nop " + changed_int_split[1]
		} else if changed_int_split[0] == "nop" {
			temp_data[changed_instr_int] = "jmp " + changed_int_split[1]
		}

		indexes_hit := make(map[int]int)
		row_index := 0
		accumulator := 0
		for {
			// Check if we've already processed this row
			if indexes_hit[row_index] == 1 {
				break
			}
			if row_index >= len(temp_data) {
				fmt.Println("Question 8 Part 2 Solution")
				fmt.Println(accumulator)
				return
			}
			indexes_hit[row_index] += 1

			row := temp_data[row_index]
			str_split := strings.Split(row, " ")
			if str_split[0] == "nop" {
				row_index += 1
			} else if str_split[0] == "acc" {
				change, _ := strconv.Atoi(str_split[1])
				accumulator += change
				row_index += 1
			} else if str_split[0] == "jmp" {
				change, _ := strconv.Atoi(str_split[1])
				row_index += change
			}
		}
	}
}
