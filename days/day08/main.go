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
				x_diff := coords[i][0] - coords[j][0]
				y_diff := coords[i][1] - coords[j][1]
				antinode_1 := [2]int{coords[i][0] + x_diff, coords[i][1] + y_diff}
				antinode_2 := [2]int{coords[i][0] - 2*x_diff, coords[i][1] - 2*y_diff}

				if coord_on_map(antinode_1, x_max, y_max) {
					_, exists := uniqute_antinodes[antinode_1]
					if !exists {
						uniqute_antinodes[antinode_1] = struct{}{}
					}
				}
				if coord_on_map(antinode_2, x_max, y_max) {
					_, exists := uniqute_antinodes[antinode_2]
					if !exists {
						uniqute_antinodes[antinode_2] = struct{}{}
					}
				}
			}
		}
	}

	return len(uniqute_antinodes)
}

func coord_on_map(coord [2]int, x_max int, y_max int) bool {
	return coord[0] >= 0 && coord[0] < x_max && coord[1] >= 0 && coord[1] < y_max
}

func Solve2(input_lines string) int {
	return 1
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func parse_input(input_lines string) (map[rune][][]int, int, int) {
	antennas := make(map[rune][][]int)

	lines := strings.Split(input_lines, "\n")

	for y, line := range lines {
		for x, antenna := range line {
			if antenna == '.' {
				continue
			}

			coord := []int{x, y}
			coords, exists := antennas[antenna]

			if exists {
				antennas[antenna] = append(coords, coord)
			} else {
				antennas[antenna] = [][]int{coord}
			}
		}
	}
	y_max := len(lines)
	x_max := len(lines[0])
	return antennas, x_max, y_max
}
