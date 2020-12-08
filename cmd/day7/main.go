package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	lineRegex     = regexp.MustCompile(`^(\w+ \w+) bags contain((?:[,. ]+(?:\d{1}) (?:\w+ \w+) bag[s]?)+|(?: no other bags)).$`)
	contentsRegex = regexp.MustCompile(`^\s?(\d+)\s(\w+ \w+)`)
)

const (
	noOtherBags = " no other bags"
	shinyGold   = "shiny-gold"
)

type bagContents struct {
	num   int
	color string
}

func main() {
	f, err := os.Open("inputs/day7/input_full.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	isInside := make(map[string][]string)
	bagContains := make(map[string][]bagContents)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		matches := lineRegex.FindStringSubmatch(line)
		rootColor := strings.Replace(matches[1], " ", "-", -1)
		contains := matches[2]
		if contains == noOtherBags {
			bagContains[rootColor] = make([]bagContents, 0)
			continue
		}
		if bagContains[rootColor] == nil {
			bagContains[rootColor] = make([]bagContents, 0)
		}
		bagsInside := strings.Split(contains, ",")
		for _, b := range bagsInside {
			bagMatches := contentsRegex.FindStringSubmatch(strings.Trim(b, " "))
			count, err := strconv.ParseInt(bagMatches[1], 10, 64)
			if err != nil {
				panic(err)
			}
			insideColor := strings.Replace(bagMatches[2], " ", "-", -1)
			bagContains[rootColor] = append(bagContains[rootColor], bagContents{
				num:   int(count),
				color: insideColor,
			})
			if isInside[insideColor] == nil {
				isInside[insideColor] = make([]string, 0)
			}
			isInside[insideColor] = append(isInside[insideColor], rootColor)
		}
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	// for color, outer := range isInside {
	// 	fmt.Printf("%s is inside %v\n", color, outer)
	// }
	// fmt.Println("")
	// for rootColor, contents := range bagContains {
	// 	fmt.Printf("%s has %v\n", rootColor, contents)
	// }

	fmt.Println(part1(shinyGold, bagContains))
	fmt.Println(part2(shinyGold, bagContains))
}

func part2(rootColor string, containsMap map[string][]bagContents) int {
	var sum int
	bag := containsMap[rootColor]
	for _, bags := range bag {
		sum += bags.num * sumFunc(bags.color, containsMap)
	}
	return sum
}

func sumFunc(rootColor string, containsMap map[string][]bagContents) int {
	sum := 1
	bag := containsMap[rootColor]
	for _, bags := range bag {
		sum += bags.num * sumFunc(bags.color, containsMap)
	}
	return sum
}

func part1(findColor string, containsMap map[string][]bagContents) int {
	count := 0
	for color := range containsMap {
		if color == findColor {
			continue
		}
		if contains(color, findColor, containsMap) {
			count++
		}
	}
	return count
}

func contains(rootColor, findColor string, containsMap map[string][]bagContents) bool {
	if rootColor == findColor {
		return true
	}

	var has bool
	for _, bag := range containsMap[rootColor] {
		has = has || contains(bag.color, findColor, containsMap)
	}
	return has
}
