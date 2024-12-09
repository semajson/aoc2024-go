package day09

import (
	"fmt"
	"log"
)

const FREE = -1

func Solve1(input_lines string) int {
	filesystem := parse_input(input_lines)

	blocks := calc_blocks(filesystem)

	last_possible_file_block := len(blocks) - 1
	for i := range blocks {
		// Find the first free block
		block := blocks[i]
		if block == FREE {
			// Now find last file block
			for j := last_possible_file_block; j > i; j-- {
				end_block := blocks[j]
				if end_block != FREE {
					// Swap them around
					blocks[i] = end_block
					blocks[j] = FREE
					last_possible_file_block = j - 1

					break
				}
			}
		}
	}

	return calc_checksum(blocks)
}

func calc_blocks(filesystem []int) []int {
	blocks := []int{}
	for i, num := range filesystem {
		if i%2 == 0 {
			// File
			file_num := i / 2
			for i := 0; i < num; i++ {
				blocks = append(blocks, file_num)
			}
		} else {
			// Free space
			for i := 0; i < num; i++ {
				blocks = append(blocks, FREE)
			}
		}
	}
	return blocks
}

func calc_checksum(blocks []int) int {
	checksum := 0
	for i, block := range blocks {
		if block != FREE {
			checksum += i * block
		}
	}
	return checksum
}

func debug_print(blocks []int) {
	if true {
		return
	}
	fmt.Printf("\n")
	for _, block := range blocks {
		if block == FREE {
			fmt.Printf(".")

		} else {
			fmt.Printf("%d", block)

		}
	}
	fmt.Printf("\n")
}

func Solve2(input_lines string) int {
	filesystem := parse_input(input_lines)
	blocks := calc_blocks(filesystem)

	debug_print(blocks)

	curr_file_block := FREE
	curr_file_len := 0
	moved_blocks := make(map[int]struct{})
	for i := len(blocks) - 1; i >= 0; i-- {
		block := blocks[i]
		if block == curr_file_block {
			curr_file_len += 1
		} else {
			_, already_moved := moved_blocks[curr_file_block]
			if curr_file_block != FREE && !already_moved {
				// Found whole file

				// Try to move file
				free_index := First_free_space(blocks, curr_file_len, i)
				if free_index > 0 {
					// Copy file to the found free space
					for j := free_index; j < free_index+curr_file_len; j++ {
						blocks[j] = curr_file_block
					}

					// Free the current space
					for j := i + 1; j < i+1+curr_file_len; j++ {
						blocks[j] = FREE
					}
					debug_print(blocks)
				}
			}

			moved_blocks[curr_file_block] = struct{}{}
			curr_file_block = block
			curr_file_len = 1
		}
	}
	debug_print(blocks)

	return calc_checksum(blocks)
}

func First_free_space(blocks []int, space_required int, curr_index int) int {
	if space_required <= 0 {
		return -1
	}

	curr_space := 0

	for i := 0; i <= curr_index; i++ {
		if blocks[i] == FREE {
			curr_space += 1
			if curr_space == space_required {
				// Found enough space!
				return i + 1 - space_required
			}
		} else {
			curr_space = 0
		}
	}
	return -1
}

func parse_input(input_lines string) []int {
	filesystem := []int{}

	for _, char := range input_lines {
		num := int(char) - int('0')

		if num > 10 {
			err := fmt.Errorf("the num found: %d is greater than 10", num)
			log.Fatal(err)
		}
		filesystem = append(filesystem, num)
	}
	return filesystem
}
