package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	lib "../lib"
)

type Rule struct {
	index int
	value byte  // if "a" or "b"
	cond1 []int // Number of rules to work through
	cond2 []int // If Rules
}

func (r Rule) isValid(pointer int, targetString string, ruleMap map[int]Rule, depth8 int, depth11 int) int {
	// If valid, return the pointer for the next instance to use
	// If invalid return -1
	//fmt.Printf("entering rule %d\n", r.index)
	// depth8 is how many times we use the recursive version.  If we hit a bad condition on this value, parse -2

	if r.value != byte(0) {
		if pointer >= len(targetString) {
			return -1
		}
		if targetString[pointer] == r.value {
			//fmt.Printf("Rule %d, Pointer: %d Valid\n", r.index, pointer)
			return pointer + 1
		} else {
			//fmt.Printf("Rule %d, Pointer: %d invalid\n", r.index, pointer)
			return -1
		}
	}
	currPointer := pointer

	skip2 := false
	skip1 := false
	if r.index == 8 {
		if depth8 > 0 {
			depth8--
		} else {
			skip2 = true
		}
	}
	if r.index == 11 {
		if depth11 > 0 {
			depth11--
		} else {
			skip2 = true
		}
	}
	if len(r.cond2) > 0 && skip2 == false {
		success := true
		for _, condition := range r.cond2 {
			nextRule := ruleMap[condition]

			pointer = nextRule.isValid(pointer, targetString, ruleMap, depth8, depth11)
			if pointer <= -1 {
				success = false
				break
			}
		}
		if success == true {
			//fmt.Printf("Exiting rule %d successfully, pointer=%d\n", r.index, pointer)
			return pointer
		}
	}
	// Reset the pointer if cond2 fails
	pointer = currPointer
	success := true
	if skip1 == false {
		for _, condition := range r.cond1 {
			nextRule := ruleMap[condition]
			pointer = nextRule.isValid(pointer, targetString, ruleMap, depth8, depth11)
			if pointer <= -1 {
				success = false
				break
			}
		}
		if success == true {
			//fmt.Printf("Exiting rule %d successfully, pointer=%d\n\n", r.index, pointer)
			return pointer
		}
	}
	//fmt.Printf("Exiting rule %d unsuccessfully, pointer=%d\n\n", r.index, pointer)
	return -1
}

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(19)
	q19part1(data)
	q19part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func stringToIntArray(input string) []int {
	values := strings.Split(input, " ")
	solution := make([]int, 0)
	for _, value := range values {
		if value == "" || value == " " {
			continue
		}
		intValue, _ := strconv.Atoi(value)
		solution = append(solution, intValue)
	}
	return solution
}

func parseInput(data []string) (map[int]Rule, []string) {
	ruleParser := make(map[int]Rule)
	rulesProcessed := false
	stringsToProcess := make([]string, 0)
	for _, row := range data {
		if row == "" {
			rulesProcessed = true
			continue
		}
		if rulesProcessed == true {
			stringsToProcess = append(stringsToProcess, row)
			continue
		}

		vals := strings.Split(row, ":")
		ruleNumber, _ := strconv.Atoi(vals[0])
		if vals[1] == ` "a"` || vals[1] == ` "b"` {
			rule := Rule{index: ruleNumber, value: vals[1][2]}
			ruleParser[ruleNumber] = rule
			continue
		}

		a := strings.Split(vals[1], "|")

		b := stringToIntArray(a[0])
		if len(a) == 1 {
			rule := Rule{index: ruleNumber, cond1: b}
			ruleParser[ruleNumber] = rule
		} else {
			c := stringToIntArray(a[1])
			rule := Rule{index: ruleNumber, cond1: b, cond2: c}
			ruleParser[ruleNumber] = rule
		}
	}
	return ruleParser, stringsToProcess
}

func q19part1(data []string) {
	ruleParser, inputs := parseInput(data)
	startingRule := ruleParser[0]
	valid := 0
	for _, input := range inputs {

		conforms := startingRule.isValid(0, input, ruleParser, 0, 0)
		success := conforms != -1 && conforms == len(input)
		if success {
			valid += 1
		}
	}
	fmt.Println("Day 19 Part 1 Solution")
	fmt.Println(valid)
}

func q19part2(data []string) {
	ruleParser, inputs := parseInput(data)
	ruleParser[8] = Rule{
		index: 8,
		cond1: []int{42},
		cond2: []int{42, 8},
	}

	ruleParser[11] = Rule{
		index: 11,
		cond1: []int{42, 31},
		cond2: []int{42, 11, 31},
	}

	startingRule := ruleParser[0]
	valid := 0
	for _, input := range inputs {

		depth8 := 0
		depth11 := 0
		success := false
		for success == false {
			for success == false {
				conforms := startingRule.isValid(0, input, ruleParser, depth8, depth11)
				depth8++
				// Longest string is only 60 odd characters long
				if depth8 >= 30 {
					break
				}
				success = conforms != -1 && conforms == len(input)

			}
			// If depth8 is failing at zero, we've exhausted depth11
			if depth8 == 0 {
				break
			}
			depth8 = 0
			depth11++
			// Longest string is only 60 odd characters long, so we don't nee to go super deep
			if depth11 >= 30 {
				break
			}
		}

		if success {
			valid += 1
		}
	}
	fmt.Println("Day 19 Part 2 Solution")
	fmt.Println(valid)
}
