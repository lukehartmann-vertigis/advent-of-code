package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const maxCap = 1024 * 1024 // 1MB
	buf := make([]byte, 0, 64*1024)

	scanner.Buffer(buf, maxCap)

	lines := make([]string, 0, 1024)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading lines: %w", err)
	}

	return lines, nil
}

func MustReadLines(path string) []string {
	lines, err := ReadLines(path)
	if err != nil {
		panic(err)
	}
	return lines
}

func ReadRuneGrid(path string) [][]rune {
	lines := MustReadLines(path)

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid
}
