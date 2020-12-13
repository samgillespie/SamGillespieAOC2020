package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	lib "../lib"
)

func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(13)
	timestamp, _ := strconv.Atoi(data[0])
	buses := strings.Split(data[1], ",")
	q13part1(timestamp, buses)
	q13part2(timestamp, buses)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func q13part1(timestamp int, buses []string) {
	min_bus_time := 9999999
	solution := 99999
	for _, bus := range buses {
		if bus != "x" {
			bus_num, _ := strconv.Atoi(bus)
			remainder := float64(timestamp)/float64(bus_num) - float64(timestamp/bus_num)
			time_to_leave := bus_num - int(math.Round(remainder*float64(bus_num)))
			//fmt.Printf("%d, %f, %d\n", bus_num, remainder, time_to_leave)
			if time_to_leave < min_bus_time {
				min_bus_time = time_to_leave
				solution = time_to_leave * bus_num
			}
		}
	}
	fmt.Println("Day 13 Part 1 Solution:")
	fmt.Println(solution)

}

func chineseRemainderTheorem(a, n []*big.Int) (*big.Int, error) {
	// Code shamelessly stolen from the internet.  I could implement by hand, but cbf
	var one = big.NewInt(1)
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func q13part2(timestamp int, buses []string) {
	var n_values []*big.Int
	var i_values []*big.Int

	for index, bus := range buses {
		if bus != "x" {

			busNumber, _ := strconv.ParseInt(bus, 10, 64)
			n_values = append(n_values, big.NewInt(busNumber))
			indexNumber := new(big.Int)
			if index == 0 {
				indexNumber = big.NewInt(0)
			} else {
				indexNumber = big.NewInt(busNumber - int64(index))
			}
			i_values = append(i_values, indexNumber)
		}
	}
	fmt.Println("Day 13 Part 2 Solution:")
	fmt.Println(chineseRemainderTheorem(i_values, n_values))
}
