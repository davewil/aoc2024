package main

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Instruction struct {
	Type string // "mul", "do", "don't"
	X    int
	Y    int
}

var (
	instructionPattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
)

func parseInput(input string) ([]Instruction, error) {
	matches := instructionPattern.FindAllStringSubmatch(input, -1)
	instructions := make([]Instruction, 0, len(matches))

	for _, m := range matches {
		switch m[0] {
		case "do()":
			instructions = append(instructions, Instruction{Type: "do"})
		case "don't()":
			instructions = append(instructions, Instruction{Type: "don't"})
		default:
			x, err := strconv.Atoi(m[1])
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(m[2])
			if err != nil {
				return nil, err
			}
			instructions = append(instructions, Instruction{Type: "mul", X: x, Y: y})
		}
	}

	return instructions, nil
}

func part1(input []Instruction) int {
	total := 0

	for _, inst := range input {
		if inst.Type == "mul" {
			total += inst.X * inst.Y
		}
	}

	return total
}

func part2(input []Instruction) int {
	enabled := true
	total := 0

	for _, inst := range input {
		switch inst.Type {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if enabled {
				total += inst.X * inst.Y
			}
		}
	}

	return total
}

func main() {
	utils.LoadEnv()

	input, err := utils.GetPuzzleInput(2024, 3)
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}

	instructions, err := parseInput(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	fmt.Println("Part 1:", part1(instructions))
	fmt.Println("Part 2:", part2(instructions))
}
