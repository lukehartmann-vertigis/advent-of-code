package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	aoc_2024_day01 "github.com/Djosar/advent-of-code/2024/1"
	aoc_2025_day01 "github.com/Djosar/advent-of-code/2025/1"
	aoc_2025_day02 "github.com/Djosar/advent-of-code/2025/2"
	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
		return
	}

	yearId, dayId := 0, 0
	yearsDict := yearsDict()

	if len(os.Args) == 1 {
		for _, y := range yearsDict {
			y.Run()
		}
		return
	}

	if len(os.Args) == 2 {
		yearId, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if len(os.Args) >= 3 {
		yearId, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		dayId, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if dayId == 0 {
		yearsDict[yearId].Run()
		return
	}

	yearsDict[yearId].RunDay(dayId)
}

func yearsDict() map[int]*aoc_helpers.AOCYear {
	return map[int]*aoc_helpers.AOCYear{
		2024: aoc_helpers.NewYear(2024, []*aoc_helpers.AOCDay{
			aoc_2024_day01.Day01(),
		}),
		2025: aoc_helpers.NewYear(2025, []*aoc_helpers.AOCDay{
			aoc_2025_day01.Day01(),
			aoc_2025_day02.Day02(),
		}),
	}
}
