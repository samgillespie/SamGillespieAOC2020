package main

import (
	"container/ring"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	//data := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	data := []int{9, 7, 4, 6, 1, 8, 3, 5, 2}
	q23part1(data)
	q23part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func intToRing(cups []int) (*ring.Ring, map[int]*ring.Ring) {
	newRing := ring.New(len(cups))
	pointerMap := make(map[int]*ring.Ring)
	for i := 0; i < len(cups); i++ {
		newRing.Value = cups[i]
		pointerMap[cups[i]] = newRing
		newRing = newRing.Next()
	}
	return newRing, pointerMap
}

func PrintRing(r *ring.Ring) {
	r.Do(func(p interface{}) {
		fmt.Printf("%d", p.(int))
	})
	fmt.Println("")
}

func loopIndex(i int, numCups int) int {
	if i >= numCups {
		return i - numCups
	} else if i < 0 {
		return i + numCups
	}
	return i
}

func loopValue(i int, numCups int) int {
	if i > numCups {
		return i - numCups
	} else if i <= 0 {
		return i + numCups
	}
	return i
}

func SelectCup(selectedCup int, cupNum int, nextElems []int) int {
	for {
		selectedCup--
		if selectedCup <= 0 {
			selectedCup = cupNum
		}
		if nextElems[0] == selectedCup || nextElems[1] == selectedCup || nextElems[2] == selectedCup {
			continue
		}
		return selectedCup
	}
}

func CupFinder(r *ring.Ring, cupNumber int) *ring.Ring {
	for {
		if r.Value.(int) == cupNumber {
			return r
		}
		r = r.Next()
	}
}

func printSolution(r *ring.Ring) {
	start := CupFinder(r, 1)
	start.Do(func(p interface{}) {
		if p.(int) != 1 {
			fmt.Printf("%d", p.(int))
		}
	})
	fmt.Println("")
}

func ResolveStep(cursor *ring.Ring, pointers map[int]*ring.Ring, length int) {

	nextElems := []int{}
	a := cursor

	a = a.Next()
	nextElems = append(nextElems, a.Value.(int))
	a = a.Next()
	nextElems = append(nextElems, a.Value.(int))
	a = a.Next()
	nextElems = append(nextElems, a.Value.(int))

	selectedCup := cursor.Value.(int)
	selectedCup = SelectCup(selectedCup, length, nextElems)
	targetCup := pointers[selectedCup]

	remove3 := cursor.Unlink(3)
	targetCup.Link(remove3)
}

func q23part1(cups []int) {
	// Part 1
	cursor, pointers := intToRing(cups)
	for i := 0; i < 100; i++ {
		ResolveStep(cursor, pointers, 9)
		cursor = cursor.Next()
	}
	fmt.Println("Day 23 Part 1 Solution:")
	printSolution(cursor)

}

func q23part2(cups []int) {
	for i := 10; i <= 1000000; i++ {
		cups = append(cups, i)

	}
	cursor, pointers := intToRing(cups)

	for i := 1; i < 10000000; i++ {
		ResolveStep(cursor, pointers, 1000000)
		cursor = cursor.Next()
	}
	cup := pointers[1]
	cup = cup.Next()
	a := cup.Value.(int)
	cup = cup.Next()
	b := cup.Value.(int)
	fmt.Println("Day 23 Part 2 Solution:")
	fmt.Println(a * b)
}
