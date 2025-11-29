package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file if it exists
// It searches in the current directory and parent directories
func LoadEnv() {
	// Try to load .env file from current dir or parent dirs
	// godotenv.Load() will search upwards for .env file
	_ = godotenv.Load()

	// Also try explicit path to project root
	// This handles cases where we run from subdirectories
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		_ = godotenv.Load("../.env")
	}
}

// GetPuzzleInput fetches the puzzle input for a given day from the Advent of Code website
// using the session cookie from the AOC_SESSION environment variable.
// It caches the input in a local file to avoid repeated requests.
func GetPuzzleInput(year, day int) (string, error) {
	// Create cache directory if it doesn't exist
	cacheDir := filepath.Join(".", fmt.Sprintf("day%02d", day))
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create cache directory: %w", err)
	}

	cacheFile := filepath.Join(cacheDir, "input.txt")

	// Check if input is already cached
	if data, err := os.ReadFile(cacheFile); err == nil {
		return string(data), nil
	}

	// Get session cookie from environment variable
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return "", fmt.Errorf("AOC_SESSION environment variable not set")
	}

	// Fetch input from Advent of Code website
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", session))
	req.Header.Set("User-Agent", "github.com/yourusername/aoc2024 by your@email.com")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch input: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch input: HTTP %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Cache the input
	if err := os.WriteFile(cacheFile, data, 0644); err != nil {
		return "", fmt.Errorf("failed to cache input: %w", err)
	}

	return string(data), nil
}
