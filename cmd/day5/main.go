package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("inputs/day5/input_full.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	ids := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		row, col, seatID := part1(line)
		ids = append(ids, seatID)
		fmt.Printf("%s: row %d, col %d, seatID %d\n", line, row, col, seatID)
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	sort.Ints(ids)
	fmt.Println("largest SeatID: ", ids[len(ids)-1])
	fmt.Println("My ID: ", part2(ids))
}

func part1(boardingPass string) (int, int, int) {
	row := findMiddle(0, 127, boardingPass[:7])
	col := findMiddle(0, 7, boardingPass[7:])

	return row, col, row*8 + col
}
func part2(ids []int) int {
	last := ids[0] - 1
	for _, id := range ids {
		if id-last > 1 {
			return id - 1
		}
		last = id
	}
	return 0
}

func findMiddle(min, max int, pass string) int {
	if pass == "" {
		return min
	}
	cmd := pass[0]
	half := ((max - min) / 2) + 1
	switch cmd {
	case 'F', 'L':
		max -= half
	case 'B', 'R':
		min += half
	}
	return findMiddle(min, max, pass[1:])
}
