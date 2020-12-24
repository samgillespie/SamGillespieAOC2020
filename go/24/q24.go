package main

import (
	"fmt"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(24)
	hexagonMap := CreateStartingGrid(data)
	GameOfLife(hexagonMap)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

type Hexagon struct {
	isBlack bool
	x       int
	y       int
	toFlip  bool
}

type Coords struct {
	x int
	y int
}

func (h Hexagon) Coords() Coords {
	return Coords{x: h.x, y: h.y}
}

func (h Hexagon) DirectionCoords(direction string) Coords {
	// Define a hexagon as size 2.  Moving cardinal is +2, moving angled is +1 +1
	switch direction {
	case "ne":
		return Coords{h.x + 1, h.y + 1}
	case "e":
		return Coords{h.x + 2, h.y}
	case "se":
		return Coords{h.x + 1, h.y - 1}
	case "nw":
		return Coords{h.x - 1, h.y + 1}
	case "w":
		return Coords{h.x - 2, h.y}
	case "sw":
		return Coords{h.x - 1, h.y - 1}
	}
	return Coords{}
}

func (h Hexagon) AdjacentBlacks(hexagonMap map[Coords]*Hexagon) ([]*Hexagon, int) {
	directions := []string{"ne", "e", "se", "nw", "w", "sw"}
	blacks := 0
	adjacentHexes := []*Hexagon{}
	for _, direction := range directions {
		i := h
		coords := i.DirectionCoords(direction)
		hex, instantiated := hexagonMap[coords]

		var adjacentHex *Hexagon
		if instantiated == true && hex.isBlack == true {
			blacks++
			adjacentHex = hexagonMap[coords]

		} else {
			adjacentHex = &Hexagon{x: coords.x, y: coords.y, isBlack: false}
		}

		adjacentHexes = append(adjacentHexes, adjacentHex)
	}
	return adjacentHexes, blacks

}

func (h *Hexagon) Move(direction string, hexagonMap map[Coords]*Hexagon) (*Hexagon, map[Coords]*Hexagon) {
	newPos := h.DirectionCoords(direction)
	//fmt.Printf("direction: %s, CurrentPos: %v NewPos: %v\n", direction, Coords{h.x, h.y}, newPos)
	targetHex, alreadyExists := hexagonMap[newPos]
	if alreadyExists == false {
		newHex := Hexagon{x: newPos.x, y: newPos.y}
		targetHex = &newHex
	}

	hexagonMap[newPos] = targetHex

	return targetHex, hexagonMap
}

func parseInput(data []string) [][]string {
	finalData := [][]string{}
	isNorS := ""
	for _, row := range data {
		instruction := []string{}
		for _, char := range row {
			if char == 'n' || char == 's' {
				isNorS = string(char)
			} else {
				instruction = append(instruction, string(isNorS)+string(char))
				isNorS = ""
			}
		}
		finalData = append(finalData, instruction)
	}
	return finalData
}

func CountBlack(hexagonMap map[Coords]*Hexagon) int {
	blackHexes := 0
	for _, hex := range hexagonMap {
		if hex.isBlack == true {
			blackHexes += 1
		}
	}
	return blackHexes
}

func CreateStartingGrid(data []string) map[Coords]*Hexagon {
	startingHexagon := Hexagon{x: 0, y: 0}
	instructions := parseInput(data)
	hexagonMap := make(map[Coords]*Hexagon)
	for _, instruction := range instructions {
		activeHexagon := &startingHexagon
		for _, direction := range instruction {
			activeHexagon, hexagonMap = activeHexagon.Move(direction, hexagonMap)
		}
		activeHexagon.isBlack = !activeHexagon.isBlack
	}
	blackHexes := CountBlack(hexagonMap)
	fmt.Println("Day 24 Part 1 Solution")
	fmt.Println(blackHexes)
	return hexagonMap
}

func SimulateStep(hexagonMap map[Coords]*Hexagon) map[Coords]*Hexagon {
	// Only resolve black cells, and white cells adjacent to black cells
	toResolve := make(map[*Hexagon]bool)
	for _, hex := range hexagonMap {
		if hex.isBlack == false {
			continue
		}
		toResolve[hex] = true
		adjacentCells, _ := hex.AdjacentBlacks(hexagonMap)
		for _, adjHex := range adjacentCells {
			toResolve[adjHex] = true
			hexagonMap[adjHex.Coords()] = adjHex
		}
	}

	toFlip := make(map[*Hexagon]bool)
	for hex := range toResolve {
		_, blackcells := hex.AdjacentBlacks(hexagonMap)
		if hex.isBlack && (blackcells == 0 || blackcells > 2) {
			hex.toFlip = true
			toFlip[hex] = true
		}

		if hex.isBlack == false && blackcells == 2 {
			hex.toFlip = true
			toFlip[hex] = true
		}
		if hex.x == 4 && hex.y == 0 {
		}
	}

	for hex := range toFlip {
		hex.isBlack = !hex.isBlack
		//fmt.Printf("%v Flipped to black=%t\n ", hex.Coords(), hex.isBlack)
		hex.toFlip = false
	}

	return hexagonMap
}

func GameOfLife(hexagonMap map[Coords]*Hexagon) {
	// Part 2
	for i := 0; i <= 100; i++ {
		fmt.Printf("Day %d: %d \n", i, CountBlack(hexagonMap))
		hexagonMap = SimulateStep(hexagonMap)

	}
}
