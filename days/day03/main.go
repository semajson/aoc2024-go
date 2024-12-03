package day03

import (
	"regexp"
	"strconv"
)

func Solve1(input_lines string) int {
	// Process the input
	commands := parse_input(input_lines)

	sum := 0
	for _, command := range commands {
		switch command.name {
		case "mul":
			sum += command.param1 * command.param2
		}
	}

	return sum
}

func Solve2(input_lines string) int {
	// Process the input
	commands := parse_input(input_lines)

	sum := 0
	mul_enabled := true
	for _, command := range commands {
		switch command.name {
		case "mul":
			if mul_enabled {
				sum += command.param1 * command.param2
			}
		case "do":
			mul_enabled = true
		case "don't":
			mul_enabled = false
		}
	}
	return sum
}

func parse_input(input_lines string) []command {
	commands := []command{}

	re, _ := regexp.Compile(`(mul|do|don't)\((\d+)?(,)?(\d+)?\)`)

	matches := re.FindAllStringSubmatch(input_lines, -1)

	for _, match := range matches {
		command_name := match[1]
		param1, _ := strconv.Atoi(match[2])
		param2, _ := strconv.Atoi(match[4])
		command := command{command_name, param1, param2}

		commands = append(commands, command)
	}
	return commands
}

type command struct {
	name   string
	param1 int
	param2 int
}
