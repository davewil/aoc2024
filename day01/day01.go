package main

import (
	"aoc2024/utils"
	"fmt"
)

func parseInput(input string) ([]string, error) {
	lines := utils.ReadLines(input)
	// Add your parsing logic here
	return lines, nil
}

func part1(input []string) int {
	// Solve part 1 here
	return 0
}

func part2(input []string) int {
	// Solve part 2 here
	return 0
}

func main() {
	utils.LoadEnv()

	input, err := utils.GetPuzzleInput(2024, 1)
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}

	data, err := parseInput(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
