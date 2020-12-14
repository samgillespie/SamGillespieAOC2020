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

	data := lib.ReadInputAsStr(14)
	q14part1(data)
	q14part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func getMemIndex(str string) int64 {
	new := strings.Split(str, "[")
	value, _ := strconv.ParseInt(new[1][0:len(new[1])-1], 10, 64)
	return value
}

func intToBinary(integer int64, length int) string {
	memBinary := strconv.FormatInt(integer, 2)
	for len(memBinary) < length {
		memBinary = "0" + memBinary
	}
	return memBinary
}

func sumAddresses(mapping map[int64]int64) int64 {
	value := int64(0)
	for index := range mapping {
		value += mapping[index]
	}
	return value
}

func q14part1(data []string) {
	addresses := make(map[int64]int64)
	mask := ""
	for _, row := range data {
		split := strings.Split(row, " ")
		// Set the mask
		if split[0] == "mask" {
			mask = split[2]
			continue
		}

		// Now process the mem[] ops
		memIndex := getMemIndex(split[0])
		memDecimal, _ := strconv.ParseInt(split[2], 10, 64)
		memBinary := intToBinary(memDecimal, 36)
		// Gross, doing this with strings.  SHould look into bitsetting
		bit_value := ""
		for i := 0; i < len(mask); i++ {
			if mask[i] != 'X' {
				bit_value += string(mask[i])
			} else {
				if i >= len(memBinary) {
					bit_value += "0"
				} else {
					bit_value += string(memBinary[i])
				}
			}
		}
		addresses[memIndex], _ = strconv.ParseInt(bit_value, 2, 64)
	}
	fmt.Println("Day 13 Part 1 Solution:")
	fmt.Println(sumAddresses(addresses))
}

func convertAddressesToAllStates(addressString string) []int64 {
	// xMap = where are the xes
	xMap := make(map[int]bool)
	xMapLength := 0
	for i, xrune := range addressString {
		if xrune == 'X' {
			xMap[i] = true
			xMapLength += 1
		}
	}
	permutations := int(math.Pow(2, float64(xMapLength)))
	var values []int64
	for i := 0; i < permutations; i++ {
		// Int values will tell us what to do with each value of the xMap
		intValues := intToBinary(int64(i), xMapLength)
		newInteger := ""
		intValuesCursor := 0
		for j := range addressString {
			if xMap[j] == true {
				newInteger += string(intValues[intValuesCursor])
				intValuesCursor += 1
			} else {
				newInteger += string(addressString[j])
			}
		}

		integerAddress, _ := strconv.ParseInt(newInteger, 2, 64)
		values = append(values, integerAddress)
	}
	return values
}

func q14part2(data []string) {

	addresses := make(map[int64]int64)
	mask := ""
	for _, row := range data {
		split := strings.Split(row, " ")
		// Set the mask
		if split[0] == "mask" {
			mask = split[2]
			continue
		}

		// Now process the mem[] ops
		memIndex := getMemIndex(split[0])
		memDecimal, _ := strconv.ParseInt(split[2], 10, 64)
		memIndexBinary := intToBinary(memIndex, 36)
		// Gross, doing this with strings.  SHould look into bitsetting
		bit_value := ""
		for i := 0; i < len(mask); i++ {
			if mask[i] == 'X' {
				bit_value += "X"
			} else if mask[i] == '1' {
				bit_value += "1"
			} else if mask[i] == '0' {
				bit_value += string(memIndexBinary[i])
			} else {
				panic("halp")
			}

		}
		allAddresses := convertAddressesToAllStates(bit_value)
		for _, addr := range allAddresses {
			addresses[addr] = memDecimal
		}
	}
	fmt.Println("Day 13 Part 2 Solution:")
	fmt.Println(sumAddresses(addresses))
}
