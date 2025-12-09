package aoc_2025_day06

import (
	"regexp"
	"strconv"
	"strings"

	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
	"github.com/Djosar/advent-of-code/util/slice_helpers"
)

const testInput = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func Day06() *aoc_helpers.AOCDay {
	day := aoc_helpers.NewDay(6, 2025)
	day.PushTask(part01)
	day.PushTask(part02)
	return day
}

func part01(d *aoc_helpers.AOCDay) (any, error) {
	lines := d.Lines

	convertedLines := make([][]int, len(lines)-1)
	operations := []string{}
	for lineIdx, line := range lines {
		if lineIdx < len(lines)-1 {
			split := slice_helpers.Filter(strings.Split(line, " "), func(item string) bool {
				return len(strings.TrimSpace(item)) > 0
			})

			vals := slice_helpers.Map(split, func(item string) int {
				val, _ := strconv.Atoi(item)
				return val
			})

			convertedLines[lineIdx] = vals
		} else {
			operations = slice_helpers.Filter(strings.Split(line, " "), func(item string) bool {
				return len(strings.TrimSpace(item)) > 0
			})
		}
	}

	res := 0

	for x := range convertedLines[0] {
		operation := operations[x]
		col := []int{}
		for y := range convertedLines {
			col = append(col, convertedLines[y][x])
		}

		start := 0
		if operation == "*" {
			start = 1
		}

		res += slice_helpers.Reduce(col, func(acc, curr int) int {
			return runOperation(operation)(acc, curr)
		}, start)

	}

	return res, nil
}

func part02(d *aoc_helpers.AOCDay) (any, error) {
	lines := d.Lines
	opsIndex := len(lines) - 1
	opsLine := lines[opsIndex]
	valueLines := lines[:opsIndex]
	opRegex := regexp.MustCompile(`[+*]\s`)
	colRanges := slice_helpers.Map(opRegex.FindAllStringIndex(opsLine, -1), func(item []int) int {
		return item[0]
	})

	newNums := make([][]int, len(colRanges))

	for rngIdx, start := range colRanges {
		end := len(opsLine)
		if rngIdx < len(colRanges)-1 {
			end = colRanges[rngIdx+1] - 1
		}

		numStringsCnt := end - start
		newNumStrings := make([]string, numStringsCnt)

		for _, line := range valueLines {
			val := line[start:end]

			for charIdx, char := range val {
				currentVal := newNumStrings[charIdx]
				newNumStrings[charIdx] = strings.TrimSpace(currentVal + string(char))
			}
		}

		newNums[rngIdx] = slice_helpers.Map(newNumStrings, func(item string) int {
			val, _ := strconv.Atoi(item)
			return val
		})
	}

	sum := 0
	for colIdx, start := range colRanges {
		op := string(opsLine[start])
		start := 0
		if op == "*" {
			start = 1
		}
		res := slice_helpers.Reduce(newNums[colIdx], func(acc, curr int) int {
			return runOperation(op)(acc, curr)
		}, start)
		sum += res
	}

	return sum, nil
}

func runOperation(op string) func(a, b int) int {
	switch op {
	case "*":
		return func(a, b int) int {
			return a * b
		}
	case "+":
		return func(a, b int) int {
			return a + b
		}
	default:
		return nil
	}
}
