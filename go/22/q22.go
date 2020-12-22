package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()
	data := lib.ReadInputAsStr(22)
	q22part1(data)
	q22part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func parseInput(data []string) ([]int, []int) {
	deck1 := []int{}
	deck2 := []int{}
	currentPlayer := 0
	for _, contents := range data {
		if contents == "" {
			continue
		}
		if strings.Contains(contents, "Player") {
			currentPlayer++
			continue
		}

		cardNum, _ := strconv.Atoi(contents)
		if currentPlayer == 1 {
			deck1 = append(deck1, cardNum)
		} else {
			deck2 = append(deck2, cardNum)
		}
	}
	return deck1, deck2
}

// Part 1
func CrabRound(deck1 []int, deck2 []int) ([]int, []int) {
	var card1, card2 int
	card1, deck1 = deck1[0], deck1[1:]
	card2, deck2 = deck2[0], deck2[1:]
	if card1 > card2 {
		deck1 = append(deck1, card1, card2)
	}
	if card2 > card1 {
		deck2 = append(deck2, card2, card1)
	}
	return deck1, deck2
}

func CrabBattle(deck1 []int, deck2 []int) {
	for len(deck1) > 0 && len(deck2) > 0 {
		deck1, deck2 = CrabRound(deck1, deck2)
	}
	var score int
	if len(deck1) == 0 {
		score = CrabScore(deck2)
	} else {
		score = CrabScore(deck1)
	}
	fmt.Println("Crab Battle Score Part 1")
	fmt.Println(score)
}

func CrabScore(deck []int) int {
	score := 0
	for i, card := range deck {
		score += card * (len(deck) - i)
	}
	return score
}

// Part 2

func DeckToInteger(deck []int) int64 {
	integer := int64(0)
	for _, card := range deck {
		integer += int64(math.Pow(2, float64(card)))
	}
	return integer
}

func SaveState(deck1 []int, deck2 []int) []int64 {
	return []int64{DeckToInteger(deck1), DeckToInteger(deck2)}
}

func CompareState(states [][]int64, deck1 []int, deck2 []int) bool {
	for _, state := range states {
		if DeckToInteger(deck1) != state[0] {
			continue
		}
		if DeckToInteger(deck2) != state[1] {
			continue
		}
		return true
	}
	return false
}

func RecursiveCrabRound(deck1 []int, deck2 []int) ([]int, []int) {
	var card1, card2 int
	card1, deck1 = deck1[0], deck1[1:]
	card2, deck2 = deck2[0], deck2[1:]

	winner := 0
	if card1 <= len(deck1) && card2 <= len(deck2) {

		newDeck1 := make([]int, card1)
		newDeck2 := make([]int, card2)
		copy(newDeck1, deck1)
		copy(newDeck2, deck2)
		winner, _, _ = RecursiveCrabBattle(newDeck1, newDeck2)

		if winner == 1 {
			deck1 = append(deck1, card1, card2)
		} else {
			deck2 = append(deck2, card2, card1)
		}
		return deck1, deck2
	}

	if card1 > card2 {
		deck1 = append(deck1, card1, card2)
	}
	if card2 > card1 {
		deck2 = append(deck2, card2, card1)
	}
	return deck1, deck2
}

func RecursiveCrabBattle(deck1 []int, deck2 []int) (int, []int, []int) { // Bool represents if Player 1 won
	previousStates := [][]int64{}
	for len(deck1) > 0 && len(deck2) > 0 {
		if CompareState(previousStates, deck1, deck2) == true {
			return 1, deck1, deck2
		}

		previousStates = append(previousStates, SaveState(deck1, deck2))
		deck1, deck2 = RecursiveCrabRound(deck1, deck2)
	}

	if len(deck1) == 0 {
		return 2, deck1, deck2
	} else {
		return 1, deck1, deck2
	}
}

func q22part1(data []string) {
	deck1, deck2 := parseInput(data)
	CrabBattle(deck1, deck2)
}

func q22part2(data []string) {
	deck1, deck2 := parseInput(data)
	winner, deck1, deck2 := RecursiveCrabBattle(deck1, deck2)
	fmt.Println("Crab Battle Score Part 2")
	if winner == 1 {
		fmt.Println(CrabScore(deck1))
	} else if winner == 2 {
		fmt.Println(CrabScore(deck2))
	}
}
