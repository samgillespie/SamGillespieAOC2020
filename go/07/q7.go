package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	lib "../lib"
)

func main() {
	data := lib.ReadInputAsStr(7)
	start := time.Now()
	q7part1(data)
	q7part2(data)
	elapsed := time.Since(start)
	fmt.Printf("Main took %s", elapsed)
}

type bag struct {
	name     string
	contents map[*bag]int
}

func AddRowToBagMap(bag_map map[string]*bag, row string) map[string]*bag {
	// Generates a tree like bag_map, that we can navigate to get our answers
	clean := strings.ReplaceAll(row, "bags", "bag")
	clean = strings.ReplaceAll(clean, ".", "")
	split := strings.Split(clean, "contain")
	bag_name := strings.TrimSpace(split[0])
	if split[1] == "no other bags." {
		return bag_map
	}

	contents := strings.Split(split[1], ",")
	contents_map := make(map[*bag]int, len(contents))
	for _, item := range contents {
		item_split := strings.Split(item, " ")
		item_number, _ := strconv.Atoi(item_split[1])
		item_name := strings.TrimSpace(strings.Join(item_split[2:], " "))
		// Get the bag from the map, create it if it doesn't exist yet
		if bag_map[item_name] == nil {
			temp_bag := bag{name: item_name, contents: make(map[*bag]int, 0)}
			bag_map[item_name] = &temp_bag
			contents_map[&temp_bag] = item_number
		} else {
			temp_bag := bag_map[item_name]
			contents_map[temp_bag] = item_number
		}
	}
	row_bag := bag{name: bag_name, contents: contents_map}

	bag_map[bag_name] = &row_bag
	return bag_map
}

func BagInList(bag_list []*bag, bag *bag) bool {
	for _, b := range bag_list {
		if b == bag {
			return true
		}
	}
	return false
}

func FindInMap(bag_map map[string]*bag, starting_bag_name string, target_name string) bool {
	bags_to_search := make([]*bag, 0)
	bags_searched := make([]*bag, 0)
	starting_bag := bag_map[starting_bag_name]

	for bag, _ := range starting_bag.contents {
		if bag.name == target_name {
			return true
		}
		bags_to_search = append(bags_to_search, bag)
	}

	for {
		if len(bags_to_search) == 0 {
			return false
		}

		active_bag := bag_map[bags_to_search[0].name]
		if active_bag.name == target_name {
			return true
		}
		for bag := range active_bag.contents {
			if BagInList(bags_searched, bag) == false {
				bags_to_search = append(bags_to_search, bag)
			}
		}
		bags_to_search = bags_to_search[1:]
	}
}

func q7part1(data []string) {
	bag_map := make(map[string]*bag)
	for _, row := range data {
		bag_map = AddRowToBagMap(bag_map, row)
	}

	counter := 0
	for bag_name, _ := range bag_map {
		search := FindInMap(bag_map, bag_name, "shiny gold bag")
		if search == true {

			counter += 1
		}
	}
	fmt.Println("Question 7 Part 1 Solution:")
	fmt.Println(counter)
}

func q7part2(data []string) {
	bag_map := make(map[string]*bag)

	for _, row := range data {
		bag_map = AddRowToBagMap(bag_map, row)
	}

	bags_to_count := map[*bag]int{bag_map["shiny gold bag"]: 1}
	counter := 0
	for {
		if len(bags_to_count) == 0 {
			// Need to subtract 1 for the shiny gold bag
			fmt.Println("Question 7 Part 2 Solution:")
			fmt.Println(counter - 1)
			return
		}

		for bag, number := range bags_to_count {
			for content, inner_number := range bag_map[bag.name].contents {
				bags_to_count[content] += inner_number * number
			}
			counter += number
			delete(bags_to_count, bag)
		}
	}
}
