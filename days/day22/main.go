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
	for _, num := range nums {
		banana_lookup := Get_num_bananas(num)
		banana_lookups = append(banana_lookups, banana_lookup)
	}

	// Brute force :(
	max_total := 0
	for i := -9; i <= 9; i++ {
		for j := -9; j <= 9; j++ {
			for k := -9; k <= 9; k++ {
				for l := -9; l <= 9; l++ {
					seq := [4]int{i, j, k, l}

					total := 0
					for _, banana_lookup := range banana_lookups {
						val, exists := banana_lookup[seq]
						if exists {
							total += val
						}
					}
					if total > max_total {
						max_total = total
					}
				}
			}
		}
	}

	return max_total
}

func Get_num_bananas(secret int) map[[4]int]int {

	lookup := make(map[[4]int]int)

	diffs := [4]int{99, 99, 99, 99}
	curr := secret
	for i := 0; i < 2000; i++ {
		next := get_next_secret(curr)
		diff := (next%10 - curr%10) % 10

		diffs[i%4] = diff

		if i >= 4 {
			key := [4]int{}
			for j := range key {
				key[j] = diffs[((i%4)+j+1)%4]
			}
			_, present := lookup[key]
			if !present {
				lookup[key] = next % 10
			}
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
