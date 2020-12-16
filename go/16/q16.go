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

	data := lib.ReadInputAsStr(16)
	q16part1(data)
	q16part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

type Rule struct {
	name string
	min1 int
	max1 int
	min2 int
	max2 int
}

func (r Rule) Validate(number int) bool {
	//fmt.Printf("%d < %d < %d\n", r.min, number, r.max)
	if (number < r.min1 || number > r.max1) && (number < r.min2 || number > r.max2) {
		return false
	}
	return true
}

type RuleTester struct {
	rules []Rule
}

func (rt RuleTester) Validate(number int) RuleTester {
	var validRules []Rule
	for _, r := range rt.rules {
		if r.Validate(number) == true {
			validRules = append(validRules, r)
		}
	}
	rt.rules = validRules
	return rt
}

func ParseRule(ruleString string) Rule {
	split := strings.Split(ruleString, ":")
	name := split[0]
	values := strings.Split(split[1], "or")
	first := strings.Split(values[0], "-")
	second := strings.Split(values[1], "-")
	rule1Min, _ := strconv.Atoi(strings.TrimSpace(first[0]))
	rule1Max, _ := strconv.Atoi(strings.TrimSpace(first[1]))
	rule2Min, _ := strconv.Atoi(strings.TrimSpace(second[0]))
	rule2Max, _ := strconv.Atoi(strings.TrimSpace(second[1]))
	return Rule{
		name: name,
		min1: rule1Min,
		max1: rule1Max,
		min2: rule2Min,
		max2: rule2Max,
	}
}

func ParseTicket(ticketString string) []int {
	numbers := strings.Split(ticketString, ",")
	var ints []int
	for _, number := range numbers {
		ticketValue, _ := strconv.Atoi(number)
		ints = append(ints, ticketValue)
	}
	return ints
}

// Returns Rules - Your Ticket, All Tickets
func ParseInput(input []string) ([]Rule, []int, [][]int) {
	stage := 0
	var rules []Rule
	var yourTicket []int
	var nearbyTickets [][]int
	for _, row := range input {
		if row == "" {
			continue
		}
		if row == "your ticket:" {
			stage = 1
			continue
		}
		if row == "nearby tickets:" {
			stage = 2
			continue
		}

		if stage == 0 {
			newRule := ParseRule(row)
			rules = append(rules, newRule)
		} else if stage == 1 {
			yourTicket = ParseTicket(row)
		} else if stage == 2 {
			nearbyTickets = append(nearbyTickets, ParseTicket(row))
		}
	}
	return rules, yourTicket, nearbyTickets
}

func q16part1(data []string) {
	rules, _, nearbyTickets := ParseInput(data)
	var invalidValues int
	for _, ticket := range nearbyTickets {
		for _, ticketValue := range ticket {
			isValid := false
			for _, rule := range rules {
				if rule.Validate(ticketValue) == true {
					isValid = true
					break
				}
			}
			if isValid == false {
				invalidValues += ticketValue
			}
		}
	}
	fmt.Println("Day 16 Part 1 Solution: ")
	fmt.Println(invalidValues)
}

func q16part2(data []string) {
	rules, yourTicket, nearbyTickets := ParseInput(data)
	var validTickets [][]int
	for _, ticket := range nearbyTickets {
		ticketValid := true
		for _, ticketValue := range ticket {
			isValid := false
			for _, rule := range rules {
				if rule.Validate(ticketValue) == true {
					isValid = true
					break
				}
			}
			if isValid == false {
				ticketValid = false
				break
			}
		}
		if ticketValid == true {
			validTickets = append(validTickets, ticket)
		}
	}
	fmt.Printf("Removed %d/%d Tickets\n", len(nearbyTickets)-len(validTickets), len(nearbyTickets))

	rulesOrder := make([]Rule, len(validTickets[0]))
	var departures []int

	// Iterate over lists multiple times, because sometimes two rules might be valid for one index
	for len(departures) < 6 {
		for i := 0; i < len(validTickets[0]); i++ {
			// Check if we've already calculated this index
			if rulesOrder[i].min1 != 0 {
				continue
			}
			rt := RuleTester{
				rules: rules,
			}
			// Include yourTicket
			rt = rt.Validate(yourTicket[i])

			// Rule Tester will drop rules that fail
			for _, ticket := range validTickets {
				rt = rt.Validate(ticket[i])
				if len(rt.rules) == 1 {
					correctRule := rt.rules[0]
					if strings.HasPrefix(correctRule.name, "departure") {
						departures = append(departures, i)
					}
					rulesOrder[i] = rt.rules[0]

					// Remove from the global rules
					for k, rule := range rules {
						if rule.name == rt.rules[0].name {
							rules = append(rules[0:k], rules[k+1:len(rules)]...)
							break
						}
					}
					break
				}
			}
		}
	}
	solution := 1
	for _, index := range departures {
		solution *= yourTicket[index]
	}

	fmt.Println("Day 16 Part 1 Solution: ")
	fmt.Println(solution)
}
