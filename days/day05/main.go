package day05

import (
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	rules, updates := parse_input(input_lines)

	// Convert rules to more useful format
	// Maybe should be a map to a set rather than a slice?
	rules_map := make(map[int][]int)
	for _, rule := range rules {
		existing_rule, present := rules_map[rule[1]]

		if present {
			rules_map[rule[1]] = append(existing_rule, rule[0])
		} else {
			rules_map[rule[1]] = []int{rule[0]}
		}
	}

	// Get valid updates
	valid_updates := [][]int{}
	for _, update := range updates {
		valid := true
		for i, x := range update {
			rule, exists := rules_map[x]

			if exists && !update_valid(update, i, rule) {
				valid = false
				break
			}
		}
		if valid {
			valid_updates = append(valid_updates, update)
		}

	}

	// Get middle values of updates
	middle_sum := 0
	for _, valid_update := range valid_updates {
		middle_value := valid_update[len(valid_update)/2]
		middle_sum += middle_value
	}

	return middle_sum
}

func update_valid(update []int, index int, rule []int) bool {
	// Check for intersection
	set := make(map[int]struct{})
	for _, num := range rule {
		set[num] = struct{}{}
	}

	for _, num := range update[index+1:] {
		if _, exists := set[num]; exists {
			return false
		}
	}
	return true
}

func Solve2(input_lines string) int {
	return 1
}

func parse_input(input_lines string) ([][]int, [][]int) {
	parts := strings.Split(input_lines, "\n\n")

	// Process rules
	rules := [][]int{}
	for _, line := range strings.Split(parts[0], "\n") {
		nums := strings.Split(line, "|")
		rule := make([]int, 2)
		for i, num := range nums {
			rule[i], _ = strconv.Atoi(num)
		}
		rules = append(rules, rule)
	}

	// Process updates
	updates := [][]int{}
	for _, line := range strings.Split(parts[1], "\n") {
		nums := strings.Split(line, ",")
		update := make([]int, len(nums))
		for i, num := range nums {
			update[i], _ = strconv.Atoi(num)
		}
		updates = append(updates, update)
	}

	return rules, updates
}
