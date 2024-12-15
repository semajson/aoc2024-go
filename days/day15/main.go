package day15

import (
	"fmt"
	"strings"
)

func Solve1(input_lines string) int {
	board, moves, robot := parse_input(input_lines)

	// Calc difference when sorted
	// println(board, moves, robot)
	// print_board(board, robot)
	// println(len(board), len(moves), robot.x)

	// Simulate
	for _, move := range moves {
		curr := coord{robot.x, robot.y}

		// Find all items to move in a given direction
		boxes := []coord{}
		space := false
		for {
			curr = coord{curr.x + move.dx, curr.y + move.dy}

			val, exists := board[curr]
			if !exists {
				// Found a space
				space = true
				break
			} else if val == "#" {
				// Found wall
				break
			} else {
				// Found box
				boxes = append(boxes, coord{curr.x, curr.y})
			}
		}

		if space {
			// Move all found boxes by 1 in the direction
			for i := len(boxes) - 1; i >= 0; i-- {
				// Remove box
				box_pos := boxes[i]
				box_val, exists := board[box_pos]
				if box_val != "O" {
					panic("unhittable code 3")
				}
				if !exists {
					panic("unhittable code 1")
				}
				delete(board, box_pos)

				// Re-add box
				box_pos = coord{box_pos.x + move.dx, box_pos.y + move.dy}
				_, occupied := board[box_pos]
				if occupied {
					panic("unhittable code 2")
				}
				board[box_pos] = box_val
			}

			// Move robot
			robot = coord{robot.x + move.dx, robot.y + move.dy}
		}
		// print_board(board, robot)
	}

	score_sum := 0
	for pos, val := range board {
		if val == "O" {
			score_sum += pos.x + 100*pos.y
		}
	}

	return score_sum
}

func Solve2(input_lines string) int {
	board, moves, robot := parse_input(input_lines)

	// Calc difference when sorted
	println(len(board), len(moves), robot.x)

	return 1
}

func print_board(board map[coord]string, robot coord) {
	x_max := 0
	y_max := 0

	for pos, _ := range board {
		x_max = max(x_max, pos.x)
		y_max = max(y_max, pos.y)
	}

	for y := 0; y <= y_max; y++ {
		for x := 0; x <= x_max; x++ {
			val, exists := board[coord{x, y}]
			if exists {
				fmt.Printf("%s", val)
			} else {
				if x == robot.x && y == robot.y {
					fmt.Printf("@")
				} else {
					fmt.Printf(".")
				}
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

}

type coord struct {
	x int
	y int
}

type move struct {
	dx int
	dy int
}

func parse_input(input_lines string) (map[coord]string, []move, coord) {
	input_parts := strings.Split(input_lines, "\n\n")

	// Parse board
	var robot coord
	board := make(map[coord]string)
	for y, line := range strings.Split(input_parts[0], "\n") {
		for x, val := range line {
			switch val {
			case '@':
				robot = coord{x, y}
			case '#':
				board[coord{x, y}] = string(val)
			case 'O':
				board[coord{x, y}] = string(val)
			case '.':
				continue
			default:
				panic("Unexpected input")
			}
		}
	}

	// Parse moves
	moves := []move{}
	for _, line := range strings.Split(input_parts[1], "\n") {
		for _, move_raw := range line {
			switch move_raw {
			case '^':
				moves = append(moves, move{0, -1})
			case '>':
				moves = append(moves, move{1, 0})
			case 'v':
				moves = append(moves, move{0, 1})
			case '<':
				moves = append(moves, move{-1, 0})
			default:
				panic("Unexpected move")
			}
		}
	}
	return board, moves, robot
}
