package day18

import (
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	falling_times, _ := parse_input(input_lines)

	// BFS
	return bfs(falling_times, coord{70, 70}, 1024)
}

func bfs(falling_times map[coord]int, end coord, seconds_to_consider int) int {
	// BFS
	start := coord{0, 0}

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
	return -1
}

func Solve2(input_lines string) string {
	falling_times, max_time := parse_input(input_lines)

	// Binary chop
	end := coord{70, 70}
	lower_time := 0
	for lower_time <= max_time {
		mid := (lower_time + max_time) / 2

		x := bfs(falling_times, end, mid)
		x_plus_1 := bfs(falling_times, end, mid+1)

		if x == -1 {
			max_time = mid - 1
		} else if x_plus_1 != -1 {
			lower_time = mid + 1
		} else {
			// Found match!
			for pos, time := range falling_times {
				if time == (mid + 1) {
					return pos.to_str()
				}
			}
		}
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
func (pos coord) to_str() string {
	return strconv.Itoa(pos.x) + "," + strconv.Itoa(pos.y)
}

func parse_input(input_lines string) (map[coord]int, int) {
	falling_times := make(map[coord]int)

	lines := strings.Split(input_lines, "\n")

	for time, line_raw := range lines {
		line := strings.Split(line_raw, ",")
		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		falling_times[coord{x, y}] = time + 1 // Add one to not zero index
	}
	return falling_times, len(lines)
}
