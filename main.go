package main

import (
	"aoc2024-go/days/day01"
	"aoc2024-go/days/day02"
	"aoc2024-go/days/day03"
	"aoc2024-go/days/day04"
	"aoc2024-go/days/day05"
	"aoc2024-go/days/day06"
	"aoc2024-go/days/day07"
	"aoc2024-go/days/day08"
	"aoc2024-go/days/day09"
	"aoc2024-go/days/day10"
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
	case "3":
		result_1 = day03.Solve1(get_input("03"))
		result_2 = day03.Solve2(get_input("03"))
	case "4":
		result_1 = day04.Solve1(get_input("04"))
		result_2 = day04.Solve2(get_input("04"))
	case "5":
		result_1 = day05.Solve1(get_input("05"))
		result_2 = day05.Solve2(get_input("05"))
	case "6":
		result_1 = day06.Solve1(get_input("06"))
		result_2 = day06.Solve2(get_input("06"))
	case "7":
		result_1 = day07.Solve1(get_input("07"))
		result_2 = day07.Solve2(get_input("07"))
	case "8":
		result_1 = day08.Solve1(get_input("08"))
		result_2 = day08.Solve2(get_input("08"))
	case "9":
		result_1 = day09.Solve1(get_input("09"))
		result_2 = day09.Solve2(get_input("09"))
	case "10":
		result_1 = day10.Solve1(get_input("10"))
		result_2 = day10.Solve2(get_input("10"))
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
