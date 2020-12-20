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

	data := lib.ReadInputAsStr(20)
	SolveDay20(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

// Make sure a tile is facing the right way before assigning tiles
type Tile struct {
	id            int
	tile          []string
	adjacentTiles []int // Holds tiles that are currently unaligned
	posx          int
	posy          int
}

type coord struct {
	x int
	y int
}

func (t Tile) Rotate() Tile {
	tileRotation := make([]string, 0)
	currentRow := ""
	for i, _ := range t.tile {
		for _, row := range t.tile {
			currentRow += string(row[i])
		}
		tileRotation = append(tileRotation, currentRow)
		currentRow = ""
	}
	t.tile = tileRotation
	t = t.FlipHorizontal()
	return t
}

func (t Tile) FlipHorizontal() Tile {
	tileFlip := make([]string, 0)
	for _, row := range t.tile {
		currentRow := ""
		for j := len(row) - 1; j >= 0; j-- {
			currentRow += string(row[j])
		}
		tileFlip = append(tileFlip, currentRow)
		currentRow = ""
	}
	t.tile = tileFlip
	return t
}

func (t Tile) FlipVertical() Tile {
	t = t.Rotate()
	t = t.FlipHorizontal()
	t = t.Rotate()
	t = t.Rotate()
	t = t.Rotate()
	return t
}

// Return the rows
func (t Tile) TopRow() string {
	return t.tile[0]
}

func (t Tile) BottomRow() string {
	return t.tile[len(t.tile)-1]
}

func (t Tile) LeftRow() string {
	res := ""
	for _, row := range t.tile {
		res += string(row[0])
	}
	return res
}

func (t Tile) RightRow() string {
	res := ""
	for _, row := range t.tile {
		res += string(row[len(row)-1])
	}
	return res
}

func (t Tile) Sides() []string {
	top := t.TopRow()
	bottom := t.BottomRow()
	left := t.LeftRow()
	right := t.RightRow()
	return []string{
		top,
		bottom,
		left,
		right,
		ReverseString(top),
		ReverseString(bottom),
		ReverseString(left),
		ReverseString(right),
	}
}

func parseInput(inputData []string) []Tile {
	tile := Tile{}
	tile.posx = -1
	tile.posy = -1
	tileData := make([]string, 0)
	tileList := make([]Tile, 0)
	for _, row := range inputData {
		if row == "" || row == " " {
			tile.tile = tileData
			tileList = append(tileList, tile)
			tile = Tile{}
			tile.posx = -1
			tile.posy = -1
			tileData = make([]string, 0)
		} else if row[0:4] == "Tile" {
			idString := strings.Split(row, " ")[1]
			idString = strings.Replace(idString, ":", "", 1)
			tile.id, _ = strconv.Atoi(idString)
			continue
		} else {
			tileData = append(tileData, row)
		}
	}
	return tileList
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FindMatch(a []string, b []string) (bool, string, string) {
	orientations := []string{
		"top",
		"bottom",
		"left",
		"right",
		"reversetop",
		"reversebottom",
		"reverseleft",
		"reverseright",
	}

	for indexi, i := range a {
		for indexj, j := range b {
			if i == j {
				return true, orientations[indexi], orientations[indexj]
			}
		}
	}
	return false, "", ""
}

func isAligned(a string, b string) bool {
	if a == "left" && b == "right" {
		return true
	} else if a == "right" && b == "left" {
		return true
	} else if a == "top" && b == "bottom" {
		return true
	} else if a == "bottom" && b == "top" {
		return true
	}
	return false
}

func FindRelations(data []string) map[int]Tile {
	tiles := parseInput(data)
	p1Solution := 1
	cornerTile := Tile{}
	for i, tile := range tiles {
		sides := tile.Sides()
		hits := 0
		adjacentTiles := make([]int, 0)
		for _, checkTile := range tiles {
			if tile.id == checkTile.id {
				continue
			}
			checkSides := checkTile.Sides()
			match, _, _ := FindMatch(sides, checkSides)
			if match == true {
				hits += 1
				adjacentTiles = append(adjacentTiles, checkTile.id)
			}
		}
		if hits == 2 {
			p1Solution *= tile.id
			cornerTile = tile
		}
		tiles[i].adjacentTiles = adjacentTiles
	}
	fmt.Println("Day 20 Part 1 Solution")
	fmt.Println(p1Solution)
	return orientTiles(tiles, cornerTile.id)
}

func ListToMap(tiles []Tile) map[int]Tile {
	mapping := make(map[int]Tile)
	for _, tile := range tiles {
		mapping[tile.id] = tile
	}
	return mapping
}

func orientTiles(tiles []Tile, startingTileId int) map[int]Tile {
	// Iterate through all the tiles in startingTile
	// Start by declaring the first tile as the Top Left tile
	tileMap := ListToMap(tiles)
	startingTile := tileMap[startingTileId]
	// Rotate the starting element to topleft, and adjust everything accordingly
	for {

		firstTile := tileMap[startingTile.adjacentTiles[0]]
		secondTile := tileMap[startingTile.adjacentTiles[1]]
		_, firstOrientation, _ := FindMatch(startingTile.Sides(), firstTile.Sides())
		_, secondOrientation, _ := FindMatch(startingTile.Sides(), secondTile.Sides())
		if firstOrientation == "reversetop" || firstOrientation == "reversebottom" || secondOrientation == "reversetop" || secondOrientation == "reversebottom" {
			startingTile = startingTile.FlipHorizontal()
			continue
		}
		if firstOrientation == "reverseleft" || firstOrientation == "reverseright" || secondOrientation == "reverseleft" || secondOrientation == "reverseright" {
			startingTile = startingTile.FlipVertical()
			continue
		}
		if (firstOrientation == "right" && secondOrientation == "bottom") || (firstOrientation == "bottom" && secondOrientation == "right") {
			break
		} else {
			startingTile = startingTile.Rotate()
		}
	}
	startingTile.posx = 0
	startingTile.posy = 0
	tileMap[startingTileId] = startingTile

	// Hold all tiles that have been placed, but do not have their adjacents in place
	unprocessedTiles := []int{startingTileId}

	for len(unprocessedTiles) > 0 {
		baseTile := tileMap[unprocessedTiles[0]]
		baseTileSides := baseTile.Sides()
		for _, tileIndex := range baseTile.adjacentTiles {
			adjacentTile := tileMap[tileIndex]
			if adjacentTile.posx != -1 {

				continue
			}
			// Flip and Rotate till Opposed (i.e. Left and Right)
			for {
				_, baseAngle, otherAngle := FindMatch(baseTileSides, adjacentTile.Sides())
				if otherAngle == "reversetop" || otherAngle == "reversebottom" {
					adjacentTile = adjacentTile.FlipHorizontal()
					continue
				}
				if otherAngle == "reverseleft" || otherAngle == "reverseright" {
					adjacentTile = adjacentTile.FlipVertical()
					continue
				}
				if isAligned(baseAngle, otherAngle) == false {
					adjacentTile = adjacentTile.Rotate()
					continue
				}
				if baseAngle == "right" {
					adjacentTile.posx = baseTile.posx + 1
					adjacentTile.posy = baseTile.posy
				} else if baseAngle == "bottom" {
					adjacentTile.posx = baseTile.posx
					adjacentTile.posy = baseTile.posy + 1
				} else if baseAngle == "top" {
					adjacentTile.posx = baseTile.posx
					adjacentTile.posy = baseTile.posy - 1
				} else if baseAngle == "left" {
					adjacentTile.posx = baseTile.posx - 1
					adjacentTile.posy = baseTile.posy
				}
				tileMap[adjacentTile.id] = adjacentTile
				unprocessedTiles = append(unprocessedTiles, adjacentTile.id)
				break
			}
		}
		unprocessedTiles = unprocessedTiles[1:]
	}
	return tileMap
}

func ConvertToCoordMap(mapData map[int]Tile) (map[coord]Tile, int, int) {
	newMap := make(map[coord]Tile)
	xmax := 0
	ymax := 0
	for _, tile := range mapData {
		c := coord{x: tile.posx, y: tile.posy}
		newMap[c] = tile
		if tile.posx > xmax {
			xmax = tile.posx
		}
		if tile.posy > ymax {
			ymax = tile.posy
		}
	}
	return newMap, xmax, ymax
}

func CombineMap(mapData map[int]Tile) []string {
	newMap, xmax, ymax := ConvertToCoordMap(mapData)

	globalMap := make([]string, 0)

	for y := 0; y <= ymax; y++ {
		cursors := make([]string, 0)
		// Prepare Cursors
		for i := 0; i < 8; i++ { // 1 -> 9 to shave off the edges
			cursors = append(cursors, "")
		}
		for x := 0; x <= xmax; x++ {
			coord := coord{x: x, y: y}
			tile := newMap[coord]
			for i := 0; i < 8; i++ {
				target := tile.tile[i+1]
				cursors[i] += target[1 : len(target)-1]
			}
		}
		globalMap = append(globalMap, cursors...)
	}
	return globalMap
}

func dragonCoordinates(x int, y int) []coord {
	// There's probably a better way
	return []coord{
		coord{x: 0 + x, y: 1 + y},
		coord{x: 1 + x, y: 2 + y},
		coord{x: 4 + x, y: 2 + y},
		coord{x: 5 + x, y: 1 + y},
		coord{x: 6 + x, y: 1 + y},
		coord{x: 7 + x, y: 2 + y},
		coord{x: 10 + x, y: 2 + y},
		coord{x: 11 + x, y: 1 + y},
		coord{x: 12 + x, y: 1 + y},
		coord{x: 13 + x, y: 2 + y},
		coord{x: 16 + x, y: 2 + y},
		coord{x: 17 + x, y: 1 + y},
		coord{x: 18 + x, y: 1 + y},
		coord{x: 19 + x, y: 1 + y},
		coord{x: 18 + x, y: 0 + y},
	}
}

func findDragons(data []string) []string {
	// We can find dragons by looking for heads
	dragonLength := 20
	for y := 0; y < len(data)-2; y++ {
		for x := 0; x < len(data[y])-dragonLength+1; x++ {
			//fmt.Printf("%d, %d\n", x, y)
			coords := dragonCoordinates(x, y)
			success := true
			for _, coord := range coords {
				if data[coord.y][coord.x] == '.' {
					success = false
					break
				}
			}

			if success == true {
				for _, coord := range coords {
					data[coord.y] = data[coord.y][0:coord.x] + "O" + data[coord.y][coord.x+1:]
				}
			}
		}
	}
	return data
}

func SolveDay20(data []string) {
	tileMap := FindRelations(data) // Calculate Part 1 in this step
	globalMap := CombineMap(tileMap)

	// For testing - Rotate to match the example doc
	bigTile := Tile{tile: globalMap}

	globalMap = bigTile.tile

	bigTile.tile = findDragons(bigTile.tile)
	bigTile = bigTile.Rotate()
	bigTile.tile = findDragons(bigTile.tile)
	bigTile = bigTile.Rotate()
	bigTile.tile = findDragons(bigTile.tile)
	bigTile = bigTile.Rotate()
	bigTile.tile = findDragons(bigTile.tile)
	bigTile = bigTile.FlipHorizontal()
	bigTile.tile = findDragons(bigTile.tile)
	bigTile = bigTile.Rotate()
	bigTile.tile = findDragons(bigTile.tile)
	bigTile = bigTile.Rotate()
	bigTile.tile = findDragons(bigTile.tile)
	bigTile = bigTile.Rotate()
	bigTile.tile = findDragons(bigTile.tile)
	counter := 0
	for _, row := range bigTile.tile {
		counter += strings.Count(row, "#")
	}
	fmt.Println("Day 20 Part 2 Solution")
	fmt.Println(counter)
}
