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

const ROBOT_MOVES = true
const ROBOT_DOES_NOT_MOVE = false

func find_what_moves_p1(robot coord, move move, board map[coord]string) ([]coord, bool) {
	point_to_check := coord{robot.x, robot.y}

	// Find all items to move in a given direction
	boxes := []coord{}
	for {
		point_to_check = coord{point_to_check.x + move.dx, point_to_check.y + move.dy}

		val, exists := board[point_to_check]
		if !exists {
			// Found a space
			return boxes, ROBOT_MOVES
		} else if val == "#" {
			// Found wall
			return []coord{}, ROBOT_DOES_NOT_MOVE
		} else {
			// Found box
			boxes = append(boxes, coord{point_to_check.x, point_to_check.y})
		}
	}
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

func find_what_moves_p2(robot coord, move move, board map[coord]string) ([]coord, bool) {
	points_to_check := []coord{coord{robot.x, robot.y}}

	// Find all items to move in a given direction
	boxes := []coord{}
	for {
		// Get next line
		for i := range points_to_check {
			points_to_check[i] = coord{points_to_check[i].x + move.dx, points_to_check[i].y + move.dy}
		}

		new_points_to_check := []coord{}
		for _, point_to_check := range points_to_check {
			val, exists := board[point_to_check]
			if !exists {
				// Found a "space"

				// Check if it is actually the right side of a box
				left := coord{point_to_check.x - 1, point_to_check.y}
				left_val, _ := board[left]
				if left_val == "O" {
					// Is a box
					// Keep tracking this
					new_points_to_check = append(new_points_to_check, point_to_check)

					if (!slices.Contains(points_to_check, left)) && (move.dx == 0) {
						// Found a new box!
						boxes = append(boxes, left)

						// Need to check left point too now
						new_points_to_check = append(new_points_to_check, left)
					}
				} else {
					// Is a real space
					// Can now stop checking this point
				}
			} else if val == "#" {
				// Found wall

				// When you hit the wall, nothing moves
				return []coord{}, ROBOT_DOES_NOT_MOVE
			} else if val == "O" {
				// Found box
				boxes = append(boxes, coord{point_to_check.x, point_to_check.y})
				new_points_to_check = append(new_points_to_check, point_to_check)

				// Check if we need to check the point to the right
				right := coord{point_to_check.x + 1, point_to_check.y}
				if move.dx == 0 && !slices.Contains(points_to_check, right) {
					new_points_to_check = append(new_points_to_check, right)
				}
			} else {
				panic("Unhitable 4")
			}

		}

		// Check exist condition
		points_to_check = new_points_to_check
		if len(points_to_check) == 0 {
			// There is space to move the robot + all boxes it is attached to
			return boxes, ROBOT_MOVES
		}
	}
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
