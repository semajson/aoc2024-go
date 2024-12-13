package day11

import (
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	nums := parse_input(input_lines)

	lookup := make(map[[2]int]int)

	count := 0
	for _, num := range nums {
		delta := get_count(num, 25, lookup)

		count += delta
	}

	return count
}

func Solve2(input_lines string) int {
	nums := parse_input(input_lines)

	lookup := make(map[[2]int]int)

	count := 0
	for _, num := range nums {
		delta := get_count(num, 75, lookup)

		count += delta
	}

	return count
}

func get_count(num int, blinks int, lookup map[[2]int]int) int {
	// Base case
	if blinks == 0 {
		return 1
	}

	// Lookups for speed
	key := [2]int{num, blinks}
	value, exists := lookup[key]
	if exists {
		return value
	}

	// Do blink
	count := 0
	num_str := strconv.Itoa(num)
	if num == 0 {
		count = get_count(1, blinks-1, lookup)
	} else if len(num_str)%2 == 0 {
		// Even num of digits
		left := num_str[:len(num_str)/2]
		left_num, _ := strconv.Atoi(left)
		right := num_str[len(num_str)/2:]
		right_num, _ := strconv.Atoi(right)

		count = get_count(left_num, blinks-1, lookup) + get_count(right_num, blinks-1, lookup)
	} else {
		count = get_count(num*2024, blinks-1, lookup)
	}

	lookup[key] = count
	return count
}

func parse_input(input_lines string) []int {
	nums := []int{}
	for _, num_raw := range strings.Split(input_lines, " ") {
		num, _ := strconv.Atoi(num_raw)

		nums = append(nums, num)

	}
	return nums
}
