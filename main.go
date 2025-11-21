package main

import (
	"log"

	aoc_2024 "github.com/Djosar/advent-of-code/2024/1"
	aoc_helpers "github.com/Djosar/advent-of-code/util/aoc-helpers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
		return
	}

	year2024 := []*aoc_helpers.AOCDay{
		aoc_2024.Day01(),
	}

	for _, day := range year2024 {
		day.Run()
	}
}
