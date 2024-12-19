package day19

import (
	"strings"
)

func Solve1(input_lines string) int {
	towels, designs := parse_input(input_lines)

	// Calc num of possible designs
	possible_designs := 0
	lookup := make(map[string]int)
	for _, design := range designs {
		if dfs_permutations(design, towels, lookup) > 0 {
			possible_designs += 1
		}
	}

	return possible_designs
}

// Returns the number of ways to make a design from a given
// list of towels
func dfs_permutations(design string, towels []string, lookup map[string]int) int {
	lookup_val, exists := lookup[design]
	if exists {
		return lookup_val
	}

	if len(design) == 0 {
		return 1
	}

	// Branch for each towel that matches start of design
	permutations := 0
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			permutations += dfs_permutations(design[len(towel):], towels, lookup)
		}
	}

	lookup[design] = permutations
	return permutations
}

func Solve2(input_lines string) int {
	towels, designs := parse_input(input_lines)

	total_permutations := 0
	lookup := make(map[string]int)
	for _, design := range designs {
		total_permutations += dfs_permutations(design, towels, lookup)
	}

	return total_permutations
}

func parse_input(input_lines string) ([]string, []string) {

	input := strings.Split(input_lines, "\n\n")

	towels := []string{}
	towels = append(towels, strings.Split(input[0], ", ")...)

	designs := []string{}
	designs = append(designs, strings.Split(input[1], "\n")...)

	return towels, designs
}
