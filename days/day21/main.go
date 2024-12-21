package day21

import (
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	codes := parse_input(input_lines)
	println(len(codes))

	complexity_sum := 0
	lookup := make(map[lookup_key]int)

	for _, code := range codes {
		code := strings.Join(code, "")
		shortest_len := robot_1_shortest(code, 1, lookup)
		num := numeric(code)
		complexity_sum += shortest_len * num
	}

	return complexity_sum
}

func Solve2(input_lines string) int {
	codes := parse_input(input_lines)
	println(len(codes))

	complexity_sum := 0
	lookup := make(map[lookup_key]int)

	for _, code := range codes {
		code := strings.Join(code, "")
		shortest_len := robot_1_shortest(code, 24, lookup)
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
	nums := []string{}
	for _, char := range code {
		char := string(char)
		if char != "A" {
			nums = append(nums, char)
		}
	}
	num_str := strings.Join(nums, "")

	num, err := strconv.Atoi(num_str)
	if err != nil {
		panic("error")
	}
	return num
}

// func robot_1_get_combos(code []string) []string {
// 	depth_1 := get_code_combos(code, robot_1_mapping)

// 	for _, x := range depth_1 {

// 	}
// }

func robot_1_shortest(code string, depth int, lookup map[lookup_key]int) int {
	board := robot_1_mapping

	valid_map := make(map[coord]struct{})
	for _, pos := range board {
		valid_map[pos] = struct{}{}
	}

	shortest_len := 0
	for i := 0; i < len(code); i++ {

		start := "A"
		if i != 0 {
			start = string(code[i-1])
		}

		end := string(code[i])
		start_coord := robot_1_mapping[start]
		end_coord := robot_1_mapping[end]

		dir_combos := get_dir_combos(start_coord, end_coord, valid_map)

		robot_2_codes := []string{}
		for _, dir_combo := range dir_combos {
			robot_2_code := dir_combo + "A"
			robot_2_codes = append(robot_2_codes, robot_2_code)
		}

		// Recursive call
		potential_paths := []int{}
		for _, robot_2_code := range robot_2_codes {
			potential_path := robot_2_shortest(robot_2_code, depth, lookup)
			potential_paths = append(potential_paths, potential_path)
		}
		// potential_paths := robot_2_codes

		// Pick the shortest one
		robot_2_shortest := -1
		for _, potential_path := range potential_paths {
			if (robot_2_shortest == -1) || potential_path < robot_2_shortest {
				robot_2_shortest = potential_path
			}
		}
		shortest_len = shortest_len + robot_2_shortest
	}
	return shortest_len
}

func robot_2_shortest(code string, depth int, lookup map[lookup_key]int) int {
	if len(lookup) > 0 && len(lookup)%1000 == 0 {
		println("tests")
	}

	key := lookup_key{code, depth}
	val, exists := lookup[key]
	if exists {
		return val
	}

	board := robot_2_mapping

	valid_map := make(map[coord]struct{})
	for _, pos := range board {
		valid_map[pos] = struct{}{}
	}

	shortest_len := 0
	for i := 0; i < len(code); i++ {

		start := "A"
		if i != 0 {
			start = string(code[i-1])
		}
		end := string(code[i])

		start_coord := robot_2_mapping[start]
		end_coord := robot_2_mapping[end]

		// dir_combos := get_dir_combos(start_coord, end_coord, valid_map)
		// dir_combos := []string{get_dir_best_2(start_coord, end_coord)}
		dir_combo := get_dir_best_2(start_coord, end_coord)

		robot_2_code := dir_combo + "A"

		// Potential recursive call
		shortest := 0
		if depth > 0 {
			// potential_path := robot_2_shortest(x, depth-1, lookup)
			// potential_paths = append(potential_paths, potential_path)
			shortest = robot_2_shortest(robot_2_code, depth-1, lookup)
		} else {
			shortest = len(robot_2_code)
		}

		// // Pick the shortest one
		// robot_2_shortest := ""
		// for _, potential_path := range potential_paths {
		// 	if (len(robot_2_shortest) == 0) || (len(potential_paths) < len(robot_2_shortest)) {
		// 		robot_2_shortest = potential_path
		// 	}
		// }

		shortest_len = shortest_len + shortest

	}

	lookup[key] = shortest_len
	return shortest_len
}

func get_dir_best_2(start coord, end coord) string {
	c1_0 := coord{1, 0}
	c2_0 := coord{2, 0}
	c0_1 := coord{0, 1}
	c1_1 := coord{1, 1}
	c2_1 := coord{2, 1}

	if start == end {
		return ""
	} else if start == c2_0 && end == c1_0 {
		return "<"
	} else if start == c2_0 && end == c0_1 {
		return "v<<"
	} else if start == c2_0 && end == c1_1 {
		return "<v" // unsure
	} else if start == c2_0 && end == c2_1 {
		return "v"
	} else if start == c1_0 && end == c2_0 {
		return ">"
	} else if start == c1_0 && end == c0_1 {
		return "v<"
	} else if start == c1_0 && end == c1_1 {
		return "v"
	} else if start == c1_0 && end == c2_1 {
		return "v>"
	} else if start == c0_1 && end == c1_0 {
		return ">^"
	} else if start == c0_1 && end == c2_0 {
		return ">>^"
	} else if start == c0_1 && end == c1_1 {
		return ">"
	} else if start == c0_1 && end == c2_1 {
		return ">>"
	} else if start == c1_1 && end == c1_0 {
		return "^"
	} else if start == c1_1 && end == c2_0 {
		return ">^"
	} else if start == c1_1 && end == c0_1 {
		return "<"
	} else if start == c1_1 && end == c2_1 {
		return ">"
	} else if start == c2_1 && end == c1_0 {
		return "<^"
	} else if start == c2_1 && end == c2_0 {
		return "^"
	} else if start == c2_1 && end == c0_1 {
		return "<<"
	} else if start == c2_1 && end == c1_1 {
		return "<"
	}
	panic("unknown")
}

func get_dir_combos(start coord, end coord, valid_map map[coord]struct{}) []string {
	coord_combos := get_coord_combos(start, end, valid_map)

	if start == end {
		return []string{""}
	}

	dir_combos := []string{}
	for _, coord_combo := range coord_combos {
		dir_combos = append(dir_combos, coords_to_dirs(coord_combo))
	}

	lowest_changes := 99999
	best := []string{}
	for _, dir_combo := range dir_combos {

		if len(dir_combo) == 0 {
			// println("test")
		}

		changes := 0
		curr := dir_combo[0]
		for i := 1; i < len(dir_combo); i++ {
			if curr != dir_combo[i] {
				changes += 1
				curr = dir_combo[i]
			}
		}
		if changes < lowest_changes {
			best = []string{dir_combo}
			lowest_changes = changes
		} else if changes == lowest_changes {
			best = append(best, dir_combo)
		}
	}

	// Need to pick the best
	if len(best) > 1 {
		println("test")
	}

	return best
}

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

func get_coord_combos(start coord, end coord, valid_map map[coord]struct{}) [][]coord {
	combos := [][]coord{}

	if start == end {
		return [][]coord{{start}}
	}

	// Branching
	if (start.x - end.x) > 0 {
		next := coord{start.x - 1, start.y}
		_, valid := valid_map[next]
		if valid {
			combos = append(combos, get_coord_combos(next, end, valid_map)...)
		}
	}
	if (start.x - end.x) < 0 {
		next := coord{start.x + 1, start.y}
		_, valid := valid_map[next]
		if valid {
			combos = append(combos, get_coord_combos(next, end, valid_map)...)
		}
	}
	if (start.y - end.y) > 0 {
		next := coord{start.x, start.y - 1}
		_, valid := valid_map[next]
		if valid {
			combos = append(combos, get_coord_combos(next, end, valid_map)...)
		}
	}
	if (start.y - end.y) < 0 {
		next := coord{start.x, start.y + 1}
		_, valid := valid_map[next]
		if valid {
			combos = append(combos, get_coord_combos(next, end, valid_map)...)
		}
	}

	new_combos := [][]coord{}
	for _, combo := range combos {
		new_combo := []coord{start}
		new_combo = append(new_combo, combo...)
		new_combos = append(new_combos, new_combo)
	}

	return new_combos
}

func coords_to_dirs(coords []coord) string {
	dirs := []string{}

	for i := range coords {
		if i == (len(coords) - 1) {
			continue
		}
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

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func parse_input(input_lines string) [][]string {
	codes := [][]string{}

	for _, line := range strings.Split(input_lines, "\n") {
		code := []string{}
		for _, char := range line {
			code = append(code, string(char))
		}
		codes = append(codes, code)
	}
	return codes
}
