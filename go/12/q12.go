package main

import (
	"fmt"
	"strconv"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(12)
	q12part1(data)
	q12part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

type Ship struct {
	direction byte
	x         int
	y         int
}

type Waypoint struct {
	x int
	y int
}

func (waypoint Waypoint) Right(num int) Waypoint {
	for i := 0; i < num; i++ {
		temp := []int{waypoint.x, waypoint.y}
		waypoint.x = temp[1]
		waypoint.y = -temp[0]
	}
	return waypoint
}

func (waypoint Waypoint) Left(num int) Waypoint {
	for i := 0; i < num; i++ {
		temp := []int{waypoint.x, waypoint.y}
		waypoint.x = -temp[1]
		waypoint.y = temp[0]
	}
	return waypoint
}

func (ship Ship) Right(num int) Ship {
	for i := 0; i < num; i++ {
		if ship.direction == 'E' {
			ship.direction = 'S'
		} else if ship.direction == 'S' {
			ship.direction = 'W'
		} else if ship.direction == 'W' {
			ship.direction = 'N'
		} else if ship.direction == 'N' {
			ship.direction = 'E'
		}
	}
	return ship
}

func (ship Ship) Left(num int) Ship {
	for i := 0; i < num; i++ {
		if ship.direction == 'E' {
			ship.direction = 'N'
		} else if ship.direction == 'N' {
			ship.direction = 'W'
		} else if ship.direction == 'W' {
			ship.direction = 'S'
		} else if ship.direction == 'S' {
			ship.direction = 'E'
		}
	}
	return ship
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func (ship Ship) Manhatten() int {
	return abs(ship.x) + abs(ship.y)
}

func P1instruction(ship Ship, instruction string) Ship {
	value, _ := strconv.Atoi(instruction[1:])
	switch instruction[0] {
	case 'N':
		ship.y += value
	case 'S':
		ship.y -= value
	case 'E':
		ship.x += value
	case 'W':
		ship.x -= value
	case 'F':
		ship = P1instruction(ship, string(ship.direction)+instruction[1:])
	case 'R':
		ship = ship.Right(value / 90)
	case 'L':
		ship = ship.Left(value / 90)
	default:
		panic(fmt.Sprintf("Cannot find value %s", string(instruction[0])))
	}
	//fmt.Printf("%s, %d, %d\n", instruction, ship.x, ship.y)
	return ship
}

func P2instruction(ship Ship, waypoint Waypoint, instruction string) (Ship, Waypoint) {
	value, _ := strconv.Atoi(instruction[1:])
	switch instruction[0] {
	case 'N':
		waypoint.y += value
	case 'S':
		waypoint.y -= value
	case 'E':
		waypoint.x += value
	case 'W':
		waypoint.x -= value
	case 'F':
		ship.x += waypoint.x * value
		ship.y += waypoint.y * value
	case 'R':
		waypoint = waypoint.Right(value / 90)
	case 'L':
		waypoint = waypoint.Left(value / 90)
	default:
		panic(fmt.Sprintf("Cannot find value %s", string(instruction[0])))
	}
	//fmt.Printf("%s, %d, %d\n", instruction, ship.x, ship.y)
	return ship, waypoint
}

func q12part1(data []string) {

	ship := Ship{direction: 'E', x: 0, y: 0}
	for _, instruction := range data {
		ship = P1instruction(ship, instruction)
	}
	fmt.Println(ship.Manhatten())
}

func q12part2(data []string) {
	ship := Ship{direction: 'E', x: 0, y: 0}
	waypoint := Waypoint{x: 10, y: 1}
	for _, instruction := range data {
		ship, waypoint = P2instruction(ship, waypoint, instruction)
	}
	fmt.Println(ship.Manhatten())
}
