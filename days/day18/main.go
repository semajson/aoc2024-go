package day18

import (
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	falling_times := parse_input(input_lines)

	// BFS
	start := coord{0, 0}

	// end := coord{6, 6}
	end := coord{70, 70}
	seconds_to_consider := 1024

	curr := []coord{start}
	seen := make(map[coord]struct{})
	path_length := 0

	for len(curr) > 0 {
		// Js9 - for bs first, what is easier?
		// Do all current at a time, or one at a time?
		next := []coord{}

		for _, node := range curr {
			// Skip if needed
			_, visited := seen[node]
			if visited {
				continue
			}
			seen[node] = struct{}{}
			if node.x < 0 || node.x > end.x || node.y < 0 || node.y > end.y {
				continue
			}

			// Exit check
			if node == end {
				return path_length
			}

			// Branch
			for _, neighbour := range node.get_neighbours() {
				time, will_fall := falling_times[neighbour]

				if !will_fall || (time > seconds_to_consider) {
					next = append(next, neighbour)
				}
			}

		}

		curr = next
		path_length += 1
	}

	return 1
}

func Solve2(input_lines string) int {
	falling_times := parse_input(input_lines)

	// Calc difference when sorted
	println(len(falling_times))

	return 1
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

func parse_input(input_lines string) map[coord]int {
	falling_times := make(map[coord]int)

	for time, line_raw := range strings.Split(input_lines, "\n") {
		line := strings.Split(line_raw, ",")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		falling_times[coord{x, y}] = time + 1 // Add one to not zero index
	}
	return falling_times
}
