package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("inputs/day9/input_full.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	nums := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, int(num))
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	invalid := part1(25, 25, nums)
	weakness := part2(invalid, nums)
	fmt.Printf("Invalid: %d | Weakness: %d\n", invalid, weakness)

}
func part2(invalid int, nums []int) int {
	for i, startNum := range nums {
		sum := startNum
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == invalid {
				min, max := findLowHigh(nums[i : j+1])
				return min + max
			}
			if sum > invalid {
				break
			}
		}
	}
	return 0
}

func findLowHigh(nums []int) (int, int) {
	max := math.MinInt64
	min := math.MaxInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return min, max
}

func part1(preamble, lookback int, nums []int) int {
	for i := preamble; i < len(nums); i++ {
		lookbackNums := nums[i-lookback : i]
		if !validate(nums[i], lookbackNums) {
			return nums[i]
		}
	}
	return 0
}

func validate(num int, lookbackNums []int) bool {
	for i := range lookbackNums {
		for j := i + 1; j < len(lookbackNums); j++ {
			if lookbackNums[i]+lookbackNums[j] == num {
				return true
			}
		}
	}
	return false
}
