# Advent of Code 2024 - Go Solutions

Solutions for [Advent of Code 2024](https://adventofcode.com/2024) written in Go.

## Setup

1. Set your Advent of Code session cookie:
   ```bash
   export AOC_SESSION=your_session_cookie_here
   ```
   
   To find your session cookie:
   - Log in to [adventofcode.com](https://adventofcode.com)
   - Open browser developer tools (F12)
   - Go to Application/Storage > Cookies
   - Copy the value of the `session` cookie

2. Run a solution:
   ```bash
   cd day01
   go run day01.go
   ```

3. Run tests:
   ```bash
   cd day01
   go test
   ```

## Project Structure

```
aoc2024/
├── go.mod
├── utils/
│   ├── input.go    # Fetch puzzle inputs from AoC website
│   └── utils.go    # Common utility functions
├── day01/
│   ├── day01.go
│   ├── day01_test.go
│   └── input.txt   # Cached puzzle input (auto-downloaded)
└── ...
```

## Features

- Automatic puzzle input fetching from adventofcode.com
- Input caching to avoid repeated requests
- Test scaffolding for example inputs
- Common utility functions
