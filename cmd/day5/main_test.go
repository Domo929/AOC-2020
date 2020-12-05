package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMiddle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		startMin int
		startMax int
		exp      int
	}{
		{"Rows1", "FBFBBFF", 0, 127, 44},
		{"Rows2", "BFFFBBF", 0, 127, 70},
		{"Rows3", "FFFBBBF", 0, 127, 14},
		{"Rows4", "BBFFBBF", 0, 127, 102},
		{"Cols1", "RLR", 0, 7, 5},
		{"Cols2", "RRR", 0, 7, 7},
		{"Cols3", "RLL", 0, 7, 4},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			row := findMiddle(test.startMin, test.startMax, test.input)
			require.Equal(t, test.exp, row)
		})
	}

}
