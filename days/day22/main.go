package day22

import (
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	nums := parse_input(input_lines)

	total_sum := 0
	for _, num := range nums {
		secret := num
		for i := 0; i < 2000; i++ {
			secret = get_next_secret(secret)
		}
		total_sum += secret
	}

	return total_sum
}

func Solve2(input_lines string) int {
	nums := parse_input(input_lines)

	// Calc difference when sorted
	println(len(nums))

	return 2
}

func get_next_secret(curr_secret int) int {
	new_num := curr_secret * 64
	new_num = new_num ^ curr_secret
	new_num = new_num % 16777216
	new_secret := new_num

	new_num = new_secret / 32
	new_num = new_num ^ new_secret
	new_num = new_num % 16777216
	new_secret = new_num

	new_num = new_secret * 2048
	new_num = new_num ^ new_secret
	new_num = new_num % 16777216
	new_secret = new_num

	return new_secret
}

func parse_input(input_lines string) []int {
	nums := []int{}

	for _, line := range strings.Split(input_lines, "\n") {

		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	return nums
}
