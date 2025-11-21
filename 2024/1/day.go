package aoc_2024

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
	"github.com/Djosar/advent-of-code/util/slice_helpers"
)

func Day01() *aoc_helpers.AOCDay {
	day := aoc_helpers.NewDay(1, 2024)

	day.PushTask(part01)
	day.PushTask(part02)

	return day
}

func part01(d *aoc_helpers.AOCDay) error {
	colA, colB := getNumCols(d.Lines)
	sum := 0

	slices.Sort(colA)
	slices.Sort(colB)

	for idx, item := range colA {
		diff := math.Abs(float64(item) - float64(colB[idx]))
		sum += int(diff)
	}

	fmt.Println(sum)

	return nil
}

func part02(d *aoc_helpers.AOCDay) error {

	return nil
}

func getNumCols(lines []string) ([]int, []int) {
	colA, colB := make([]int, 0, len(lines)), make([]int, 0, len(lines))
	for _, ln := range lines {
		split := slice_helpers.Map(strings.Split(ln, "   "), func(item string) int {
			val, _ := strconv.Atoi(item)
			return val
		})

		colA = append(colA, split[0])
		colB = append(colB, split[1])
	}

	return colA, colB
}
