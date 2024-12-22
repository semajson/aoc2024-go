package day22

import (
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	nums := parse_input(input_lines)

	// Calc difference when sorted
	total_sum := 0

	for _, num := range nums {
		secret := Secret{
			num:   num,
			index: 0,
		}
		for i := 0; i < 2000; i++ {
			secret.evolve()
		}
		total_sum += secret.num
	}

	return total_sum
}

func Solve2(input_lines string) int {
	nums := parse_input(input_lines)

	// Calc difference when sorted
	println(len(nums))

	return 2
}

type Secret struct {
	num   int
	index int
}

func (s *Secret) evolve() {
	// switch s.index {
	// case 0:
	// 	new_num := s.num * 64
	// 	new_num = new_num ^ s.num
	// 	new_num = new_num % 16777216
	// 	s.num = new_num
	// case 1:
	// 	new_num := s.num / 32
	// 	new_num = new_num ^ s.num
	// 	new_num = new_num % 16777216
	// 	s.num = new_num
	// case 2:
	// 	new_num := s.num * 2048
	// 	new_num = new_num ^ s.num
	// 	new_num = new_num % 16777216
	// 	s.num = new_num
	// }

	new_num := s.num * 64
	new_num = new_num ^ s.num
	new_num = new_num % 16777216
	s.num = new_num

	new_num = s.num / 32
	new_num = new_num ^ s.num
	new_num = new_num % 16777216
	s.num = new_num

	new_num = s.num * 2048
	new_num = new_num ^ s.num
	new_num = new_num % 16777216
	s.num = new_num

	s.index += 1
	s.index %= 3
}

func parse_input(input_lines string) []int {
	nums := []int{}

	for _, line := range strings.Split(input_lines, "\n") {

		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	return nums
}
