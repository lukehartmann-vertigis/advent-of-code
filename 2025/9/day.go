package aoc_2025_day09

import (
	"math"
	"slices"
	"strconv"
	"strings"

	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
	"github.com/Djosar/advent-of-code/util/slice_helpers"
)

const testInput = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func Day09() *aoc_helpers.AOCDay {
	day := aoc_helpers.NewDay(9, 2025, testInput)
	day.PushTask(part01)
	day.PushTask(part02)
	return day
}

func part01(d *aoc_helpers.AOCDay) (any, error) {
	lines := d.Lines

	coords := calculateCoords(lines)

	biggestArea := 0
	for aIdx, coordA := range coords {
		for bIdx, coordB := range coords {
			if aIdx == bIdx {
				continue
			}

			if coordA[0] == coordB[0] || coordA[1] == coordB[1] {
				continue
			}

			sideA := math.Abs(float64(coordB[0]-coordA[0])) + 1
			sideB := math.Abs(float64(coordB[1]-coordA[1])) + 1
			area := int(sideA * sideB)
			if area > biggestArea {
				biggestArea = area
			}
		}
	}

	return biggestArea, nil
}

func part02(d *aoc_helpers.AOCDay) (any, error) {
	return 0, nil
}

func calculateCoords(lines []string) [][]int {
	coords := [][]int{}
	for _, ln := range lines {
		vals := strings.Split(ln, ",")
		coords = append(coords, slice_helpers.Map(vals, func(item string) int {
			c, _ := strconv.Atoi(item)
			return c
		}))
	}
	slices.SortFunc(coords, func(a, b []int) int {
		return a[0] - b[0]
	})
	return coords
}
