package day07

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	equations := parse_input(input_lines)

	valid_sum := 0
	for _, equation := range equations {
		curr_total := equation.numbers[0]
		equation.numbers = equation.numbers[1:]
		if equation_valid(equation, curr_total) {
			valid_sum += equation.test_value
		}
	}

	return valid_sum
}

type Equation struct {
	test_value int
	numbers    []int
}

func equation_valid(equation Equation, curr_total int) bool {
	// Exit condition
	if curr_total > equation.test_value {
		return false
	} else if len(equation.numbers) == 0 {
		if curr_total == equation.test_value {
			return true
		} else {
			return false
		}
	}

	curr_value := equation.numbers[0]

	// Add branch
	if equation_valid(Equation{equation.test_value, equation.numbers[1:]}, curr_total+curr_value) {
		return true
	}

	// Multiple branch
	return equation_valid(Equation{equation.test_value, equation.numbers[1:]}, curr_total*curr_value)
}

func Solve2(input_lines string) int {
	equations := parse_input(input_lines)

	valid_sum := 0
	for _, equation := range equations {
		curr_total := equation.numbers[0]
		equation.numbers = equation.numbers[1:]
		if equation_valid_2(equation, curr_total) {
			valid_sum += equation.test_value
		}
	}

	return valid_sum
}

func equation_valid_2(equation Equation, curr_total int) bool {
	// Exit condition
	if curr_total > equation.test_value {
		return false
	} else if len(equation.numbers) == 0 {
		if curr_total == equation.test_value {
			return true
		} else {
			return false
		}
	}

	curr_value := equation.numbers[0]

	// Add branch
	if equation_valid_2(Equation{equation.test_value, equation.numbers[1:]}, curr_total+curr_value) {
		return true
	}

	// Multiple branch
	if equation_valid_2(Equation{equation.test_value, equation.numbers[1:]}, curr_total*curr_value) {
		return true
	}

	// || branch
	new_total, _ := strconv.Atoi(fmt.Sprint(curr_total) + fmt.Sprint(curr_value))
	return equation_valid_2(Equation{equation.test_value, equation.numbers[1:]}, new_total)
}

func parse_input(input_lines string) []Equation {
	equations := []Equation{}

	for _, line := range strings.Split(input_lines, "\n") {
		equation_raw := strings.Split(line, ":")

		test_number, _ := strconv.Atoi(equation_raw[0])

		numbers_raw := strings.Fields(equation_raw[1])
		numbers := make([]int, len(numbers_raw))
		for i, number_raw := range numbers_raw {
			numbers[i], _ = strconv.Atoi(number_raw)
		}

		equations = append(equations, Equation{test_number, numbers})
	}
	return equations
}
