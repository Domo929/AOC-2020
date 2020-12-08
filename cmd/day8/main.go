package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op      string
	val     int
	visited bool
}

func main() {
	f, err := os.Open("inputs/day8/input_full.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	instructions := make([]instruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		in, err := parseLine(line)
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, *in)

	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(part1Execute(instructions))
	fmt.Println(part2(instructions))
}

func part1Execute(instructions []instruction) int {
	acc := 0
	ndx := 0

	for {
		in := instructions[ndx]
		if in.visited {
			return acc
		}
		instructions[ndx].visited = true

		switch in.op {
		case "nop":
			ndx++
		case "acc":
			acc += in.val
			ndx++
		case "jmp":
			ndx += in.val
		}
	}
}

func part2(instructions []instruction) int {
	for ndx, in := range instructions {
		newInstructions := make([]instruction, len(instructions))
		for n, i := range instructions {
			newInstructions[n] = i
			newInstructions[n].visited = false
		}

		switch in.op {
		case "nop":
			newInstructions[ndx].op = "jmp"
		case "jmp":
			newInstructions[ndx].op = "nop"
		case "acc":
			continue
		}
		if acc, ok := part2Execute(newInstructions); ok {
			return acc
		}
	}
	return 0
}
func part2Execute(instructions []instruction) (int, bool) {
	acc := 0

	lastNdx := -1
	ndx := 0

	for {
		if lastNdx == ndx {
			return acc, false
		}
		if ndx == len(instructions) {
			return acc, true
		}
		in := instructions[ndx]
		if in.visited {
			return acc, false
		}
		instructions[ndx].visited = true

		lastNdx = ndx
		switch in.op {
		case "nop":
			ndx++
		case "acc":
			acc += in.val
			ndx++
		case "jmp":
			ndx += in.val
		}

	}
}

func parseLine(line string) (*instruction, error) {
	in := new(instruction)
	parts := strings.Split(line, " ")
	in.op = strings.Trim(parts[0], " ")

	val, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, err
	}
	in.val = int(val)

	return in, nil
}
