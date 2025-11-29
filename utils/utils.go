package utils

import (
	"strconv"
	"strings"
)

// ReadLines splits input text into lines, trimming the final newline if present
func ReadLines(input string) []string {
	input = strings.TrimSuffix(input, "\n")
	return strings.Split(input, "\n")
}

// ParseInts converts a slice of strings to a slice of integers
func ParseInts(lines []string) ([]int, error) {
	nums := make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		nums[i] = n
	}
	return nums, nil
}

// Abs returns the absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
