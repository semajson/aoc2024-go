package day14

import (
	"regexp"
	"strconv"
)

type coord struct {
	x int
	y int
}

func Solve1(input_lines string) int {
	robots := parse_input(input_lines)

	seconds := 100
	x_max := 101
	y_max := 103

	// Move robots
	for i, robot := range robots {
		robot.move(seconds, x_max, y_max)
		robots[i] = robot
	}

	// Count robots
	robot_count := make(map[coord]int)
	for _, robot := range robots {
		pos := coord{robot.x, robot.y}
		count, exists := robot_count[pos]

		if exists {
			robot_count[pos] = count + 1
		} else {
			robot_count[pos] = 1
		}
	}

	top_left_count := 0
	top_right_count := 0
	bottom_left_count := 0
	bottom_right_count := 0
	for y := 0; y < y_max; y++ {
		for x := 0; x < x_max; x++ {
			pos := coord{x, y}
			count, exists := robot_count[pos]
			if !exists {
				continue
			}

			left := x < (x_max / 2)
			right := x > (x_max / 2)
			top := y < (y_max / 2)
			bottom := y > (y_max / 2)

			if top && left {
				top_left_count += count
			} else if top && right {
				top_right_count += count
			} else if bottom && left {
				bottom_left_count += count
			} else if bottom && right {
				bottom_right_count += count
			}
		}
	}

	return top_left_count * top_right_count * bottom_left_count * bottom_right_count
}

func Solve2(input_lines string) int {
	robots := parse_input(input_lines)

	println(robots)

	return 1
}

type robot struct {
	x  int
	y  int
	dx int
	dy int
}

func (r *robot) move(seconds int, x_max int, y_max int) {
	x_moved := seconds * r.dx
	y_moved := seconds * r.dy

	// Deal with wrap around
	r.x += x_moved
	r.x %= x_max
	r.y += y_moved
	r.y %= y_max

	if r.x < 0 {
		r.x += x_max
	}
	if r.y < 0 {
		r.y += y_max
	}
}

func parse_input(input_lines string) []robot {
	robots := []robot{}

	re, _ := regexp.Compile(`p\=(-?\d+),(-?\d+) v\=(-?\d+),(-?\d+)`)

	matches := re.FindAllStringSubmatch(input_lines, -1)

	for _, match := range matches {
		x, err_1 := strconv.Atoi(match[1])
		y, err_2 := strconv.Atoi(match[2])
		dx, err_3 := strconv.Atoi(match[3])
		dy, err_4 := strconv.Atoi(match[4])
		if err_1 != nil || err_2 != nil || err_3 != nil || err_4 != nil {
			panic("error passing input")
		}
		robots = append(robots, robot{x, y, dx, dy})
	}
	return robots
}
