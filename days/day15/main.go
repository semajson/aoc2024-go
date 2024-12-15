package day15

import (
	"fmt"
	"slices"
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
	p1_board, moves, p1_robot := parse_input(input_lines)

	// Calc difference when sorted
	// println(board, moves, robot)
	// print_board(board, robot)
	// println(len(board), len(moves), robot.x)

	board, robot := transform_board(p1_board, p1_robot)
	print_board(board, robot)

	// Simulate
	for _, move := range moves {
		if move.dx == -1 {
			println("Move is: <")
		}
		if move.dx == 1 {
			println("Move is: >")
		}
		if move.dy == -1 {
			println("Move is: ^")
		}
		if move.dy == 1 {
			println("Move is: v")
		}

		// Curr is now a line!
		curr_line := []coord{coord{robot.x, robot.y}}

		// Find all items to move in a given direction
		boxes := []coord{}
		space := false
		for {
			// Get next line
			for i := range curr_line {
				curr_line[i] = coord{curr_line[i].x + move.dx, curr_line[i].y + move.dy}
			}
			// space_count := 0
			wall := false
			new_curr := []coord{}
			for _, curr := range curr_line {
				val, exists := board[curr]
				if !exists {
					// Found a "space"

					// Check if it is actually the right side of a box
					left := coord{curr.x - 1, curr.y}
					left_val, _ := board[left]
					if left_val == "O" {
						new_curr = append(new_curr, curr)
						if (!slices.Contains(curr_line, left)) && (move.dx == 0) {
							// Found another box!
							boxes = append(boxes, left)

							// Extend curr to left
							new_curr = append(new_curr, left)
						}
					} else {
						// It is a real space!
						// Can now ignore this curr, as it has space for it
					}

				} else if val == "#" {
					// Found wall
					wall = true
					break
				} else if val == "O" {
					// Found box
					boxes = append(boxes, coord{curr.x, curr.y})
					right := coord{curr.x + 1, curr.y}
					new_curr = append(new_curr, curr)
					if move.dx == 0 && !slices.Contains(curr_line, right) {
						// Extend curr to right
						new_curr = append(new_curr, right)
					}
				} else {
					panic("Unhitable 4")
				}

			}

			if wall {
				break
			}

			if (move.dy == 0) && len(curr_line) > 1 {
				panic("more unhitable code")
			}

			curr_line = new_curr
			if len(curr_line) == 0 {
				space = true
				break
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
		print_board(board, robot)
	}

	score_sum := 0
	for pos, val := range board {
		if val == "O" {
			score_sum += pos.x + 100*pos.y
		}
	}

	return score_sum
}

func transform_board(board map[coord]string, robot coord) (map[coord]string, coord) {
	new_board := make(map[coord]string)

	for pos, val := range board {
		switch val {
		case "#":
			new_board[coord{pos.x * 2, pos.y}] = val
			new_board[coord{pos.x*2 + 1, pos.y}] = val
		case "O":
			new_board[coord{pos.x * 2, pos.y}] = val
		}
	}
	new_robot := coord{robot.x * 2, robot.y}

	return new_board, new_robot
}

func print_board(board map[coord]string, robot coord) {
	x_max := 0
	y_max := 0

	return

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
