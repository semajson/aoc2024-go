package day10

import (
	"strconv"
	"strings"

	"slices"
)

func Solve1(input_lines string) int {
	heights := parse_input(input_lines)

	// Find zeros
	zeros := []coord{}
	for coord, height := range heights {
		if height == 0 {
			zeros = append(zeros, coord)
		}
	}

	// Find score sum
	score_sum := 0
	lookup := make(map[coord][]coord)
	for _, zero_pos := range zeros {
		// DFS with lookup
		score_sum += len(dfs_distinct_9s(zero_pos, heights, lookup))
	}

	return score_sum
}

func dfs_distinct_9s(pos coord, heights map[coord]int, lookup map[coord][]coord) []coord {
	// Cheeky lookup
	lookup_score, exists := lookup[pos]
	if exists {
		return lookup_score
	}

	// Check if at end
	curr_height, exists := heights[pos]
	if !exists {
		panic("Un-hittable code")
	}
	if curr_height == 9 {
		return []coord{pos}
	}

	// Get options
	neighbours := pos.get_neighbours()
	reachable_9s := []coord{}
	for _, neighbour := range neighbours {
		neighbour_height, exists := heights[neighbour]
		if !exists {
			// Not on board
			continue
		}
		if (neighbour_height - curr_height) != 1 {
			// Too steep / goes down
			continue
		}

		for _, reachable_9 := range dfs_distinct_9s(neighbour, heights, lookup) {
			if !slices.Contains(reachable_9s, reachable_9) {
				reachable_9s = append(reachable_9s, reachable_9)
			}
		}
	}

	// Add to lookup
	lookup[pos] = reachable_9s
	return reachable_9s
}

func Solve2(input_lines string) int {
	heights := parse_input(input_lines)

	// Find zeros
	zeros := []coord{}
	for coord, height := range heights {
		if height == 0 {
			zeros = append(zeros, coord)
		}
	}

	// Find score sum
	score_sum := 0
	lookup := make(map[coord]int)
	for _, zero_pos := range zeros {
		// DFS with lookup
		score_sum += dfs_paths(zero_pos, heights, lookup)
	}

	// debug_print(lookup)

	return score_sum
}

func dfs_paths(pos coord, heights map[coord]int, lookup map[coord]int) int {
	// Cheeky lookup
	lookup_score, exists := lookup[pos]
	if exists {
		return lookup_score
	}

	// Check if at end
	curr_height, exists := heights[pos]
	if !exists {
		panic("Un-hittable code")
	}
	if curr_height == 9 {
		return 1
	}

	// Get options
	neighbours := pos.get_neighbours()
	score := 0
	for _, neighbour := range neighbours {
		neighbour_height, exists := heights[neighbour]
		if !exists {
			// Not on board
			continue
		}
		if (neighbour_height - curr_height) != 1 {
			// Too steep / goes down
			continue
		}

		score += dfs_paths(neighbour, heights, lookup)
	}

	// Add to lookup
	lookup[pos] = score
	return score
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

func debug_print_board(test map[coord]int) {
	new_test := [10][10]int{}

	for pos, value := range test {
		new_test[pos.y][pos.x] = value
	}

	println("hello")
}

func parse_input(input_lines string) map[coord]int {
	heights := make(map[coord]int)

	for y, line := range strings.Split(input_lines, "\n") {
		for x := range line {
			pos := coord{x, y}
			height, err := strconv.Atoi(string(line[x]))

			if err == nil {
				heights[pos] = height
			}
		}

	}
	return heights
}
