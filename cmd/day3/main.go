package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("inputs/day3/input_full.tsv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	csvR := csv.NewReader(f)

	// [y][x]
	grid := make([][]rune, 0)
	for {
		line, err := csvR.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		row := []rune(line[0])
		grid = append(grid, row)
	}

	fmt.Println("Part 1 trees hit: ", slope(3, 1, grid))
	fmt.Println("Part 2 total trees hit: ", multiSlope([]int{1, 3, 5, 7, 1}, []int{1, 1, 1, 1, 2}, grid))

}

func slope(dx, dy int, grid [][]rune) int {
	width := len(grid[0])
	depth := len(grid)

	x := 0
	y := 0

	count := 0
	for {
		y += dy
		x = (x + dx) % width
		if y >= depth {
			break
		}
		switch grid[y][x] {
		case '#':
			count++
		default:
			continue
		}
	}
	return count
}

func multiSlope(dxs, dys []int, grid [][]rune) int {
	total := 1
	for i := 0; i < len(dxs); i++ {
		total *= slope(dxs[i], dys[i], grid)
	}
	return total
}
