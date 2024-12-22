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

	banana_lookups := []map[[4]int]int{}
	all_seen := make(map[[4]int]struct{})
	for _, num := range nums {
		banana_lookup := get_banana_lookup(num, all_seen)
		banana_lookups = append(banana_lookups, banana_lookup)
	}
	max_total := 0
	for key, _ := range all_seen {
		total := 0
		for _, banana_lookup := range banana_lookups {
			val, exists := banana_lookup[key]
			if exists {
				total += val
			}
		}
		if total > max_total {
			max_total = total
		}
	}
	return max_total
}

func get_banana_lookup(secret int, all_seen map[[4]int]struct{}) map[[4]int]int {

	lookup := make(map[[4]int]int)

	diffs_buffer := [4]int{99, 99, 99, 99}
	curr := secret
	for i := 0; i < 2000; i++ {
		next := get_next_secret(curr)
		diff := (next%10 - curr%10) % 10

		diffs_buffer[i%4] = diff

		if i >= 4 {
			diffs := [4]int{}
			for j := range diffs {
				diffs[j] = diffs_buffer[((i%4)+j+1)%4]
			}
			_, present := lookup[diffs]
			if !present {
				lookup[diffs] = next % 10
			}

			all_seen[diffs] = struct{}{}
		}
		curr = next
	}

	return lookup
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
