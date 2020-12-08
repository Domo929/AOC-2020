package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("inputs/day6/input_full.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	counts, sum := part1(lines)
	fmt.Println(counts)
	fmt.Println(sum)

	counts, sum = part2(lines)
	fmt.Println(counts)
	fmt.Println(sum)
}

func part1(lines []string) ([]int, int) {
	overallCounts := make([]int, 0)
	foundQuestions := make(map[rune]struct{})
	for _, line := range lines {
		if line == "" {
			overallCounts = append(overallCounts, len(foundQuestions))
			foundQuestions = make(map[rune]struct{})
			continue
		}

		for _, r := range line {
			foundQuestions[r] = struct{}{}
		}
	}
	overallCounts = append(overallCounts, len(foundQuestions))

	sum := 0
	for _, count := range overallCounts {
		sum += count
	}

	return overallCounts, sum
}

func part2(lines []string) ([]int, int) {
	overallCounts := make([]int, 0)
	foundQuestions := make(map[rune]struct{})
	firstInGroup := true
	for _, line := range lines {
		if line == "" {
			overallCounts = append(overallCounts, len(foundQuestions))
			foundQuestions = make(map[rune]struct{})
			firstInGroup = true
			continue
		}
		lineMap := make(map[rune]struct{})
		for _, r := range line {
			lineMap[r] = struct{}{}
		}
		if firstInGroup {
			foundQuestions = lineMap
		} else {
			for r := range foundQuestions {
				if _, ok := lineMap[r]; !ok {
					delete(foundQuestions, r)
				}
			}
		}
		firstInGroup = false
	}
	overallCounts = append(overallCounts, len(foundQuestions))

	sum := 0
	for _, count := range overallCounts {
		sum += count
	}

	return overallCounts, sum
}
