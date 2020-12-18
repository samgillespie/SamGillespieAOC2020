package main

import (
	"fmt"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(17)
	q17part1(data)
	q17part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

type coord struct {
	x int
	y int
	z int
	w int
}

func (c coord) adjacentElems(dimensions int) []coord {
	adjacent := make([]coord, 0)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if dimensions == 4 {
					for w := -1; w <= 1; w++ {
						if x == 0 && y == 0 && z == 0 && w == 0 {
							continue
						}
						adjacent = append(adjacent, coord{x: c.x + x, y: c.y + y, z: c.z + z, w: c.w + w})
					}
				} else {
					if x == 0 && y == 0 && z == 0 {
						continue
					}
					adjacent = append(adjacent, coord{x: c.x + x, y: c.y + y, z: c.z + z, w: 0})
				}
			}
		}
	}
	return adjacent
}

func (c coord) adjacentActive(dataMap map[coord]bool, dimensions int) int {
	elems := c.adjacentElems(dimensions)
	active := 0
	for _, elem := range elems {
		if dataMap[elem] == true {
			active++
		}
	}
	return active
}

func calculateBoundaries(mapData map[coord]bool) (int, int, int, int, int, int, int, int) {
	xMin := 0
	xMax := 0
	yMin := 0
	yMax := 0
	zMin := 0
	zMax := 0
	wMin := 0
	wMax := 0
	for index := range mapData {
		if index.x < xMin {
			xMin = index.x
		}
		if index.x > xMax {
			xMax = index.x
		}
		if index.y < yMin {
			yMin = index.y
		}
		if index.y > yMax {
			yMax = index.y
		}
		if index.z < zMin {
			zMin = index.z
		}
		if index.z > zMax {
			zMax = index.z
		}
		if index.w < wMin {
			wMin = index.w
		}
		if index.w > wMax {
			wMax = index.w
		}
	}
	xMin--
	yMin--
	zMin--
	wMin--
	xMax++
	yMax++
	zMax++
	wMax++
	return xMin, xMax, yMin, yMax, zMin, zMax, wMin, wMax
}

func parseInput(input []string) map[coord]bool {
	mapData := make(map[coord]bool)
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			coordinates := coord{x: x, y: y, z: 0, w: 0}
			mapData[coordinates] = input[y][x] == '#'
		}
	}
	return mapData
}

func printMap(mapData map[coord]bool) []string {
	visualMap := make([]string, 0)
	xMin, xMax, yMin, yMax, zMin, zMax, _, _ := calculateBoundaries(mapData)
	for z := zMin + 1; z <= zMax-1; z++ {
		layer := "\n"
		for y := yMin + 1; y <= yMax-1; y++ {
			for x := xMin + 1; x <= xMax-1; x++ {
				if mapData[coord{x: x, y: y, z: z, w: 0}] == true {
					layer += "#"
				} else {
					layer += "."
				}
			}
			layer += "\n"
		}
		visualMap = append(visualMap, layer)
		fmt.Printf("Layer %d", z)
		fmt.Println(layer)
	}
	return visualMap
}

func simulateStep(mapData map[coord]bool, dimensions int) map[coord]bool {
	nextStep := make(map[coord]bool)
	xMin, xMax, yMin, yMax, zMin, zMax, wMin, wMax := calculateBoundaries(mapData)
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			for z := zMin; z <= zMax; z++ {
				for w := wMin; w <= wMax; w++ {
					if dimensions == 3 && w != 0 {
						continue
					}
					coordinate := coord{x: x, y: y, z: z, w: w}
					currentActive := mapData[coordinate]
					adjacentCells := coordinate.adjacentActive(mapData, dimensions)
					if currentActive == true {
						if adjacentCells == 2 || adjacentCells == 3 {
							nextStep[coordinate] = true
						} else {
							nextStep[coordinate] = false
						}
					} else {
						if adjacentCells == 3 {
							nextStep[coordinate] = true
						} else {
							nextStep[coordinate] = false
						}
					}
				}

			}
		}
	}
	return nextStep
}

func q17part1(data []string) {
	mapData := parseInput(data)
	// Simulate a step
	for step := 0; step < 6; step++ {
		mapData = simulateStep(mapData, 3)
	}
	actives := 0
	for _, value := range mapData {
		if value == true {
			actives++
		}
	}
	fmt.Println("Day 17 Part 1 Solution:")
	fmt.Println(actives)

}

func q17part2(data []string) {
	mapData := parseInput(data)
	// Simulate a step
	for step := 0; step < 6; step++ {
		mapData = simulateStep(mapData, 4)
	}
	actives := 0
	for _, value := range mapData {
		if value == true {
			actives++
		}
	}
	fmt.Println("Day 17 Part 2 Solution:")
	fmt.Println(actives)
}
