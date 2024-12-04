package day04

import (
	"strings"
)

func Solve1(input_lines string) int {
	board := parse_input(input_lines)
	word_to_find := string("XMAS")
	occurrences := 0

	directions := [][]int{
		{0, -1},  // N
		{0, 1},   // S
		{1, 0},   //E
		{-1, 0},  //W
		{1, 1},   // NE
		{1, -1},  // SE
		{-1, 1},  // SW
		{-1, -1}, // NW
	}

	for y, row := range board {
		for x := range row {
			for _, dxy := range directions {
				if is_match(board, x, y, dxy[0], dxy[1], word_to_find) {
					occurrences += 1
				}
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

	// Check values are equal
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
	board := parse_input(input_lines)
	occurrences := 0

	directions := [][]int{
		{1, 1},   // NE
		{1, -1},  // SE
		{-1, 1},  // SW
		{-1, -1}, // NW
	}

	for y, row := range board {
		for x := range row {

			for _, dxy := range directions {
				if is_x_mas(board, x, y, dxy[0], dxy[1]) {
					occurrences += 1
				}
			}

		}

	}

	return occurrences
}

// Checks if the pattern is an orientation of
// M . M
// . A .
// S . S
// Where the:
// - (x_centre, y_centre) is the 'A'
// - the orientation is defined by (m_x, my)
func is_x_mas(board []string, x_centre int, y_centre int, m_x int, m_y int) bool {
	// Check in range
	y_max := len(board)
	x_max := len(board[0])
	if y_centre-1 < 0 || y_centre+1 >= y_max || x_centre-1 < 0 || x_centre+1 >= x_max {
		return false
	}

	// Check values match
	if board[y_centre][x_centre] != 'A' {
		return false
	}
	cross_pattern := string("MMSS")
	for i := range cross_pattern {
		if board[y_centre+m_y][x_centre+m_x] != cross_pattern[i] {
			return false
		}
		// Rotation 90 degrees clockwise
		m_x, m_y = m_x*0+m_y*-1, m_x*1+m_y*0
	}

	// All pass
	return true
}

func parse_input(input_lines string) []string {
	lines := strings.Split(input_lines, "\n")

	return lines
}
