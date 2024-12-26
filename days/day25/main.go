package day25

import (
	"strings"
)

func Solve1(input_lines string) int {
	keys, locks := parse_input(input_lines)

	// Just brute check each key/lock combo
	valid_combos := 0
	for _, key := range keys {
		for _, lock := range locks {
			valid := true
			for i := range key {
				if (key[i] + lock[i]) > 5 {
					valid = false
					break
				}
			}
			if valid {
				valid_combos += 1
			}
		}
	}

	return valid_combos
}

func Solve2(input_lines string) int {
	keys, locks := parse_input(input_lines)

	println(len(keys), len(locks))

	return 1
}

func parse_input(input_lines string) ([][]int, [][]int) {
	keys := [][]int{}
	locks := [][]int{}

	for _, lines := range strings.Split(input_lines, "\n\n") {
		lines := strings.Split(lines, "\n")

		if strings.Contains(lines[0], "#") {
			// Lock
			lock := []int{}
			for i := range lines[0] {
				length := 0
				for j := 1; j < len(lines); j++ {
					if lines[j][i] == '#' {
						length += 1
					} else {
						break
					}
				}
				lock = append(lock, length)
			}
			locks = append(locks, lock)
		} else {
			// Key
			key := []int{}
			for i := range lines[0] {
				length := 0
				for j := len(lines) - 2; j >= 0; j-- {
					if lines[j][i] == '#' {
						length += 1
					} else {
						break
					}
				}
				key = append(key, length)
			}
			keys = append(keys, key)
		}
	}
	return keys, locks
}
