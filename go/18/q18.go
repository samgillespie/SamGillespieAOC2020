package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(18)
	q18part1(data)
	q18part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func evalP1(feed string) int {
	if feed[0] == '(' {
		if feed[len(feed)-1] == ')' {
			feed = feed[1 : len(feed)-1]
		}
	}
	splitValues := strings.Split(feed, " ")
	value, _ := strconv.Atoi(splitValues[0])
	operator := "@"
	for _, elem := range splitValues[1:len(splitValues)] {
		if elem == "" || elem == " " {
			continue
		}
		if elem == "+" || elem == "*" {
			operator = elem
			continue
		}

		iteratorValue, _ := strconv.Atoi(elem)
		if operator == "+" {
			value += iteratorValue
		} else if operator == "*" {
			value *= iteratorValue
		} else {
			panic("Oh no, weird operator")
		}
		operator = "!"
	}

	return value
}

func evalP2(feed string) int {
	// This one applies the addition before the multiplication
	if feed[0] == '(' {
		if feed[len(feed)-1] == ')' {
			feed = feed[1 : len(feed)-1]
		}
	}
	splitValues := strings.Split(feed, " ")
	// Resolve Addition
	toResolve := make([]int, 0)
	for i, term := range splitValues {
		if term == "+" {
			toResolve = append(toResolve, i)
		}
	}
	// Resolve in reverse order
	sort.Sort(sort.Reverse(sort.IntSlice(toResolve)))

	for _, i := range toResolve {
		valuea, _ := strconv.Atoi(splitValues[i-1])
		valueb, _ := strconv.Atoi(splitValues[i+1])
		splitValues[i-1] = strconv.Itoa(valuea + valueb)
		splitValues = append(splitValues[0:i], splitValues[i+2:]...)
	}
	if len(splitValues) == 1 {
		solution, _ := strconv.Atoi(splitValues[0])
		return solution
	}

	// Probably inefficient to throw to P1, but cbf reimplementing everything
	return evalP1(strings.Join(splitValues, " "))
}

func parseLine(values string, part int) int {
	// Find all the bracketed terms, and resolve them
	// Evaluate all the bracketed elements
	r, _ := regexp.Compile(`\((.*?)\)`)
	subStrings := r.FindAllString(values, -1)
	if len(subStrings) == 0 {
		if part == 1 {
			return evalP1(values)
		} else {
			return evalP2(values)
		}
	}

	for _, subString := range subStrings {
		lastIndex := strings.LastIndex(subString, "(")
		subValue := subString[lastIndex:len(subString)]
		var subSolution int
		if part == 1 {
			subSolution = evalP1(subValue)
		} else {
			subSolution = evalP2(subValue)
		}
		values = strings.Replace(values, subValue, strconv.Itoa(subSolution), 1)
	}
	return parseLine(values, part)
}

func q18part1(data []string) {
	totalValues := 0
	for _, row := range data {
		solution := parseLine(row, 1)
		totalValues += solution
	}
	fmt.Println("Day 18 Part 1 Solution:")
	fmt.Println(totalValues)
}

func q18part2(data []string) {
	totalValues := 0
	for _, row := range data {
		solution := parseLine(row, 2)
		totalValues += solution
	}
	fmt.Println("Day 18 Part 2 Solution:")
	fmt.Println(totalValues)
}
