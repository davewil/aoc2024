package main

import (
	"aoc2024/utils"
	"fmt"
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
		if isValid(report) {
			count++
			continue
		}

		for i := range report {
			variation := make([]int, 0, len(report)-1)
			variation = append(variation, report[:i]...)
			variation = append(variation, report[i+1:]...)

			if isValid(variation) {
				count++
				break
			}
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

	direction := 0 // 1 for increasing, -1 for decreasing
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff == 0 {
			return false
		}

		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else {
				direction = -1
			}
			continue
		}

		if direction > 0 && diff < 0 {
			return false
		}
		if direction < 0 && diff > 0 {
			return false
		}
	}

	return true
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
