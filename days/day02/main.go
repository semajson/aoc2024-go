package day02

import (
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	// Process the input
	reports := parse_input(input_lines)

	safe_count := 0

	for _, report := range reports {

		if is_safe(report) {
			safe_count += 1
		}
	}

	return safe_count
}

func is_safe(report []int) bool {
	safe := true
	curr := report[0]
	increasing := report[1] > curr

	for _, level := range report[1:] {

		if abs(curr, level) > 3 {
			safe = false
			break
		}
		if increasing && level <= curr {
			safe = false
			break

		} else if !increasing && level >= curr {
			safe = false
			break
		}
		curr = level
	}
	return safe
}

func Solve2(input_lines string) int {
	reports := parse_input(input_lines)

	safe_count := 0

	for _, report := range reports {

		if is_safe_with_dampener(report) {
			safe_count += 1
		}
	}

	return safe_count
}

func is_safe_with_dampener(report []int) bool {
	safe := true
	curr := report[0]
	increasing := report[1] > curr
	num_removed := 0

	for _, level := range report[1:] {

		if abs(curr, level) > 3 {
			safe = false
		}
		if increasing && level <= curr {
			safe = false

		} else if !increasing && level >= curr {
			safe = false
		}

		if !safe {
			if num_removed < 1 {
				safe = true
				num_removed += 1
			} else {
				break
			}
		} else {
			curr = level
		}

	}

	if !safe {
		// Deal with case where you could remove either the first or second element
		// What a horrible hack lol
		return is_safe(report[1:]) || is_safe(append(report[:1], report[2:]...))
	} else {
		return true
	}
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func parse_input(input_lines string) [][]int {
	lines := strings.Split(input_lines, "\n")

	reports := make([][]int, len(lines))

	for i, line := range lines {
		levels := strings.Fields(line)
		report := make([]int, len(levels))
		for j, level := range levels {
			num, _ := strconv.Atoi(level)
			report[j] = num
		}
		reports[i] = report
	}
	return reports
}
