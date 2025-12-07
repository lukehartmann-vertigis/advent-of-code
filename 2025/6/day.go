package aoc_2025_day06

import (
	"fmt"
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
		fmt.Println(operation, col, res)

	}

	return res, nil
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
