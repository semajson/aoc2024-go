package day12

import (
	"slices"
	"strings"
)

func Solve1(input_lines string) int {
	pos_map := parse_input(input_lines)

	cost := 0
	seen := map[coord]struct{}{}

	for pos, value := range pos_map {
		_, visited := seen[pos]
		if visited {
			continue
		}

		region := find_region(pos, seen, pos_map, value)
		area := len(region)
		perimeter := calc_peri_1(region, pos_map, value)

		cost += perimeter * area
	}

	return cost
}

func Solve2(input_lines string) int {
	pos_map := parse_input(input_lines)

	cost := 0
	seen := map[coord]struct{}{}

	for pos, value := range pos_map {
		_, visited := seen[pos]
		if visited {
			continue
		}

		region := find_region(pos, seen, pos_map, value)
		area := len(region)
		perimeter := calc_peri_2(region)

		cost += perimeter * area
	}

	return cost
}

func calc_peri_1(region []coord, pos_map map[coord]string, value string) int {
	perimeter := 0
	for _, node := range region {
		for _, neighbour := range node.get_neighbours() {
			neighbour_value, exists := pos_map[neighbour]

			if !exists || value != neighbour_value {
				perimeter += 1
			}
		}
	}
	return perimeter
}

func find_region(pos coord, seen map[coord]struct{}, pos_map map[coord]string, value string) []coord {
	region := []coord{}
	current := []coord{pos}
	for len(current) > 0 {
		// Pop one element from current
		node := current[len(current)-1]
		current = current[:len(current)-1]

		_, visited := seen[node]
		if visited {
			continue
		}

		seen[node] = struct{}{}
		region = append(region, node)

		if !slices.Contains(region, node) {
			panic("Unhitable code 2")
		}

		// Flood fill to get surround region values
		for _, neighbour := range node.get_neighbours() {
			neighbour_value, exists := pos_map[neighbour]

			if !exists || neighbour_value != value {
				continue
			}

			_, visited_neighbour := seen[neighbour]

			if visited_neighbour {
				continue
			}

			current = append(current, neighbour)
		}

	}
	return region
}

func calc_peri_2(region []coord) int {
	// Count the number of corners.
	// This isn't very efficient as it is lots of list lookups

	// Find concave corners
	// X
	// R X
	concave_corners := 0
	for _, pos := range region {

		if !slices.Contains(region, coord{pos.x + 1, pos.y}) && !slices.Contains(region, coord{pos.x, pos.y + 1}) {
			concave_corners += 1
		}
		if !slices.Contains(region, coord{pos.x + 1, pos.y}) && !slices.Contains(region, coord{pos.x, pos.y - 1}) {
			concave_corners += 1
		}
		if !slices.Contains(region, coord{pos.x - 1, pos.y}) && !slices.Contains(region, coord{pos.x, pos.y + 1}) {
			concave_corners += 1
		}
		if !slices.Contains(region, coord{pos.x - 1, pos.y}) && !slices.Contains(region, coord{pos.x, pos.y - 1}) {
			concave_corners += 1
		}
	}

	// Find convex corners
	// R X
	// R R
	convex_corners := 0
	for _, pos := range region {

		if slices.Contains(region, coord{pos.x + 1, pos.y}) && slices.Contains(region, coord{pos.x, pos.y + 1}) && !slices.Contains(region, coord{pos.x + 1, pos.y + 1}) {
			convex_corners += 1
		}
		if slices.Contains(region, coord{pos.x + 1, pos.y}) && slices.Contains(region, coord{pos.x, pos.y - 1}) && !slices.Contains(region, coord{pos.x + 1, pos.y - 1}) {
			convex_corners += 1
		}
		if slices.Contains(region, coord{pos.x - 1, pos.y}) && slices.Contains(region, coord{pos.x, pos.y + 1}) && !slices.Contains(region, coord{pos.x - 1, pos.y + 1}) {
			convex_corners += 1
		}
		if slices.Contains(region, coord{pos.x - 1, pos.y}) && slices.Contains(region, coord{pos.x, pos.y - 1}) && !slices.Contains(region, coord{pos.x - 1, pos.y - 1}) {
			convex_corners += 1
		}
	}

	perimeter := concave_corners + convex_corners
	return perimeter
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
