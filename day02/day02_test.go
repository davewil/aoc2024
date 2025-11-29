package main

import (
	"testing"
)

var exampleInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestPart1(t *testing.T) {
	data, err := parseInput(exampleInput)
	if err != nil {
		t.Fatalf("parseInput() error = %v", err)
	}

	result := part1(data)
	expected := 2

	if result != expected {
		t.Errorf("part1() = %d; want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	data, err := parseInput(exampleInput)
	if err != nil {
		t.Fatalf("parseInput() error = %v", err)
	}

	result := part2(data)
	expected := 4

	if result != expected {
		t.Errorf("part2() = %d; want %d", result, expected)
	}
}
