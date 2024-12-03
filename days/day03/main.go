package day03

import (
	"fmt"
	"regexp"
	"strconv"
)

func Solve1(input_lines string) int {
	// Process the input
	multiply_commands := parse_input(input_lines)
	fmt.Println("Usage: go run main.go <day>", multiply_commands)

	sum := 0

	for _, multiply_command := range multiply_commands {
		fmt.Println("Usage: go run main.go <day>", multiply_command[0], multiply_command[1])
		sum += multiply_command[0] * multiply_command[1]
	}

	return sum
}

func Solve2(input_lines string) int {
	// Process the input
	return 1
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func parse_input(input_lines string) [][]int {
	multiply_commands := [][]int{}

	re, _ := regexp.Compile("mul\\((\\d*),(\\d*)\\)")

	matches := re.FindAllStringSubmatch(input_lines, -1)

	for _, match := range matches {
		left_num, _ := strconv.Atoi(match[1])
		right_num, _ := strconv.Atoi(match[2])
		multiply_command := []int{left_num, right_num}

		multiply_commands = append(multiply_commands, multiply_command)
	}
	return multiply_commands
}
