package aoc_2025_day04

import (
	"strings"

	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
	"github.com/Djosar/advent-of-code/util/collections"
	"github.com/Djosar/advent-of-code/util/slice_helpers"
)

const testInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func Day04() *aoc_helpers.AOCDay {
	day := aoc_helpers.NewDay(4, 2025)
	day.PushTask(part01)
	day.PushTask(part02)
	return day
}

func part01(d *aoc_helpers.AOCDay) (any, error) {
	testLines := strings.Split(testInput, "\n")
	testItems := strings.Split(strings.ReplaceAll(testInput, "\n", ""), "")

	matrix := collections.NewMatrix(testItems, len(testLines[0]))

	rollCount := 0
	for y := 0; y < matrix.Height; y++ {
		for x := 0; x < matrix.Width; x++ {
			char, _ := matrix.At(x, y)
			if char != "@" {
				continue
			}
			moore, err := matrix.Moore(x, y)

			if err != nil {
				return nil, err
			}
			count := slice_helpers.Count(moore, "@")

			if count < 4 {
				rollCount++
			}
		}
	}
	return rollCount, nil
}

func part02(d *aoc_helpers.AOCDay) (any, error) {
	/*testLines := strings.Split(testInput, "\n")
	testItems := strings.Split(strings.ReplaceAll(testInput, "\n", ""), "")*/

	matrix := d.Matrix

	rollsCount := 0
	run := true
	for run {
		run = false
		for y := 0; y < matrix.Height; y++ {
			for x := 0; x < matrix.Width; x++ {
				char, _ := matrix.At(x, y)
				if char != "@" {
					continue
				}
				moore, err := matrix.Moore(x, y)

				if err != nil {
					return nil, err
				}
				count := slice_helpers.Count(moore, "@")

				if count < 4 {
					matrix.Set(x, y, "x")
					rollsCount++
					run = true
				}
			}
		}
	}

	return rollsCount, nil
}
