package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var (
	extractor = regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<letter>\w): (?P<password>\w+)`)
)

func main() {
	f, err := os.Open("inputs/day2/input_full.tsv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	csvR := csv.NewReader(f)

	validPart1 := 0
	validPart2 := 0
	for {
		line, err := csvR.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		pwd, err := toPassword(line[0])
		if err != nil {
			panic(err)
		}
		if pwd.ValidPart1() {
			validPart1++
		}
		if pwd.ValidPart2() {
			validPart2++
		}
	}
	fmt.Printf("Part 1 Valid: %d\n", validPart1)
	fmt.Printf("Part 2 Valid: %d\n", validPart2)
}

type password struct {
	Min      int
	Max      int
	Letter   rune
	password string
}

func toPassword(line string) (*password, error) {
	parts := extractor.FindStringSubmatch(line)

	min, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		return nil, err
	}
	max, err := strconv.ParseInt(parts[2], 10, 32)
	if err != nil {
		return nil, err
	}
	return &password{
		Min:      int(min),
		Max:      int(max),
		Letter:   []rune(parts[3])[0],
		password: parts[4],
	}, nil
}

func (pwd *password) ValidPart1() bool {
	n := 0
	for _, r := range []rune(pwd.password) {
		if r == pwd.Letter {
			n++
		}
	}
	return pwd.Min <= n && n <= pwd.Max
}

func (pwd *password) ValidPart2() bool {
	pos1R := []rune(pwd.password)[pwd.Min-1]
	pos2R := []rune(pwd.password)[pwd.Max-1]

	return (pos1R == pwd.Letter || pos2R == pwd.Letter) && !(pos1R == pwd.Letter && pos2R == pwd.Letter)
}
