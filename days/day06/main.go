package day06

import (
	"strings"
)

var DIRECTIONS = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func Solve1(input_lines string) int {
	board, guard := parse_input(input_lines)

	coords_visited := make(map[[2]int]struct{})
	for guard_on_board(guard, board) {
		coord := [2]int{guard.x, guard.y}
		_, visited := coords_visited[coord]
		if !visited {
			coords_visited[coord] = struct{}{}
		}

		move_guard(&guard, board)
	}
	return len(coords_visited)
}

func Solve2(input_lines string) int {
	board, guard_initial := parse_input(input_lines)

	num_repeated_patterns := 0

	for y, row := range board {
		for x := range row {
			if guard_initial.x == x && guard_initial.y == y {
				continue
			}
			if board[y][x] != '.' {
				continue
			}

			// Change pattern to add block
			// Maybe better way to avoid copying idk
			new_board := make([]string, len(board))
			copy(new_board, board)
			new_row := []rune(new_board[y])
			new_row[x] = '#'
			new_board[y] = string(new_row)

			// Check for repeats
			guard := guard_initial
			if does_repeat(guard, new_board) {
				num_repeated_patterns += 1
			}
		}
	}

	return num_repeated_patterns
}

func does_repeat(guard player, board []string) bool {
	guard_history := make(map[player]struct{})
	for guard_on_board(guard, board) {
		_, repeated := guard_history[guard]
		if repeated {
			return true
		} else {
			guard_history[guard] = struct{}{}
		}

		move_guard(&guard, board)
	}
	return false
}

type player struct {
	x         int
	y         int
	dir_index int
}

func guard_on_board(guard player, board []string) bool {
	return coord_on_board(guard.x, guard.y, board)
}

func coord_on_board(x int, y int, board []string) bool {
	return x >= 0 && x < len(board[0]) && y >= 0 && y < len(board)
}

func move_guard(guard *player, board []string) {
	for {
		new_x := guard.x + DIRECTIONS[guard.dir_index][0]
		new_y := guard.y + DIRECTIONS[guard.dir_index][1]

		if coord_on_board(new_x, new_y, board) && board[new_y][new_x] == '#' {
			guard.dir_index = (guard.dir_index + 1) % len(DIRECTIONS)
		} else {
			guard.x = new_x
			guard.y = new_y
			break
		}
	}

}

func parse_input(input_lines string) ([]string, player) {
	board := strings.Split(input_lines, "\n")

	var guard player
	for y, row := range board {
		for x := range row {
			if board[y][x] == '^' {
				guard = player{x: x, y: y, dir_index: 3}
				break
			}
		}
	}

	// Remove guard from board
	row := []rune(board[guard.y])
	row[guard.x] = '.'
	board[guard.y] = string(row)

	return board, guard
}
