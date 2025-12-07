package aoc_2025_day05

import (
	"slices"
	"strconv"
	"strings"

	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
	"github.com/Djosar/advent-of-code/util/slice_helpers"
)

const testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func Day05() *aoc_helpers.AOCDay {
	day := aoc_helpers.NewDay(5, 2025)

	day.PushTask(part01)
	day.PushTask(part02)

	return day
}

func inputParser(input string) ([][]int, []int) {
	split := strings.Split(input, "\n\n")
	rangesStrings := strings.Split(split[0], "\n")
	ranges := slice_helpers.Map(rangesStrings, func(item string) []int {
		vals := strings.Split(item, "-")
		numA, _ := strconv.Atoi(vals[0])
		numB, _ := strconv.Atoi(vals[1])

		return []int{numA, numB}
	})

	slices.SortStableFunc(ranges, func(a, b []int) int {
		if a[0] == b[0] {
			return 0
		} else if a[0] < b[0] {
			return -1
		} else {
			return 1
		}
	})

	ids := slice_helpers.Map(strings.Split(split[1], "\n"), func(item string) int {
		num, _ := strconv.Atoi(item)
		return num
	})

	return ranges, ids
}

func part01(d *aoc_helpers.AOCDay) (any, error) {
	idsCnt := 0
	orderedRanges, ids := inputParser(d.Input)

	for _, id := range ids {
		for _, rng := range orderedRanges {
			if id < rng[0] {
				break
			}
			if id >= rng[0] && id <= rng[1] {
				idsCnt++
				break
			}
		}
	}

	return idsCnt, nil
}

func part02(d *aoc_helpers.AOCDay) (any, error) {
	orderedRanges, _ := inputParser(d.Input)

	merged := [][]int{orderedRanges[0]}
	for _, rng := range orderedRanges {
		lastMerged := merged[len(merged)-1]
		if rng[0] == lastMerged[0] && rng[1] == lastMerged[1] {
			continue
		}

		if rng[0] <= lastMerged[1] {
			if rng[1] <= lastMerged[1] {
				continue
			}

			if rng[1] > lastMerged[1] {
				lastMerged[1] = rng[1]
			}

		} else {
			merged = append(merged, rng)
		}
	}

	diffSum := 0
	for _, item := range merged {
		diff := item[1] - item[0]

		diffSum += diff + 1
	}

	return diffSum, nil
}
