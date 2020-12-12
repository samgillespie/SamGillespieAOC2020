package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(11)
	q11part1(data)
	q11part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func raycast(raycastMap []string, x int, y int, direction string, distance int) byte {
	var change []int
	if direction == "N" {
		change = []int{0, -1}
	} else if direction == "NE" {
		change = []int{1, -1}
	} else if direction == "E" {
		change = []int{1, 0}
	} else if direction == "SE" {
		change = []int{1, 1}
	} else if direction == "S" {
		change = []int{0, 1}
	} else if direction == "SW" {
		change = []int{-1, 1}
	} else if direction == "W" {
		change = []int{-1, 0}
	} else if direction == "NW" {
		change = []int{-1, -1}
	}

	ray_dist := 0
	currPos := []int{x, y}
	for {
		currPos = []int{currPos[0] + change[0], currPos[1] + change[1]}
		// Check in bounds
		if currPos[1] < 0 || currPos[1] >= len(raycastMap) {
			return '.'
		}
		if currPos[0] < 0 || currPos[0] >= len(raycastMap[0]) {
			return '.'
		}

		if raycastMap[currPos[1]][currPos[0]] != '.' {
			return raycastMap[currPos[1]][currPos[0]]
		}

		ray_dist += 1
		if ray_dist == distance {
			return '.'
		}

	}
}

func calculateAdjacency(adjMap []string, x int, y int, x_max int, y_max int, distance int) int {
	adjacent := 0
	directions := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	for _, direction := range directions {
		if raycast(adjMap, x, y, direction, distance) == '#' {
			adjacent += 1
		}
	}
	return adjacent
}

func simulate_step(currConfig []string, distance int, adjacency_limit int) []string {
	nextStep := make([]string, len(currConfig))
	y_max := len(currConfig)
	x_max := len(currConfig[0])
	for y := 0; y < y_max; y++ {
		row := ""
		for x := 0; x < x_max; x++ {
			if currConfig[y][x] == '.' {
				row += "."
				continue
			}
			adj := calculateAdjacency(currConfig, x, y, x_max, y_max, distance)
			//fmt.Printf("%d, %d, %d \n", x, y, adj)

			if currConfig[y][x] == 'L' && adj == 0 {
				row += "#"
			} else if currConfig[y][x] == 'L' && adj > 0 {
				row += "L"
			} else if currConfig[y][x] == '#' && adj >= adjacency_limit {
				row += "L"
			} else if currConfig[y][x] == '#' {
				row += "#"
			}
		}
		nextStep[y] = row
	}

	return nextStep
}

func q11part1(data []string) {
	prevCount := -1
	nextStep := make([]string, len(data))
	copy(nextStep, data)
	for {
		nextStep = simulate_step(nextStep, 1, 4)
		currCount := 0
		for y := range nextStep {
			currCount += strings.Count(nextStep[y], "#")
		}
		if currCount == prevCount {
			fmt.Println("Day 11 Part 1 Answer:")
			fmt.Println(currCount)
			break
		}
		prevCount = currCount
	}
}

func printMap(mapInfo []string) {
	// For visualizing
	printable := strings.Join(mapInfo, "\n")
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Printf(printable)
	fmt.Printf("\n\n")
	time.Sleep(100 * time.Millisecond)
}

func q11part2(data []string) {
	prevCount := -1
	nextStep := make([]string, len(data))
	copy(nextStep, data)
	for {
		nextStep = simulate_step(nextStep, 9999, 5)
		currCount := 0
		for y := range nextStep {
			currCount += strings.Count(nextStep[y], "#")
		}
		if currCount == prevCount {
			fmt.Println("Day 11 Part 2 Answer:")
			fmt.Println(currCount)
			break
		}
		prevCount = currCount
	}
}
