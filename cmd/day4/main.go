package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport map[string]string

var (
	hairRegex = regexp.MustCompile(`#[a-f0-9]{6}`)
)

func main() {
	f, err := os.Open("inputs/day4/input_full.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	part1Valid := 0
	part2Valid := 0

	ppFields := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			pport := toPassport(ppFields)
			if pport.isValid(requiredFields, false) {
				part1Valid++
			}
			if pport.isValid(requiredFields, true) {
				part2Valid++
			}
			ppFields = make([]string, 0)
		}
		if len(line) > 0 {
			ppFields = append(ppFields, strings.Split(line, " ")...)
		}
	}
	fmt.Printf("Valid Part 1: %d\n", part1Valid)
	fmt.Printf("Valid Part 2: %d\n", part2Valid)
}

func toPassport(fields []string) passport {
	p := make(passport)
	for _, field := range fields {
		parts := strings.Split(field, ":")
		p[parts[0]] = parts[1]
	}
	return p
}

func (p passport) isValid(required []string, verifyFields bool) bool {
	for _, fieldKey := range required {
		fieldVal, ok := p[fieldKey]
		if !ok {
			return false
		}
		if verifyFields {
			switch fieldKey {
			case "byr":
				if !validNumber(fieldVal, 4, 1920, 2002) {
					return false
				}
			case "iyr":
				if !validNumber(fieldVal, 4, 2010, 2020) {
					return false
				}
			case "eyr":
				if !validNumber(fieldVal, 4, 2020, 2030) {
					return false
				}
			case "hgt":
				if !validHeight(fieldVal) {
					return false
				}
			case "hcl":
				if !validHair(fieldVal) {
					return false
				}
			case "ecl":
				if !validEye(fieldVal) {
					return false
				}
			case "pid":
				if !validPassportID(fieldVal) {
					return false
				}
			}
		}
	}
	return true
}

func validNumber(field string, length, min, max int) bool {
	if len(field) != length {
		return false
	}
	num, err := strconv.ParseInt(field, 10, 32)
	if err != nil {
		return false
	}
	valid := int(num) >= min && int(num) <= max

	return valid
}

func validHeight(field string) bool {
	if strings.Contains(field, "cm") {
		height := strings.TrimSuffix(field, "cm")
		return validNumber(height, 3, 150, 193)
	}
	if strings.Contains(field, "in") {
		height := strings.TrimSuffix(field, "in")
		return validNumber(height, 2, 59, 76)
	}
	return false
}

func validHair(field string) bool {
	return hairRegex.MatchString(field)
}

func validEye(field string) bool {
	allowedColors := map[string]struct{}{
		"amb": {},
		"blu": {},
		"brn": {},
		"gry": {},
		"grn": {},
		"hzl": {},
		"oth": {},
	}
	_, ok := allowedColors[field]
	return ok
}

func validPassportID(field string) bool {
	return len(field) == 9
}
