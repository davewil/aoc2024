package main

import (
	"aoc2024/utils"
	"fmt"
	"slices"
)

func parseInput(input string) ([]int, []int, error) {
	lines := utils.ReadLines(input)
	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for i, line := range lines {
		var l, r int
		if _, err := fmt.Sscanf(line, "%d %d", &l, &r); err != nil {
			return nil, nil, err
		}
		left[i] = l
		right[i] = r
	}
	return left, right, nil
}

func part1(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	total := 0
	for i := range left {
		total += utils.Abs(left[i] - right[i])
	}
	return total
}

func part2(left []int, right []int) int {
	rightFreq := make(map[int]int)
	for _, v := range right {
		rightFreq[v]++
	}

	total := 0
	for _, v := range left {
		total += v * rightFreq[v]
	}
	return total
}

func main() {
	utils.LoadEnv()

	input, err := utils.GetPuzzleInput(2024, 1)
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}

	left, right, err := parseInput(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	fmt.Println("Part 1:", part1(left, right))
	fmt.Println("Part 2:", part2(left, right))
}
