package day04

import (
	"strings"
)

func Solve1(input_lines string) int {
	// Process the input
	board := parse_input(input_lines)
	word_to_find := string("XMAS")
	occurrences := 0

	for y, row := range board {
		for x := range row {
			// North
			if is_match(board, x, y, 0, -1, word_to_find) {
				occurrences += 1
			}

			// South
			if is_match(board, x, y, 0, 1, word_to_find) {
				occurrences += 1
			}

			// East
			if is_match(board, x, y, 1, 0, word_to_find) {
				occurrences += 1
			}

			// West
			if is_match(board, x, y, -1, 0, word_to_find) {
				occurrences += 1
			}

			// NE
			if is_match(board, x, y, 1, -1, word_to_find) {
				occurrences += 1
			}

			// SE
			if is_match(board, x, y, 1, 1, word_to_find) {
				occurrences += 1
			}

			// SW
			if is_match(board, x, y, -1, 1, word_to_find) {
				occurrences += 1
			}

			// NW
			if is_match(board, x, y, -1, -1, word_to_find) {
				occurrences += 1
			}

		}

	}

	return occurrences
}

func is_match(board []string, x_start int, y_start int, dx int, dy int, word_to_find string) bool {
	// Check in range
	word_to_find_len := len(word_to_find)
	y_max := len(board)
	x_max := len(board[0])
	x_end := x_start + (word_to_find_len-1)*dx
	y_end := y_start + (word_to_find_len-1)*dy
	if y_end < 0 || y_end >= y_max || x_end < 0 || x_end >= x_max {
		return false
	}

	// Check values
	x := x_start
	y := y_start
	match := true
	for i := range word_to_find {
		if board[y][x] != word_to_find[i] {
			match = false
			break
		}
		y += dy
		x += dx
	}
	return match
}

func Solve2(input_lines string) int {
	// Process the input
	return 1
}

func parse_input(input_lines string) []string {
	lines := strings.Split(input_lines, "\n")

	return lines
}
