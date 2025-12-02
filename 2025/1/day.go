package aoc_2025_day01

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
)

const testInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func Day01() *aoc_helpers.AOCDay {
	day := aoc_helpers.NewDay(1, 2025)

	day.PushTask(part01)
	day.PushTask(part02)

	return day
}

func part01(d *aoc_helpers.AOCDay) error {
	// testLines := strings.Split(testInput, "\n")
	count := calculatePart01(d.Lines)

	fmt.Println("PART 01:", count)
	return nil
}

func calculatePart01(lines []string) (count int) {
	count = 0
	dial := 50

	for _, line := range lines {
		val, err := parseLine(line)
		if err != nil {
			fmt.Println(err)
			break
		}
		dial += val
		if dial%100 == 0 {
			count += 1
		}
	}

	return count
}

func part02(d *aoc_helpers.AOCDay) error {
	// testLines := strings.Split(testInput, "\n")
	count := calculatePart02(d.Lines)
	fmt.Println("PART 02:", count)
	return nil
}

func calculatePart02(lines []string) (count int) {
	dial := 50
	count = 0
	for _, line := range lines {
		val, err := parseLine(line)
		if err != nil {
			break
		}

		absVal := int(math.Abs(float64(val)))
		mod := 1
		if val < 0 {
			mod = -1
		}

		for i := 0; i < absVal; i++ {
			dial += mod
			if dial%100 == 0 {
				count += 1
			}
		}
	}

	return count
}

func parseLine(line string) (int, error) {
	r := regexp.MustCompile(`([L,R])(\d+)`)
	matches := r.FindStringSubmatch(line)
	direction := matches[1]
	val, err := strconv.Atoi(matches[2])
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	if direction == "L" {
		val = -1 * val
	}
	return val, nil
}
