package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("inputs/day1/input.tsv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	csvR := csv.NewReader(f)

	inputs := make([]int, 0)
	for {
		line, err := csvR.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		inInt, err := strconv.ParseInt(line[0], 10, 32)
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, int(inInt))
	}

	part1(inputs)
	part2(inputs)
}

func part1(inputs []int) {
	inputMap := make(map[int]struct{})
	for _, input := range inputs {
		inputMap[input] = struct{}{}
	}

	found := make(map[int]struct{})
	for _, input := range inputs {
		other := 2020 - input
		if _, ok := inputMap[other]; ok {
			if _, ok := found[other]; !ok {
				found[input] = struct{}{}
				fmt.Printf("Part1 - Inputs: (%d | %d) - Product: %d\n", input, other, input*other)
			}
		}
	}
}

func part2(inputs []int) {
	inputMap := make(map[int]struct{})
	for _, input := range inputs {
		inputMap[input] = struct{}{}
	}

	found := make(map[int]struct{})
	for _, input1 := range inputs {
		for _, input2 := range inputs {
			if input1 != input2 {
				other := 2020 - (input1 + input2)
				if _, ok := inputMap[other]; ok {
					if _, ok := found[other]; !ok {
						found[input1] = struct{}{}
						found[input2] = struct{}{}
						found[other] = struct{}{}
						fmt.Printf("Part 2 - Inputs: (%d | %d | %d) - Product: %d\n", input1, input2, other, input1*input2*other)
					}
				}
			}
		}
	}
}
