package aoc_2025_day03

import (
	"fmt"
	"strconv"
	"strings"

	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
)

const testInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func Day03() *aoc_helpers.AOCDay {
	day := aoc_helpers.NewDay(3, 2025)
	day.PushTask(part01)
	day.PushTask(part02)
	return day
}

func part01(d *aoc_helpers.AOCDay) (any, error) {
	sum := 0
	for _, line := range d.Lines {
		sum += joltage(line)
	}
	return sum, nil
}

func part02(d *aoc_helpers.AOCDay) (any, error) {

	return 0, nil
}

func joltage(line string) (joltage int) {
	nums := strings.Split(line, "")
	maxes := []string{"0", "0"}
	maxIdx := 0
	secondMaxIdx := 0
	for idx, val := range nums {
		if val <= maxes[0] {
			continue
		}
		maxes[0] = val
		maxIdx = idx
	}

	for idx, val := range nums {
		if maxIdx == len(nums)-1 && idx == maxIdx {
			continue
		}
		if maxIdx < len(nums)-1 && idx <= maxIdx {
			continue
		}
		if val <= maxes[1] {
			continue
		}

		maxes[1] = val
		secondMaxIdx = idx
	}

	if maxIdx > secondMaxIdx {
		joltage, _ = strconv.Atoi(fmt.Sprintf("%s%s", maxes[1], maxes[0]))
	} else {
		joltage, _ = strconv.Atoi(fmt.Sprintf("%s%s", maxes[0], maxes[1]))
	}

	return joltage
}
