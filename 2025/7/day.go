package aoc_2025_day07

import (
	"regexp"

	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
	"github.com/Djosar/advent-of-code/util/slice_helpers"
)

const testInput = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

func Day07() *aoc_helpers.AOCDay {
	day := aoc_helpers.NewDay(7, 2025, testInput)
	day.PushTask(part01)
	day.PushTask(part02)
	return day
}

func part01(d *aoc_helpers.AOCDay) (any, error) {
	lines := d.Lines
	startRegex := regexp.MustCompile("S")
	// splitPointRegex := regexp.MustCompile(`\^`)

	// x indexes
	beamTracker := make([][]int, len(lines))
	splitCount := 0

	for y, ln := range lines {
		if y == 0 {
			match := startRegex.FindStringIndex(ln)
			beamTracker[0] = []int{
				match[0],
			}
			continue
		}

		beamStore := []int{}
		for idx, char := range ln {
			currentBeamTrack := beamTracker[y-1]
			idxRelevant := slice_helpers.Count(currentBeamTrack, idx) >= 1

			if !idxRelevant {
				continue
			}

			strChar := string(char)
			switch strChar {
			case ".":
				beamStore = append(beamStore, idx)

			case "^":
				left := idx - 1
				right := idx + 1
				if left >= 0 && slice_helpers.Count(beamStore, left) == 0 {
					beamStore = append(beamStore, left)
				}
				if right < len(ln) && slice_helpers.Count(beamStore, right) == 0 {
					beamStore = append(beamStore, right)
				}
				splitCount++
			}
		}
		beamTracker[y] = beamStore
	}

	return splitCount, nil
}

func part02(d *aoc_helpers.AOCDay) (any, error) {
	return 0, nil
}
