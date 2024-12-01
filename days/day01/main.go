package day01

import (
	"sort"
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	// Process the input
	left_nums, right_nums := parse_input(input_lines)

	// Calc difference when sorted
	sort.Ints(left_nums)
	sort.Ints(right_nums)

	diff_sum := 0

	for i := range left_nums {
		diff_sum += abs(left_nums[i], right_nums[i])
	}

	return diff_sum
}

func Solve2(input_lines string) int {
	// Process the input
	left_nums, right_nums := parse_input(input_lines)

	// Count numbers in right
	right_count := make(map[int]int)
	for _, num := range right_nums {
		count, exists := right_count[num]
		if exists {
			right_count[num] = count + 1
		} else {
			right_count[num] = 1
		}
	}

	// Calc similarity score
	similarity_score := 0
	for _, num := range left_nums {
		count, exists := right_count[num]
		if exists {
			similarity_score += count * num
		}
	}

	return similarity_score
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func parse_input(input_lines string) ([]int, []int) {
	left_nums := []int{}
	right_nums := []int{}

	for _, line := range strings.Split(input_lines, "\n") {
		nums := strings.Fields(line)
		left_num, _ := strconv.Atoi(nums[0])
		left_nums = append(left_nums, left_num)

		right_num, _ := strconv.Atoi(nums[1])
		right_nums = append(right_nums, right_num)
	}
	return left_nums, right_nums
}
