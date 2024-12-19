package day19

import (
	"strings"
)

func Solve1(input_lines string) int {
	towels, designs := parse_input(input_lines)

	// Calc difference when sorted
	valid_designs := 0
	lookup := make(map[string]int)
	for _, design := range designs {
		if dfs(design, towels, lookup) != -1 {
			valid_designs += 1
		}
	}

	return valid_designs
}

// Returns the min towel number to make the design
func dfs(design string, towels []string, lookup map[string]int) int {
	lookup_val, exists := lookup[design]
	if exists {
		return lookup_val
	}

	if len(design) == 0 {
		return 0
	}

	min_towel_options := []int{}

	// Branch for each towel that matches start of pattern
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			min_towels_path := dfs(design[len(towel):], towels, lookup)

			if min_towels_path != -1 {
				min_towel_options = append(min_towel_options, min_towels_path)
			}
		}
	}

	// Get the min towels needed to make this pattern
	min_towels := -1
	for _, min_towel_option := range min_towel_options {
		if min_towels == -1 || min_towel_option < min_towels {
			min_towels = min_towel_option
		}
	}

	lookup[design] = min_towels
	return min_towels
}

func Solve2(input_lines string) int {
	towels, designs := parse_input(input_lines)

	// Calc difference when sorted
	println(len(towels), len(designs))

	return 1
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func parse_input(input_lines string) ([]string, []string) {

	input := strings.Split(input_lines, "\n\n")

	towels := []string{}
	for _, line := range strings.Split(input[0], ", ") {
		towels = append(towels, line)
	}

	designs := []string{}
	for _, line := range strings.Split(input[1], "\n") {
		designs = append(designs, line)
	}
	return towels, designs
}
