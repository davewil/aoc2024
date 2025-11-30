package main

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
	"strconv"
)

var (
	mulPattern         = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	instructionPattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
)

func parseInput(input string) ([][2]int, error) {
	matches := mulPattern.FindAllStringSubmatch(input, -1)
	allPairs := make([][2]int, 0, len(matches))

	for _, m := range matches {
		x, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(m[2])
		if err != nil {
			return nil, err
		}

		allPairs = append(allPairs, [2]int{x, y})
	}

	return allPairs, nil
}

func part1(input [][2]int) int {
	total := 0
	for _, pair := range input {
		total += pair[0] * pair[1]
	}
	return total
}

func part2(input string) int {
	matches := instructionPattern.FindAllStringSubmatch(input, -1)
	enabled := true
	total := 0

	for _, m := range matches {
		switch m[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if !enabled {
				continue
			}

			if len(m) < 3 {
				continue
			}

			x, err := strconv.Atoi(m[1])
			if err != nil {
				continue
			}
			y, err := strconv.Atoi(m[2])
			if err != nil {
				continue
			}

			total += x * y
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

	data, err := parseInput(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(input))
}
