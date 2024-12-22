package day21

import (
	"strconv"
	"strings"
)

// Numeric keypad
var numeric_keypad_mapping = map[string]coord{
	"7": {0, 0},
	"8": {1, 0},
	"9": {2, 0},
	"4": {0, 1},
	"5": {1, 1},
	"6": {2, 1},
	"1": {0, 2},
	"2": {1, 2},
	"3": {2, 2},
	"0": {1, 3},
	"A": {2, 3},
}
var numeric_keypad_valid = map[coord]struct{}{
	{0, 0}: {},
	{1, 0}: {},
	{2, 0}: {},
	{0, 1}: {},
	{1, 1}: {},
	{2, 1}: {},
	{0, 2}: {},
	{1, 2}: {},
	{2, 2}: {},
	{1, 3}: {},
	{2, 3}: {},
}

// Directional keypad
var directional_keypad_mapping = map[string]coord{
	"^": {1, 0},
	"A": {2, 0},
	"<": {0, 1},
	"v": {1, 1},
	">": {2, 1},
}
var directional_keypad_valid = map[coord]struct{}{
	{1, 0}: {},
	{2, 0}: {},
	{0, 1}: {},
	{1, 1}: {},
	{2, 1}: {},
}

func Solve1(input_lines string) int {
	codes := parse_input(input_lines)

	return calc_complexity_sum(codes, 2)
}

func Solve2(input_lines string) int {
	codes := parse_input(input_lines)

	return calc_complexity_sum(codes, 25)
}

func calc_complexity_sum(codes []string, depth int) int {
	complexity_sum := 0
	lookup := make(map[lookup_key]int)

	for _, code := range codes {
		shortest_len := get_min_len(code, depth, lookup, true)
		num := numeric(code)
		complexity_sum += shortest_len * num
	}

	return complexity_sum
}

type lookup_key struct {
	code  string
	depth int
}

func numeric(code string) int {
	num_str := strings.TrimSuffix(code, "A")
	num, err := strconv.Atoi(num_str)
	if err != nil {
		panic("error")
	}
	return num
}

func get_min_len(code string, depth int, lookup map[lookup_key]int, num_pad bool) int {
	// Check lookup
	key := lookup_key{code, depth}
	val, exists := lookup[key]
	if exists {
		return val
	}

	// Get the correct button mapping depending on the
	// keypad type
	var valid_map map[coord]struct{}
	var coord_map map[string]coord
	if num_pad {
		coord_map = numeric_keypad_mapping
		valid_map = numeric_keypad_valid
	} else {
		coord_map = directional_keypad_mapping
		valid_map = directional_keypad_valid
	}

	// Always start from the A key
	code = "A" + code

	// Take 2 adjacent parts of the code at a time and workout the
	// min instruction length needed for them
	total_shortest_len := 0
	for i := 0; i < len(code)-1; i++ {
		start := string(code[i])
		end := string(code[i+1])

		start_coord := coord_map[start]
		end_coord := coord_map[end]

		dir_paths := get_dir_paths(start_coord, end_coord, valid_map)
		robot_codes := []string{}
		for _, dir_path := range dir_paths {
			robot_codes = append(robot_codes, dir_path+"A")
		}

		// Potential recursive call
		path_lens := []int{}
		for _, robot_code := range robot_codes {
			if depth > 0 {
				path_len := get_min_len(robot_code, depth-1, lookup, false)
				path_lens = append(path_lens, path_len)
			} else {
				path_lens = append(path_lens, len(robot_code))
			}
		}

		// Pick the shortest option
		shortest_len := -1
		for _, path_len := range path_lens {
			if (shortest_len == -1) || path_len < shortest_len {
				shortest_len = path_len
			}
		}

		total_shortest_len = total_shortest_len + shortest_len
	}

	lookup[key] = total_shortest_len
	return total_shortest_len
}

func get_dir_paths(start coord, end coord, valid_map map[coord]struct{}) []string {
	if start == end {
		return []string{""}
	}

	// Find all dir paths (with min number of moves)
	coord_paths := get_coord_paths(start, end, valid_map)
	dir_paths := []string{}
	for _, coord_path := range coord_paths {
		dir_paths = append(dir_paths, coords_to_dirs(coord_path))
	}

	return dir_paths
}

func get_coord_paths(start coord, end coord, valid_map map[coord]struct{}) [][]coord {
	paths := [][]coord{}

	if start == end {
		return [][]coord{{start}}
	}

	// Branching
	if (start.x - end.x) > 0 {
		next := coord{start.x - 1, start.y}
		_, valid := valid_map[next]
		if valid {
			paths = append(paths, get_coord_paths(next, end, valid_map)...)
		}
	}
	if (start.x - end.x) < 0 {
		next := coord{start.x + 1, start.y}
		_, valid := valid_map[next]
		if valid {
			paths = append(paths, get_coord_paths(next, end, valid_map)...)
		}
	}
	if (start.y - end.y) > 0 {
		next := coord{start.x, start.y - 1}
		_, valid := valid_map[next]
		if valid {
			paths = append(paths, get_coord_paths(next, end, valid_map)...)
		}
	}
	if (start.y - end.y) < 0 {
		next := coord{start.x, start.y + 1}
		_, valid := valid_map[next]
		if valid {
			paths = append(paths, get_coord_paths(next, end, valid_map)...)
		}
	}

	new_paths := [][]coord{}
	for _, path := range paths {
		new_path := []coord{start}
		new_path = append(new_path, path...)
		new_paths = append(new_paths, new_path)
	}

	return new_paths
}

func coords_to_dirs(coords []coord) string {
	dirs := []string{}

	for i := 0; i < len(coords)-1; i++ {
		if (coords[i+1].x - coords[i].x) == 1 {
			dirs = append(dirs, ">")
		} else if (coords[i+1].x - coords[i].x) == -1 {
			dirs = append(dirs, "<")
		} else if (coords[i+1].y - coords[i].y) == -1 {
			dirs = append(dirs, "^")
		} else if (coords[i+1].y - coords[i].y) == 1 {
			dirs = append(dirs, "v")
		}
	}
	return strings.Join(dirs, "")
}

type coord struct {
	x int
	y int
}

func parse_input(input_lines string) []string {
	return strings.Split(input_lines, "\n")
}
