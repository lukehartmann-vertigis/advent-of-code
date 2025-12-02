package main

import (
	"log"

	aoc_2025 "github.com/Djosar/advent-of-code/2025/1"
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

	year2025 := []*aoc_helpers.AOCDay{
		aoc_2025.Day01(),
		aoc_2025_day02.Day02(),
	}

	for _, day := range year2025 {
		day.Run()
	}
}
