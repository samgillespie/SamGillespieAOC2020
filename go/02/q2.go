package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(2)
	q2part1(data)
	q2part2(data)
	elapsed := time.Since(start)

	log.Printf("Main took %s", elapsed)
}

func q2part1(data []string) {
	// Part 1
	valid_passwords := 0
	for i := 0; i < len(data); i++ {
		substring := strings.Split(data[i], " ")
		split_interval := strings.Split(substring[0], "-")
		interval_min, _ := strconv.Atoi(split_interval[0])
		interval_max, _ := strconv.Atoi(split_interval[1])
		letter := substring[1][0]
		password := substring[2]

		counter := 0
		for char_index := 0; char_index < len(password); char_index++ {
			if password[char_index] == letter {
				counter += 1
			}
		}

		if counter >= interval_min && counter <= interval_max {
			valid_passwords += 1
		}
	}
	fmt.Println("Question 2 Part 1 Solution:")
	fmt.Println(valid_passwords)
}

func q2part2(data []string) {
	// Part 1
	valid_passwords := 0
	for i := 0; i < len(data); i++ {
		substring := strings.Split(data[i], " ")
		split_interval := strings.Split(substring[0], "-")
		interval_min, _ := strconv.Atoi(split_interval[0])
		interval_max, _ := strconv.Atoi(split_interval[1])
		letter := substring[1][0]
		password := substring[2]

		if password[interval_min-1] == letter || password[interval_max-1] == letter {
			if password[interval_min-1] != password[interval_max-1] {
				valid_passwords += 1
			}
		}
	}
	fmt.Println("Question 2 Part 2 Solution:")
	fmt.Println(valid_passwords)
}
