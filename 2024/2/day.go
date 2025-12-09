package aoc_2024_day02

import (
	"math"
	"strconv"
	"strings"

	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
	"github.com/Djosar/advent-of-code/util/slice_helpers"
)

const testInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Day02() *aoc_helpers.AOCDay {
	day := aoc_helpers.NewDay(2, 2024, testInput)
	day.PushTask(part01)
	day.PushTask(part02)
	return day
}

func part01(d *aoc_helpers.AOCDay) (any, error) {
	lines := d.Lines

	safeReportCount := 0
	for _, ln := range lines {
		if safeReport(ln) {
			safeReportCount += 1
		}
	}

	return safeReportCount, nil
}

func part02(d *aoc_helpers.AOCDay) (any, error) {
	return 0, nil
}

func safeReport(line string) bool {
	report := slice_helpers.Map(strings.Split(line, " "), func(item string) int {
		val, _ := strconv.Atoi(item)
		return val
	})

	var increasing bool
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff == 0 {
			return false
		}

		if math.Abs(float64(diff)) > 3 {
			return false
		}

		if i == 1 {
			increasing = diff > 0
			continue
		}

		currentIncreasing := diff > 0

		if currentIncreasing != increasing {
			return false
		}
	}

	return true
}

func testLines() []string {
	return strings.Split(testInput, "\n")
}
