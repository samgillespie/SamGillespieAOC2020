package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	lib "../lib"
)

// Data Structures
type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (pass *passport) Print() {
	fmt.Printf("byr:"+pass.byr, "iyr:"+pass.iyr, "eyr:"+pass.eyr, "hgt:"+pass.hgt, "hcl:"+pass.hcl, "ecl:"+pass.ecl, "pid:"+pass.pid, "cid:"+pass.cid)
}

func (pass *passport) SetValue(field_name string, field_value string) {
	if field_name == "byr" {
		pass.byr = field_value
	}
	if field_name == "iyr" {
		pass.iyr = field_value
	}
	if field_name == "eyr" {
		pass.eyr = field_value
	}
	if field_name == "hgt" {
		pass.hgt = field_value
	}
	if field_name == "hcl" {
		pass.hcl = field_value
	}
	if field_name == "ecl" {
		pass.ecl = field_value
	}
	if field_name == "pid" {
		pass.pid = field_value
	}
	if field_name == "cid" {
		pass.cid = field_value
	}
}

func (pass *passport) isValidP1() bool {
	if pass.byr == "" {
		return false
	} else if pass.iyr == "" {
		return false
	} else if pass.eyr == "" {
		return false
	} else if pass.hgt == "" {
		return false
	} else if pass.hcl == "" {
		return false
	} else if pass.ecl == "" {
		return false
	} else if pass.pid == "" {
		return false
	}
	//} else if pass.cid == "" {
	//	return false
	//}
	return true
}

func (pass *passport) ValidBirthYear() bool {
	// Birth Year
	if pass.byr == "" {
		return false
	}
	birth_year, _ := strconv.Atoi(pass.byr)
	if birth_year < 1920 || birth_year > 2002 {
		return false
	}
	return true
}

func (pass *passport) ValidIssueYear() bool {
	// Issue Year
	if pass.iyr == "" {
		return false
	}
	issue_year, _ := strconv.Atoi(pass.iyr)
	if issue_year < 2010 || issue_year > 2020 {
		return false
	}
	return true
}

func (pass *passport) ValidExpirationYear() bool {
	// Expiration year
	if pass.eyr == "" {
		return false
	}
	expiration_year, _ := strconv.Atoi(pass.eyr)
	if expiration_year < 2020 || expiration_year > 2030 {
		return false
	}
	return true
}

func (pass *passport) ValidHeight() bool {
	// Height
	if pass.hgt == "" {
		return false
	}
	if len(pass.hgt) <= 2 {
		return false
	}
	height_suffix := pass.hgt[len(pass.hgt)-2:]
	height_value, err := strconv.Atoi(pass.hgt[:len(pass.hgt)-2])
	if err != nil {
		log.Printf(pass.hgt)
		panic(err.Error())
	}
	if height_suffix != "cm" && height_suffix != "in" {
		return false
	}
	if height_suffix == "cm" && (height_value < 150 || height_value > 193) {
		return false
	}
	if height_suffix == "in" && (height_value < 59 || height_value > 76) {
		return false
	}
	return true
}

func (pass *passport) ValidHairColor() bool {
	//Hair color
	if pass.hcl == "" {
		return false
	}

	if len(pass.hcl) != 7 {
		return false
	}

	if pass.hcl[0:1] != "#" {
		return false
	}
	var is_color = regexp.MustCompile(`([a-f0-9]{5})\w+`).MatchString
	if is_color(pass.hcl[1:len(pass.hcl)]) == false {
		return false
	}
	return true
}

func (pass *passport) ValidEyeColor() bool {
	// eye color
	if pass.ecl == "" {
		return false
	}

	invalid_eye := true
	switch pass.ecl {
	case
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth":
		invalid_eye = false
	}
	if invalid_eye == true {
		return false
	}
	return true
}

func (pass *passport) ValidPassport() bool {
	if pass.pid == "" {
		return false
	}
	if len(pass.pid) != 9 {
		return false
	}
	return true
}

func (pass *passport) isValidP2() bool {

	if pass.ValidBirthYear() == false {
		return false
	}
	if pass.ValidIssueYear() == false {
		return false
	}
	if pass.ValidExpirationYear() == false {
		return false
	}
	if pass.ValidHeight() == false {
		return false
	}
	if pass.ValidHairColor() == false {
		return false
	}
	if pass.ValidEyeColor() == false {
		return false
	}
	if pass.ValidPassport() == false {
		return false
	}
	return true

}

// Logic
func main() {
	start := time.Now()

	data := lib.ReadInputAsStr(4)
	q4part1(data)
	q4part2(data)
	elapsed := time.Since(start)

	fmt.Printf("Main took %s", elapsed)
}

func q4part1(data []string) {

	active_passport := new(passport)
	valid_passports := 0
	invalid_passports := 0
	for i := 0; i < len(data); i++ {
		row_split := strings.Split(data[i], " ")
		// If we have an empty row, we're moving to a new passport
		if data[i] == "" {
			//active_passport.Print()
			if active_passport.isValidP1() {
				valid_passports += 1
			} else {
				invalid_passports += 1
			}
			active_passport = new(passport)
			continue
		}

		for j := 0; j < len(row_split); j++ {
			entry := strings.Split(row_split[j], ":")

			active_passport.SetValue(entry[0], entry[1])
		}
	}

	if active_passport.isValidP1() {
		valid_passports += 1
	} else {
		invalid_passports += 1
	}
	fmt.Println("Question 4 Part 1 Solution:")
	fmt.Println(valid_passports)
}

func q4part2(data []string) {

	active_passport := new(passport)
	valid_passports := 0
	invalid_passports := 0
	for i := 0; i < len(data); i++ {
		row_split := strings.Split(data[i], " ")
		// If we have an empty row, we're moving to a new passport
		if data[i] == "" {
			//active_passport.Print()
			if active_passport.isValidP2() {
				valid_passports += 1
			} else {
				invalid_passports += 1
			}
			active_passport = new(passport)
			continue
		}

		for j := 0; j < len(row_split); j++ {
			entry := strings.Split(row_split[j], ":")

			active_passport.SetValue(entry[0], entry[1])
		}
	}

	if active_passport.isValidP2() {
		valid_passports += 1
	} else {
		invalid_passports += 1
	}
	fmt.Println("Question 4 Part 2 Solution:")
	fmt.Println(valid_passports)
}
