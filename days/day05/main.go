package day05

import (
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	rules, updates := parse_input(input_lines)

	// Convert rules to more useful format
	rules_map := build_rules_map(rules)

	// Get valid updates
	valid_updates := [][]int{}
	for _, update := range updates {
		if update_valid(update, rules_map) {
			valid_updates = append(valid_updates, update)
		}

	}

	return get_middle_sum(valid_updates)
}

func Solve2(input_lines string) int {
	rules, updates := parse_input(input_lines)

	// Convert rules to more useful format
	rules_map := build_rules_map(rules)

	// Get invalid updates
	invalid_updates := [][]int{}
	for _, update := range updates {
		if !update_valid(update, rules_map) {
			invalid_updates = append(invalid_updates, update)
		}

	}

	// Fix invalid updates
	fixed_updates := [][]int{}
	for _, invalid_update := range invalid_updates {
		fixed_update := invalid_update

		// Loop until the update is fixed
		for !update_valid(fixed_update, rules_map) {
			for i, x := range fixed_update {
				rule, exists := rules_map[x]
				if exists {
					intersection, not_intersect := intersection(rule, fixed_update[i+1:])
					if len(intersection) > 0 {
						fixed_update = append(fixed_update[:i], intersection...)
						fixed_update = append(fixed_update, x)
						fixed_update = append(fixed_update, not_intersect...)

						break
					}
				}
			}
		}
		fixed_updates = append(fixed_updates, fixed_update)
	}

	return get_middle_sum(fixed_updates)
}

func get_middle_sum(valid_updates [][]int) int {
	middle_sum := 0
	for _, valid_update := range valid_updates {
		middle_value := valid_update[len(valid_update)/2]
		middle_sum += middle_value
	}
	return middle_sum
}

func build_rules_map(rules [][]int) map[int][]int {
	rules_map := make(map[int][]int)
	for _, rule := range rules {
		existing_rule, present := rules_map[rule[1]]

		if present {
			rules_map[rule[1]] = append(existing_rule, rule[0])
		} else {
			rules_map[rule[1]] = []int{rule[0]}
		}
	}
	return rules_map
}

func update_valid(update []int, rules_map map[int][]int) bool {
	valid := true
	for i, x := range update {
		rule, exists := rules_map[x]

		if exists && !rule_valid(update, i, rule) {
			valid = false
			break
		}
	}
	return valid
}

func rule_valid(update []int, index int, rule []int) bool {
	// Check for intersection
	intersect, _ := intersection(rule, update[index+1:])
	return len(intersect) == 0
}

func intersection(a []int, b []int) ([]int, []int) {
	set := make(map[int]struct{})
	for _, num := range a {
		set[num] = struct{}{}
	}
	intersect := []int{}
	not_intersect := []int{}
	for _, num := range b {
		if _, exists := set[num]; exists {
			intersect = append(intersect, num)
		} else {
			not_intersect = append(not_intersect, num)
		}

	}
	return intersect, not_intersect
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
