package main

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseInput(input string) ([][]int, error) {
	lines := utils.ReadLines(input)
	result := make([][]int, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)
		nums := make([]int, len(fields))

		for i, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			nums[i] = n
		}

		result = append(result, nums)
	}

	return result, nil
}

func part1(input [][]int) int {
	count := 0
	for _, report := range input {
		if isValid(report) {
			count++
		}
	}
	return count
}

func part2(input [][]int) int {
	count := 0
	for _, report := range input {
		variations := generateVariations(report)
		if slices.ContainsFunc(variations, isValid) {
			count++
		}
	}
	return count
}

func isValid(report []int) bool {
	return test1(report) && test2(report)
}

func test1(report []int) bool {
	if len(report) < 2 {
		return true
	}

	increasing := 0
	decreasing := 0

	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1] {
			increasing++
		} else if report[i] < report[i-1] {
			decreasing++
		} else {
			return false
		}
	}

	return increasing == len(report)-1 || descreasing == len(report)-1
}

func test2(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		diff := utils.Abs(report[i] - report[i+1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func generateVariations(report []int) [][]int {
	variations := make([][]int, len(report))
	for i := 0; i < len(report); i++ {
		variation := make([]int, 0, len(report))
		variation = append(variation, report[:i]...)
		variation = append(variation, report[i+1:]...)
		variations[i] = variation
	}
	return variations
}

func main() {
	utils.LoadEnv()

	input, err := utils.GetPuzzleInput(2024, 2)
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}

	reports, err := parseInput(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	fmt.Println("Part 1:", part1(reports))
	fmt.Println("Part 2:", part2(reports))
}
