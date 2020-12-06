package main

import (
	"fmt"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(6)
	q6part1(data)
	q6part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func add_letter_to_set(list []byte, letter byte) []byte {
	for i := 0; i < len(list); i++ {
		if list[i] == letter {
			return list
		}
	}

	return append(list, letter)

}

func valid_characters(mapping map[byte]int, people_count int) int {
	counter := 0
	for i := range mapping {
		if mapping[i] == people_count {
			counter += 1
		}
	}
	return counter
}

func q6part1(data []string) {

	letters := make([]byte, 0)
	counts := make([]int, 0)

	for row := 0; row < len(data); row++ {

		yeses := data[row]
		if yeses == "" {
			counts = append(counts, len(letters))
			letters = make([]byte, 0)
		} else {
			for char_num := 0; char_num < len(yeses); char_num++ {
				character := yeses[char_num]
				letters = add_letter_to_set(letters, character)
			}
		}
	}
	counter := 0
	for i := 0; i < len(counts); i++ {
		counter += counts[i]
	}
	fmt.Println("Question 6 Part 1 Solution:")
	fmt.Println(counter)
}

func q6part2(data []string) {
	letters := make(map[byte]int, 0)
	counts := make([]int, 0)
	people := 0
	for row := 0; row < len(data); row++ {

		yeses := data[row]
		if yeses == "" {
			counts = append(counts, valid_characters(letters, people))
			letters = make(map[byte]int, 0)
			people = 0
		} else {
			people += 1
			for char_num := 0; char_num < len(yeses); char_num++ {
				character := yeses[char_num]
				letters[character] += 1
			}
		}
	}
	counter := 0
	for i := 0; i < len(counts); i++ {
		counter += counts[i]
	}
	fmt.Println("Question 6 Part 2 Solution:")
	fmt.Println(counter)
}
