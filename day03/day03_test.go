package main

import "testing"

var exampleInputPart1 = `xmul(2,4)%&mul[3,7]!@^don't()^mul(5,5)+mul(32,64]then(mul(11,8)do()mul(8,5))`
var exampleInputPart2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestPart1(t *testing.T) {
	instructions, err := parseInput(exampleInputPart1)
	if err != nil {
		t.Fatalf("parseInput() error = %v", err)
	}

	result := part1(instructions)
	expected := 161

	if result != expected {
		t.Errorf("part1() = %d; want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	instructions, err := parseInput(exampleInputPart2)
	if err != nil {
		t.Fatalf("parseInput() error = %v", err)
	}

	result := part2(instructions)
	expected := 48

	if result != expected {
		t.Errorf("part2() = %d; want %d", result, expected)
	}
}
