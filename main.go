package main

import (
	"aoc2024-go/days/day01"
	"aoc2024-go/days/day02"
	"fmt"
	"log"
	"os"
)

func main() {
	// Get args
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")

		os.Exit(1)
	}
	day := os.Args[1]

	// Solve
	var result_1 int
	var result_2 int
	switch day {
	case "1":
		result_1 = day01.Solve1(get_input("01"))
		result_2 = day01.Solve2(get_input("01"))
	case "2":
		result_1 = day02.Solve1(get_input("02"))
		result_2 = day02.Solve2(get_input("02"))
	default:
		fmt.Printf("Solution for day %s is not implemented.\n", day)
		os.Exit(1)
	}

	// Display sol
	fmt.Printf("Solution for day %s: part1: %d, part2: %d\n", day, result_1, result_2)
}

func get_input(num string) string {
	path := "days/day" + num + "/input.txt"
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", path, err)
	}
	return string(input)
}
