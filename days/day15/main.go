package day15

import (
	"fmt"
	"slices"
	"strings"
)

func Solve1(input_lines string) int {
	board, moves, robot := parse_input(input_lines)

	// Simulate
	for _, move := range moves {
		boxes, robot_moves := find_what_moves_p1(robot, move, board)

		if robot_moves {
			move_boxes(boxes, board, move)
			robot = coord{robot.x + move.dx, robot.y + move.dy}
		}
		// print_board(board, robot)
	}

	return calc_score_sum(board)
}

func find_what_moves_p1(robot coord, move move, board map[coord]string) ([]coord, bool) {
	curr := coord{robot.x, robot.y}

	// Find all items to move in a given direction
	boxes := []coord{}
	robot_moves := false
	for {
		curr = coord{curr.x + move.dx, curr.y + move.dy}

		val, exists := board[curr]
		if !exists {
			// Found a space
			robot_moves = true
			break
		} else if val == "#" {
			// Found wall
			break
		} else {
			// Found box
			boxes = append(boxes, coord{curr.x, curr.y})
		}
	}
	return boxes, robot_moves
}

func calc_score_sum(board map[coord]string) int {
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
	board, robot := transform_board_for_p2(p1_board, p1_robot)

	// Simulate
	// print_board(board, robot)
	for _, move := range moves {
		// print_move(move)
		boxes, robot_moves := find_what_moves_p2(robot, move, board)

		if robot_moves {
			move_boxes(boxes, board, move)
			robot = coord{robot.x + move.dx, robot.y + move.dy}
		}
		// print_board(board, robot)
	}

	return calc_score_sum(board)
}

func find_what_moves_p2(robot coord, move move, board map[coord]string) ([]coord, bool) {
	curr_line := []coord{coord{robot.x, robot.y}}

	// Find all items to move in a given direction
	boxes := []coord{}
	for {
		// Get next line
		for i := range curr_line {
			curr_line[i] = coord{curr_line[i].x + move.dx, curr_line[i].y + move.dy}
		}

		wall := false
		new_curr_line := []coord{}
		for _, curr := range curr_line {
			val, exists := board[curr]
			if !exists {
				// Found a "space"

				// Check if it is actually the right side of a box
				left := coord{curr.x - 1, curr.y}
				left_val, _ := board[left]
				if left_val == "O" {
					// Is a box
					new_curr_line = append(new_curr_line, curr)

					if (!slices.Contains(curr_line, left)) && (move.dx == 0) {
						// Found a new box!
						boxes = append(boxes, left)

						// Extend curr to left one
						new_curr_line = append(new_curr_line, left)
					}
				} else {
					// It is a real space!
					// Can now ignore this curr, as there is space for this
				}
			} else if val == "#" {
				// Found wall
				wall = true
				break
			} else if val == "O" {
				// Found box
				boxes = append(boxes, coord{curr.x, curr.y})

				// Check if we need to extend curr to right
				right := coord{curr.x + 1, curr.y}
				new_curr_line = append(new_curr_line, curr)
				if move.dx == 0 && !slices.Contains(curr_line, right) {
					// Extend curr to right
					new_curr_line = append(new_curr_line, right)
				}
			} else {
				panic("Unhitable 4")
			}

		}

		// Check exist conditions

		if wall {
			// Hit the wall, nothing moves
			robot_moves := false
			return []coord{}, robot_moves
		}

		curr_line = new_curr_line
		if len(curr_line) == 0 {
			// There is space to move the robot + all boxes it is attached to
			robot_moves := true
			return boxes, robot_moves
		}
	}
}

func move_boxes(boxes []coord, board map[coord]string, move move) {
	// Move All found boxes by 1 in the direction
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
}

func print_move(move move) {
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
}

func transform_board_for_p2(board map[coord]string, robot coord) (map[coord]string, coord) {
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
