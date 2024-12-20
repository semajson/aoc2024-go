package day20

import (
	"strings"
)

func Solve1(input_lines string) int {
	return Solve(input_lines, 100, 2)
}

func Solve2(input_lines string) int {
	return Solve(input_lines, 100, 20)
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Solve(input_lines string, save_threshold int, cheat_range int) int {
	board, start, end := parse_input(input_lines)

	// Build useful lookups
	distances_from_start := distance_map(board, start, end)
	distances_to_end := distance_map(board, end, start)
	end_distance := distances_from_start[end]

	// Loop through the path, at each point using a cheat to
	// jump to another section on it.
	// If the cheat shortens the start->end path to be within the required "saved threshold",
	// then count it.
	cheats_within_threshold := 0
	for cheat_start, start_to_cheat_start := range distances_from_start {
		if cheat_start == end {
			continue
		}
		for dx := -cheat_range; dx <= cheat_range; dx++ {
			for dy := -cheat_range; dy <= cheat_range; dy++ {
				cheat_start_to_cheat_end := AbsInt(dx) + AbsInt(dy)
				cheat_end := coord{x: cheat_start.x + dx, y: cheat_start.y + dy}
				if cheat_start_to_cheat_end <= cheat_range {
					val, exists := board[cheat_end]

					if exists && val == "." {
						// Have a valid cheat route, now just check if saves enough time
						cheat_end_to_end := distances_to_end[cheat_end]
						distance_with_cheat := start_to_cheat_start + cheat_start_to_cheat_end + cheat_end_to_end

						if (end_distance - distance_with_cheat) >= save_threshold {
							cheats_within_threshold += 1
						}
					}
				}
			}
		}
	}

	return cheats_within_threshold
}

// Use DFS to build a map of shortest distances from the "start" coord until the
// end coord is reached
func distance_map(board map[coord]string, start coord, end coord) map[coord]int {
	current := []coord{start}
	seen := make(map[coord]int)

	depth := 0
	for len(current) > 0 {
		new_current := []coord{}
		for _, current_node := range current {
			_, visited := seen[current_node]
			if visited {
				continue
			}
			seen[current_node] = depth

			// Exit
			if current_node == end {
				return seen
			}

			// Branch
			for _, neighbour := range current_node.get_neighbours() {
				val, exists := board[neighbour]

				if exists && val == "." {
					new_current = append(new_current, neighbour)
				}
			}
		}
		depth += 1
		current = new_current
	}
	panic("Didn't find answer")
}

type coord struct {
	x int
	y int
}

func (pos coord) get_neighbours() []coord {
	return []coord{
		{pos.x + 1, pos.y},
		{pos.x - 1, pos.y},
		{pos.x, pos.y + 1},
		{pos.x, pos.y - 1}}
}

func parse_input(input_lines string) (map[coord]string, coord, coord) {
	board := make(map[coord]string)
	start := coord{}
	end := coord{}

	for y, line := range strings.Split(input_lines, "\n") {

		for x, char := range line {
			switch char {
			case '.':
				board[coord{x, y}] = "."
			case '#':
				board[coord{x, y}] = "#"
			case 'S':
				board[coord{x, y}] = "."
				start = coord{x, y}
			case 'E':
				board[coord{x, y}] = "."
				end = coord{x, y}
			}
		}
	}
	return board, start, end
}
