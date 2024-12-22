package day21

import (
	"strconv"
	"strings"
)

// Robot 1
var robot_1_mapping = map[string]coord{
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

// Robot 1
var robot_2_mapping = map[string]coord{
	"^": {1, 0},
	"A": {2, 0},
	"<": {0, 1},
	"v": {1, 1},
	">": {2, 1},
}

func Solve1(input_lines string) int {
	codes := parse_input(input_lines)

	complexity_sum := 0
	r2_lookup := make(map[lookup_key]int)

	for _, code := range codes {
		shortest_len := r1_shortest(code, 1, r2_lookup)
		num := numeric(code)
		complexity_sum += shortest_len * num
	}

	return complexity_sum
}

func Solve2(input_lines string) int {
	codes := parse_input(input_lines)

	complexity_sum := 0
	r2_lookup := make(map[lookup_key]int)

	for _, code := range codes {
		shortest_len := r1_shortest(code, 24, r2_lookup)
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

func r1_shortest(code string, depth int, lookup map[lookup_key]int) int {
	// Build useful lookup structs
	valid_map_r1 := make(map[coord]struct{})
	for _, pos := range robot_1_mapping {
		valid_map_r1[pos] = struct{}{}
	}
	valid_map_r2 := make(map[coord]struct{})
	for _, pos := range robot_2_mapping {
		valid_map_r2[pos] = struct{}{}
	}

	total_shortest_len := 0
	for i := 0; i < len(code); i++ {
		// Robots start on A
		start := "A"
		if i != 0 {
			start = string(code[i-1])
		}
		end := string(code[i])

		start_coord := robot_1_mapping[start]
		end_coord := robot_1_mapping[end]

		dir_paths := get_dir_paths(start_coord, end_coord, valid_map_r1)

		robot_2_codes := []string{}
		for _, dir_path := range dir_paths {
			robot_2_code := dir_path + "A"
			robot_2_codes = append(robot_2_codes, robot_2_code)
		}

		// Recursive call
		options := []int{}
		for _, robot_2_code := range robot_2_codes {
			option := r2_shortest(robot_2_code, depth, lookup, valid_map_r2)
			options = append(options, option)
		}

		// Pick the shortest one
		shortest_len := -1
		for _, option := range options {
			if (shortest_len == -1) || option < shortest_len {
				shortest_len = option
			}
		}
		total_shortest_len = total_shortest_len + shortest_len
	}
	return total_shortest_len
}

func r2_shortest(code string, depth int, lookup map[lookup_key]int, valid_map_r2 map[coord]struct{}) int {
	key := lookup_key{code, depth}
	val, exists := lookup[key]
	if exists {
		return val
	}

	total_shortest_len := 0
	for i := 0; i < len(code); i++ {
		start := "A"
		if i != 0 {
			start = string(code[i-1])
		}
		end := string(code[i])

		start_coord := robot_2_mapping[start]
		end_coord := robot_2_mapping[end]

		dir_paths := get_dir_paths(start_coord, end_coord, valid_map_r2)

		robot_2_codes := []string{}
		for _, dir_path := range dir_paths {
			robot_2_codes = append(robot_2_codes, dir_path+"A")
		}

		// Potential recursive call
		options := []int{}
		for _, x := range robot_2_codes {
			if depth > 0 {
				option := r2_shortest(x, depth-1, lookup, valid_map_r2)
				options = append(options, option)
			} else {
				options = append(options, len(x))
			}
		}

		// Pick the shortest option
		shortest_len := -1
		for _, option := range options {
			if (shortest_len == -1) || option < shortest_len {
				shortest_len = option
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
