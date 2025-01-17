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
	"aoc2024-go/days/day11"
	"aoc2024-go/days/day12"
	"aoc2024-go/days/day13"
	"aoc2024-go/days/day14"
	"aoc2024-go/days/day15"
	"aoc2024-go/days/day16"
	"aoc2024-go/days/day17"
	"aoc2024-go/days/day18"
	"aoc2024-go/days/day19"
	"aoc2024-go/days/day20"
	"aoc2024-go/days/day21"
	"aoc2024-go/days/day22"
	"aoc2024-go/days/day23"
	"aoc2024-go/days/day24"
	"aoc2024-go/days/day25"
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

	var result_1_str string
	var result_2_str string

	input := get_input(day)
	switch day {
	case "1":
		result_1 = day01.Solve1(input)
		result_2 = day01.Solve2(input)
	case "2":
		result_1 = day02.Solve1(input)
		result_2 = day02.Solve2(input)
	case "3":
		result_1 = day03.Solve1(input)
		result_2 = day03.Solve2(input)
	case "4":
		result_1 = day04.Solve1(input)
		result_2 = day04.Solve2(input)
	case "5":
		result_1 = day05.Solve1(input)
		result_2 = day05.Solve2(input)
	case "6":
		result_1 = day06.Solve1(input)
		result_2 = day06.Solve2(input)
	case "7":
		result_1 = day07.Solve1(input)
		result_2 = day07.Solve2(input)
	case "8":
		result_1 = day08.Solve1(input)
		result_2 = day08.Solve2(input)
	case "9":
		result_1 = day09.Solve1(input)
		result_2 = day09.Solve2(input)
	case "10":
		result_1 = day10.Solve1(input)
		result_2 = day10.Solve2(input)
	case "11":
		result_1 = day11.Solve1(input)
		result_2 = day11.Solve2(input)
	case "12":
		result_1 = day12.Solve1(input)
		result_2 = day12.Solve2(input)
	case "13":
		result_1 = day13.Solve1(input)
		result_2 = day13.Solve2(input)
	case "14":
		result_1 = day14.Solve1(input)
		result_2 = day14.Solve2(input)
	case "15":
		result_1 = day15.Solve1(input)
		result_2 = day15.Solve2(input)
	case "16":
		result_1 = day16.Solve1(input)
		result_2 = day16.Solve2(input)
	case "17":
		result_1_str = day17.Solve1(input)
		result_2 = day17.Solve2(input)
	case "18":
		result_1 = day18.Solve1(input)
		result_2_str = day18.Solve2(input)
	case "19":
		result_1 = day19.Solve1(input)
		result_2 = day19.Solve2(input)
	case "20":
		result_1 = day20.Solve1(input)
		result_2 = day20.Solve2(input)
	case "21":
		result_1 = day21.Solve1(input)
		result_2 = day21.Solve2(input)
	case "22":
		result_1 = day22.Solve1(input)
		result_2 = day22.Solve2(input)
	case "23":
		result_1 = day23.Solve1(input)
		result_2_str = day23.Solve2(input)
	case "24":
		result_1 = day24.Solve1(input)
		result_2_str = day24.Solve2(input)
	case "25":
		result_1 = day25.Solve1(input)
		result_2 = day25.Solve2(input)
	default:
		fmt.Printf("Solution for day %s is not implemented.\n", day)
		os.Exit(1)
	}

	// Display solution
	if result_1_str == "" {
		fmt.Printf("Solution for day %s: part1: %d\n", day, result_1)
	} else {
		fmt.Printf("Solution for day %s: part1: %s\n", day, result_1_str)
	}

	if result_2_str == "" {
		fmt.Printf("Solution for day %s: part2: %d\n", day, result_2)
	} else {
		fmt.Printf("Solution for day %s: part2: %s\n", day, result_2_str)
	}
}

func get_input(num string) string {
	// Zero pad
	if len(num) == 1 {
		num = "0" + num
	}

	path := "days/day" + num + "/input.txt"
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", path, err)
	}
	return string(input)
}
