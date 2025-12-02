package aoc_2025_day02

import (
	"fmt"
	"strconv"
	"strings"

	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
	"github.com/Djosar/advent-of-code/util/slice_helpers"
)

const testInput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func Day02() *aoc_helpers.AOCDay {
	day := aoc_helpers.NewDay(2, 2025)

	day.PushTask(part01)
	day.PushTask(part02)

	return day
}

func part01(d *aoc_helpers.AOCDay) error {
	ids := strings.Split(d.Input, ",")
	sums := slice_helpers.Map(ids, func(item string) int {
		return sumInvalids(item)
	})
	sum := slice_helpers.Reduce(sums, func(acc int, elem int) int {
		return acc + elem
	}, 0)

	fmt.Println(sum)
	return nil
}

func sumInvalids(rng string) int {
	split := strings.Split(rng, "-")
	splitInts := make([]int, len(split))
	for idx, item := range split {
		val, err := strconv.Atoi(item)
		if err != nil {
			fmt.Println(err)
			continue
		}
		splitInts[idx] = val
	}

	idRng := buildRange(splitInts[0], splitInts[1])
	invalidSum := 0
	for _, item := range idRng {
		idStr := strconv.Itoa(item)
		invalid := invalidIDs(idStr)
		if invalid {
			invalidSum += item
		}
	}

	return invalidSum
}

func invalidIDs(id string) bool {

	for i := 1; i < len(id); i++ {
		prefix := id[:i]
		tail := id[i:]
		if tail[0] == '0' {
			continue
		}
		if len(tail) > len(prefix) {
			continue
		}
		if strings.Contains(tail, prefix) {
			return true
		}
	}

	return false
}

func buildRange(start int, end int) []int {
	if start > end {
		return []int{}
	}

	diff := end - start
	rng := make([]int, 0, diff)

	for i := 0; i <= diff; i++ {
		rng = append(rng, start+i)
	}

	return rng
}

func part02(d *aoc_helpers.AOCDay) error {
	return nil
}

func testIDRanges() []string {
	return strings.Split(testInput, ",")
}
