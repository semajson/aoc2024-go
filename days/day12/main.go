package day12

import (
	"slices"
	"strings"
)

func Solve1(input_lines string) int {
	pos_map := parse_input(input_lines)

	// Calc difference when sorted
	cost := 0
	seen := map[coord]struct{}{}

	for pos, value := range pos_map {

		_, visited := seen[pos]
		if visited {
			continue
		}

		// Find all coords in region
		region := []coord{}
		current := []coord{pos}
		for len(current) > 0 {
			new_current := []coord{}
			for _, node := range current {
				_, visited := seen[node]

				if visited {
					continue
				}
				seen[node] = struct{}{}
				region = append(region, node)

				if !slices.Contains(region, node) {
					panic("Unhitable code 2")
				}

				for _, neighbour := range node.get_neighbours() {
					neighbour_value, exists := pos_map[neighbour]

					if !exists || neighbour_value != value {
						continue
					}

					_, visited_neighbour := seen[neighbour]

					if visited_neighbour {
						continue
					}

					new_current = append(new_current, neighbour)

				}
			}
			current = new_current
		}
		area := len(region)

		// Calc perimeta
		perimeta := 0
		for _, node := range region {
			for _, neighbour := range node.get_neighbours() {
				neighbour_value, exists := pos_map[neighbour]

				if !exists || value != neighbour_value {
					perimeta += 1
				}
			}
		}

		// Calc cost
		cost += perimeta * area
	}

	return cost
}

func Solve2(input_lines string) int {
	pos_map := parse_input(input_lines)

	// Calc difference when sorted
	println(pos_map)

	return 1
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
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

func parse_input(input_lines string) map[coord]string {
	pos_map := make(map[coord]string)

	for y, line := range strings.Split(input_lines, "\n") {
		for x, value := range line {
			pos := coord{x, y}
			pos_map[pos] = string(value)
		}
	}
	return pos_map
}
