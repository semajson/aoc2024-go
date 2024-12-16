package day16

import (
	"container/heap"
	"strings"
)

func Solve1(input_lines string) int {
	board, start_pos, end_pos := parse_input(input_lines)

	// Dijkstra solve
	start := node{pos: start_pos, dir: 1, score: 0}

	return Dijkstra(board, start, end_pos)
}

func Dijkstra(board map[coord]struct{}, start node, end_pos coord) int {
	queue := &NodeHeap{}
	heap.Init(queue)
	heap.Push(queue, start)
	seen := make(map[node_key]struct{})

	for queue.Len() > 0 {
		curr_node := heap.Pop(queue).(node)
		lookup_key := node_key{curr_node.pos, curr_node.dir}

		_, visited := seen[lookup_key]
		if visited {
			continue
		}
		seen[lookup_key] = struct{}{}

		// Exit check
		if curr_node.pos == end_pos {
			return curr_node.score
		}

		// Straight branch
		dx := DIRECTIONS[curr_node.dir][0]
		dy := DIRECTIONS[curr_node.dir][1]
		new_pos := coord{curr_node.pos.x + dx, curr_node.pos.y + dy}

		_, exists := board[new_pos]
		if exists {
			straight_branch := node{pos: new_pos,
				dir:   curr_node.dir,
				score: curr_node.score + 1}
			heap.Push(queue, straight_branch)
		}

		// Clockwise branch
		clockwise_branch := node{pos: curr_node.pos,
			dir:   (curr_node.dir + 1) % len(DIRECTIONS),
			score: curr_node.score + 1000}
		heap.Push(queue, clockwise_branch)

		// Anti clockwise branch
		anti_clockwise_branch := node{pos: curr_node.pos,
			dir:   (curr_node.dir - 1 + len(DIRECTIONS)) % len(DIRECTIONS),
			score: curr_node.score + 1000}
		heap.Push(queue, anti_clockwise_branch)
	}
	panic("Didn't solve maze")
}

func Solve2(input_lines string) int {
	board, start, end := parse_input(input_lines)

	// Calc difference when sorted
	println(len(board), start.x, end.x)

	return 1
}

type coord struct {
	x int
	y int
}

var DIRECTIONS = [][]int{
	{0, -1}, // N
	{1, 0},  // E
	{0, 1},  // S
	{-1, 0}} // W

type node struct {
	pos   coord
	dir   int
	score int
}

type node_key struct {
	pos coord
	dir int
}

// Define a type that implements heap.Interface
type NodeHeap []node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].score < h[j].score } // Min-heap (smallest first)
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(node))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func parse_input(input_lines string) (map[coord]struct{}, coord, coord) {
	board := make(map[coord]struct{})
	var start coord
	var end coord

	for y, line := range strings.Split(input_lines, "\n") {
		for x, val := range line {
			switch val {
			case '.':
				board[coord{x, y}] = struct{}{}
			case 'S':
				board[coord{x, y}] = struct{}{}
				start = coord{x, y}
			case 'E':
				board[coord{x, y}] = struct{}{}
				end = coord{x, y}
			}
		}
	}
	return board, start, end
}
