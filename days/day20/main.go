package day20

import (
	"fmt"
	"strings"
)

func Solve1(input_lines string) int {
	board, start, end := parse_input(input_lines)

	println(len(board), start.x, end.x)
	distance_from_start := solve_no_cheat(board, start, end)
	distance_to_end := solve_no_cheat(board, end, start)
	end_distance, _ := distance_from_start[end]

	// Cheats
	cheat_count := 0
	cheat_range := 2
	cheats_within := 100
	for cheat_start, moves := range distance_from_start {
		for dx := -cheat_range; dx <= cheat_range; dx++ {
			for dy := -cheat_range; dy <= cheat_range; dy++ {
				cheat_start_to_cheat_end := AbsInt(dx) + AbsInt(dy)
				cheat_end := coord{x: cheat_start.x + dx, y: cheat_start.y + dy}
				if cheat_start_to_cheat_end <= cheat_range {
					val, exists := board[cheat_end]

					if exists && val == "." {
						move_to_end, exists_2 := distance_to_end[cheat_end]

						if !exists_2 {
							panic("Error")
						}

						distance_with_cheat := moves + cheat_start_to_cheat_end + move_to_end
						if (end_distance - distance_with_cheat) >= cheats_within {
							cheat_count += 1
						}
					}
				}
			}
		}
	}

	return cheat_count

	// total_count := 0
	// for _, moves := range cheat_moves {

	// 	if (best_no_cheat - moves) >= 20 {
	// 		total_count += 1
	// 	}
	// }

	// return total_count

}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve_no_cheat(board map[coord]string, start coord, end coord) map[coord]int {
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

func bfs_get_route_2(board map[coord]string, start coord, end coord, best_no_cheat int) map[[2]coord]int {
	cheat_shortest_paths := make(map[[2]coord]int)

	// Get max
	x_max := 0
	y_max := 0
	for pos, _ := range board {
		x_max = max(x_max, pos.x)
		y_max = max(y_max, pos.y)
	}

	for y := 0; y <= y_max; y++ {
		println("Trying y:", y)
		for x := 0; x <= x_max; x++ {
			cheat_start := coord{x, y}
			cheat_start_val, exists := board[cheat_start]

			if exists && cheat_start_val == "#" {
				board[cheat_start] = "."

				shortest_path_with_cheat := bfs(board, start, end)

				if shortest_path_with_cheat < best_no_cheat {
					key := [2]coord{cheat_start, cheat_start}
					cheat_shortest_paths[key] = shortest_path_with_cheat
				}

				board[cheat_start] = "#"

				// for _, cheat_end := range start.get_neighbours() {
				// 	cheat_end_val, exists_2 := board[end]

				// 	if exists_2 && cheat_end_val == "." {
				// 		// Found cheat candidate

				// 		// board[cheat_end] = "."

				// 		print_board(board, start)

				// 		shortest_path_with_cheat := bfs(board, start, end)

				// 		if shortest_path_with_cheat < no_cheat_shortest_path {
				// 			key := [2]coord{cheat_start, cheat_end}
				// 			cheat_shortest_paths[key] = shortest_path_with_cheat
				// 		}

				// 		board[cheat_start] = "#"
				// 		// board[cheat_end] = "#"
				// 	}
				// }
			}
		}
	}

	return cheat_shortest_paths
}

func bfs(board map[coord]string, start coord, end coord) int {
	current := []coord{start}
	seen := make(map[coord]struct{})

	depth := 0
	for len(current) > 0 {
		new_current := []coord{}
		for _, current_node := range current {
			_, visited := seen[current_node]
			if visited {
				continue
			}
			seen[current_node] = struct{}{}

			// Exit
			if current_node == end {
				return depth
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
	return depth
}

func bfs_get_route(board map[coord]string, start coord, end coord, _ int) map[[2]coord]int {
	current := []node{node{start, 0, 0, 0, coord{-1, -1}, coord{-1, -1}}}

	seen := make(map[node_lookup]struct{})
	no_cheat_shortest_path := make(map[coord]int)

	cheat_quickest_path := make(map[[2]coord]int)

	max_cheats := 1
	max_cheat_seconds := 2

	for len(current) > 0 {
		current_node := current[0]
		current = current[1:]

		if len(seen)%1000000 == 0 {
			// println("len(seen): ", len(seen), ". Len(current): ", len(current), ". len(cheat_quickest_path):",len(cheat_quickest_path))
		}

		// Pruning
		seen_lookup := node_lookup{
			pos:                     current_node.pos,
			moves:                   current_node.moves,
			cheats_used:             current_node.cheats_used,
			cheat_seconds_remaining: current_node.cheat_seconds_remaining}
		_, visited := seen[seen_lookup]
		if visited {
			continue
		}
		seen[seen_lookup] = struct{}{}

		moves, visited_2 := no_cheat_shortest_path[current_node.pos]
		if visited_2 && current_node.moves >= moves {
			// Ignore as we are on a slower path than the non-cheat path
			continue
		}
		if current_node.cheats_used == 0 {
			no_cheat_shortest_path[current_node.pos] = current_node.moves
		}

		cheat_key := [2]coord{current_node.cheat_start, current_node.cheat_end}
		_, done_cheat := cheat_quickest_path[cheat_key]

		if done_cheat {
			continue
		}

		// Exit condition
		if current_node.pos == end {
			// println("at exit")
			if current_node.cheats_used == 0 {
				// best = current_node.moves
				// Lets exist
				break
			} else if current_node.cheats_used == 1 {
				if current_node.cheat_end.x == -1 {
					panic("un-hitable")
					// if current_node.cheat_seconds_remaining == 0 {
					// 	panic("un-hitable code")
					// }
					// current_node.cheat_end == current_node.pos
				}
				cheat_quickest_path[cheat_key] = current_node.moves

			}
			continue
		}

		// Branch
		for _, neighbour := range current_node.pos.get_neighbours() {
			val, exists := board[neighbour]

			if !exists {
				continue
			}

			new_cheat_seconds_remaining := 0
			new_cheat_end := current_node.cheat_end
			if current_node.cheat_seconds_remaining > 0 {
				new_cheat_seconds_remaining = current_node.cheat_seconds_remaining - 1
				new_cheat_end = neighbour
			}

			switch val {
			case ".":
				current = append(current,
					node{
						pos:                     neighbour,
						cheats_used:             current_node.cheats_used,
						cheat_seconds_remaining: new_cheat_seconds_remaining,
						cheat_start:             current_node.cheat_start,
						cheat_end:               new_cheat_end,
						moves:                   current_node.moves + 1})
			case "#":
				if current_node.cheat_seconds_remaining > 1 {
					// Continue current cheat
					current = append(current,
						node{
							pos:                     neighbour,
							cheats_used:             current_node.cheats_used,
							cheat_seconds_remaining: new_cheat_seconds_remaining,
							cheat_start:             current_node.cheat_start,
							cheat_end:               neighbour,
							moves:                   current_node.moves + 1})
				} else if current_node.cheats_used < max_cheats {
					// Start cheat
					current = append(current,
						node{
							pos:                     neighbour,
							cheats_used:             current_node.cheats_used + 1,
							cheat_seconds_remaining: max_cheat_seconds - 1,
							cheat_start:             neighbour,
							cheat_end:               current_node.cheat_end,
							moves:                   current_node.moves + 1})
				}

			}
		}

	}

	return cheat_quickest_path
}

func print_board(board map[coord]string, current coord) {
	x_max := 0
	y_max := 0

	for pos, _ := range board {
		x_max = max(x_max, pos.x)
		y_max = max(y_max, pos.y)
	}

	for y := 0; y <= y_max; y++ {
		for x := 0; x <= x_max; x++ {
			pos := coord{x, y}
			if pos == current {
				fmt.Printf("X")
			} else {
				val, exists := board[coord{x, y}]
				if exists {
					fmt.Printf("%s", val)
				}
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

}

func Solve2(input_lines string) int {
	board, start, end := parse_input(input_lines)

	println(len(board), start.x, end.x)

	return 1
}

type node_lookup struct {
	pos                     coord
	cheats_used             int
	cheat_seconds_remaining int
	moves                   int
}

type node struct {
	pos                     coord
	cheats_used             int
	cheat_seconds_remaining int
	moves                   int
	cheat_start             coord
	cheat_end               coord
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

func (pos coord) get_right_down_neighbours() []coord {
	return []coord{
		{pos.x + 1, pos.y},
		{pos.x, pos.y + 1}}
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
