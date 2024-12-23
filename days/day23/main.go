package day23

import (
	"slices"
	"sort"
	"strings"
)

func Solve1(input_lines string) int {
	edges := parse_input(input_lines)

	// Build the graph
	reachable, all_nodes := build_graph(edges)

	multiplayer_three_games := make(map[[3]string]struct{})
	println(len(reachable))
	for node, _ := range all_nodes {
		for _, neighbour := range reachable[node] {
			neighbours := reachable[node]
			neighbour_reachable := reachable[neighbour]

			// Find intersection between node and neighbour - that gives list of 3
			for _, neighbour_neighbour := range neighbour_reachable {
				if slices.Contains(neighbours, neighbour_neighbour) {
					found_game := [3]string{node, neighbour, neighbour_neighbour}
					slice := found_game[:]
					sort.Strings(slice)
					copy(slice, found_game[:])

					_, exists := multiplayer_three_games[found_game]
					if !exists {
						multiplayer_three_games[found_game] = struct{}{}
					}
				}
			}

		}
	}

	games_with_t := 0
	for game, _ := range multiplayer_three_games {
		for _, node := range game {
			if strings.HasPrefix(node, "t") {
				games_with_t += 1
				break
			}
		}
	}
	return games_with_t
}

func build_graph(edges [][2]string) (map[string][]string, map[string]struct{}) {
	reachable := make(map[string][]string)
	all_nodes := make(map[string]struct{})
	for _, edge := range edges {
		node_1 := edge[0]
		node_2 := edge[1]

		reachable[node_1] = append(reachable[node_1], node_2)
		reachable[node_2] = append(reachable[node_2], node_1)

		all_nodes[node_1] = struct{}{}
		all_nodes[node_2] = struct{}{}
	}
	return reachable, all_nodes
}

func Solve2(input_lines string) string {
	edges := parse_input(input_lines)
	println(len(edges))

	// Build the graph
	reachable, all_nodes := build_graph(edges)

	best := []string{}
	seen := map[string]struct{}{}
	for node := range all_nodes {
		neighbours := reachable[node]

		// Skip nodes already processed
		neighbours_not_seen := []string{}
		for _, neighbour := range neighbours {
			_, visited := seen[neighbour]
			if !visited {
				neighbours_not_seen = append(neighbours_not_seen, neighbour)
			}
		}

		// Generate subsets + look for connected users
		for _, subset := range get_subsets(neighbours_not_seen) {
			if (len(subset) + 1) > len(best) {
				// Contender for "best" game
				if is_connected(subset, reachable) {
					best = append(subset, node)
				}
			}
		}

		seen[node] = struct{}{}
	}

	sort.Strings(best)

	return strings.Join(best, ",")
}

func get_subsets(nodes []string) [][]string {
	combinations := [][]string{}

	// Clever bitwise stuff to build combinations
	// Implemented with help from:
	// github.com/mxschmitt/golang-combinations combinations.All()
	n := len(nodes) * len(nodes)
	for i := 0; i < n; i++ {
		subset := []string{}
		for j, node := range nodes {
			if (i>>j)&1 == 0 {
				subset = append(subset, node)
			}
		}
		combinations = append(combinations, subset)
	}
	return combinations
}

func is_connected(nodes []string, reachable map[string][]string) bool {
	lookup := make(map[string]bool)
	for _, elem := range nodes {
		lookup[elem] = true
	}

	for _, node := range nodes {
		neighbours := reachable[node]

		for _, other_node := range nodes {
			if node == other_node {
				continue
			}

			if !slices.Contains(neighbours, other_node) {
				return false
			}
		}
	}
	return true
}

func parse_input(input_lines string) [][2]string {
	edges := [][2]string{}

	for _, line := range strings.Split(input_lines, "\n") {
		nodes := strings.Split(line, "-")
		edges = append(edges, [2]string{nodes[0], nodes[1]})
	}

	return edges
}
