package day08

import (
	"strings"
)

func Solve1(input_lines string) int {
	antennas, x_max, y_max := parse_input(input_lines)

	// Calc difference when sorted
	uniqute_antinodes := make(map[[2]int]struct{})

	for _, coords := range antennas {
		for i := range coords {
			for j := range coords {
				if i <= j {
					continue
				}
				antinodes := get_antinodes_1(coords[i], coords[j], x_max, y_max)

				for _, antinode := range antinodes {
					if coord_on_map(antinode, x_max, y_max) {
						_, exists := uniqute_antinodes[antinode]
						if !exists {
							uniqute_antinodes[antinode] = struct{}{}
						}
					}
				}

			}
		}
	}
	return len(uniqute_antinodes)
}

func get_antinodes_1(a [2]int, b [2]int, x_max int, y_max int) [][2]int {
	antinodes := [][2]int{}
	x_diff := a[0] - b[0]
	y_diff := a[1] - b[1]
	antinode_1 := [2]int{a[0] + x_diff, a[1] + y_diff}
	antinodes = append(antinodes, antinode_1)
	antinode_2 := [2]int{a[0] - 2*x_diff, a[1] - 2*y_diff}
	antinodes = append(antinodes, antinode_2)

	return antinodes
}

func coord_on_map(coord [2]int, x_max int, y_max int) bool {
	return coord[0] >= 0 && coord[0] < x_max && coord[1] >= 0 && coord[1] < y_max
}

func Solve2(input_lines string) int {

	antennas, x_max, y_max := parse_input(input_lines)

	// Calc difference when sorted
	uniqute_antinodes := make(map[[2]int]struct{})

	for _, coords := range antennas {
		for i := range coords {
			for j := range coords {
				if i <= j {
					continue
				}
				antinodes := get_antinodes_2(coords[i], coords[j], x_max, y_max)

				for _, antinode := range antinodes {
					if coord_on_map(antinode, x_max, y_max) {
						_, exists := uniqute_antinodes[antinode]
						if !exists {
							uniqute_antinodes[antinode] = struct{}{}
						}
					}
				}

			}
		}
	}
	return len(uniqute_antinodes)
}

func get_antinodes_2(a [2]int, b [2]int, x_max int, y_max int) [][2]int {
	antinodes := [][2]int{a}
	x_diff := a[0] - b[0]
	y_diff := a[1] - b[1]

	// Increase
	increase_node := [2]int{a[0] + x_diff, a[1] + y_diff}
	for coord_on_map(increase_node, x_max, y_max) {
		antinodes = append(antinodes, increase_node)
		increase_node[0] = increase_node[0] + x_diff
		increase_node[1] = increase_node[1] + y_diff
	}

	// Decrease
	decrease_node := [2]int{a[0] - x_diff, a[1] - y_diff}
	for coord_on_map(decrease_node, x_max, y_max) {
		antinodes = append(antinodes, decrease_node)
		decrease_node[0] = decrease_node[0] - x_diff
		decrease_node[1] = decrease_node[1] - y_diff
	}

	return antinodes
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func parse_input(input_lines string) (map[rune][][2]int, int, int) {
	antennas := make(map[rune][][2]int)

	lines := strings.Split(input_lines, "\n")

	for y, line := range lines {
		for x, antenna := range line {
			if antenna == '.' {
				continue
			}

			coord := [2]int{x, y}
			coords, exists := antennas[antenna]

			if exists {
				antennas[antenna] = append(coords, coord)
			} else {
				antennas[antenna] = [][2]int{coord}
			}
		}
	}
	y_max := len(lines)
	x_max := len(lines[0])
	return antennas, x_max, y_max
}
