package day13

import (
	"container/heap"
	"regexp"
	"strconv"
)

func Solve1(input_lines string) int {
	machines := parse_input(input_lines)

	output := 0
	for _, machine := range machines {
		output += get_min_tokens_fast(machine)
	}

	return output
}

func get_min_token_slow(machine machine) int {
	// Could use dikstra algo to find the minimum

	seen := make(map[node]struct{})

	queue := &NodeHeap{}
	heap.Init(queue)
	heap.Push(queue, node{0, 0})

	for queue.Len() > 0 {
		curr_node := heap.Pop(queue).(node)

		_, visited := seen[curr_node]

		if visited {
			continue
		}
		seen[curr_node] = struct{}{}

		// Check for exit
		x := curr_node.a_count*machine.a_dx + curr_node.b_count*machine.b_dx
		y := curr_node.a_count*machine.a_dy + curr_node.b_count*machine.b_dy
		if x == machine.prize_x && y == machine.prize_y {
			// Found the end!
			return curr_node.tokens()
		} else if x > machine.prize_x || y > machine.prize_y {
			continue
		}

		// B branch
		b_branch := node{curr_node.a_count, curr_node.b_count + 1}
		_, visited_b := seen[b_branch]
		if !visited_b {
			heap.Push(queue, b_branch)
		}

		// A branch
		a_branch := node{curr_node.a_count + 1, curr_node.b_count}
		_, visited_a := seen[a_branch]
		if !visited_a {
			heap.Push(queue, a_branch)
		}

	}
	return 0
}

func get_min_tokens_fast(machine machine) int {
	b_top := (machine.a_dx*machine.prize_y - machine.a_dy*machine.prize_x)
	b_bottom := (machine.a_dx*machine.b_dy - machine.a_dy*machine.b_dx)
	if b_top%b_bottom != 0 {
		return 0
	}
	b_count := b_top / b_bottom
	if b_count < 0 {
		return 0
	}

	a_top := (machine.b_dx*machine.prize_y - machine.b_dy*machine.prize_x)
	a_bottom := (machine.a_dy*machine.b_dx - machine.a_dx*machine.b_dy)
	if a_top%a_bottom != 0 {
		return 0
	}
	a_count := a_top / a_bottom
	if b_count < 0 {
		return 0
	}

	return a_count*3 + b_count
}

func Solve2(input_lines string) int {
	machines := parse_input(input_lines)

	output := 0
	for _, machine := range machines {
		machine.prize_x += 10000000000000
		machine.prize_y += 10000000000000
		output += get_min_tokens_fast(machine)
	}

	return output
}

type node struct {
	a_count int
	b_count int
}

func (x node) tokens() int {
	return x.a_count*3 + x.b_count
}

// Define a type that implements heap.Interface
type NodeHeap []node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].tokens() < h[j].tokens() } // Min-heap (smallest first)
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

type machine struct {
	a_dx    int
	a_dy    int
	b_dx    int
	b_dy    int
	prize_x int
	prize_y int
}

func parse_input(input_lines string) []machine {
	machines := []machine{}

	re, _ := regexp.Compile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X\=(\d+), Y\=(\d+)`)

	matches := re.FindAllStringSubmatch(input_lines, -1)

	for _, match := range matches {
		a_dx, err_1 := strconv.Atoi(match[1])
		a_dy, err_2 := strconv.Atoi(match[2])
		b_dx, err_3 := strconv.Atoi(match[3])
		b_dy, err_4 := strconv.Atoi(match[4])
		prize_x, err_5 := strconv.Atoi(match[5])
		prize_y, err_6 := strconv.Atoi(match[6])

		if err_1 != nil || err_2 != nil || err_3 != nil || err_4 != nil || err_5 != nil || err_6 != nil {
			panic("error passing input")
		}
		machines = append(machines, machine{a_dx, a_dy, b_dx, b_dy, prize_x, prize_y})
	}

	return machines
}
