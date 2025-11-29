package main

import (
	"testing"
)

var exampleInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	left, right, err := parseInput(exampleInput)
	if err != nil {
		t.Fatalf("parseInput() error = %v", err)
	}

	result := part1(left, right)
	expected := 11

	if result != expected {
		t.Errorf("part1() = %d; want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	left, right, err := parseInput(exampleInput)
	if err != nil {
		t.Fatalf("parseInput() error = %v", err)
	}

	result := part2(left, right)
	expected := 31

	if result != expected {
		t.Errorf("part2() = %d; want %d", result, expected)
	}
}
