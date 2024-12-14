package day14

import (
	"fmt"
	"math"
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

	x_max := 101
	y_max := 103

	// Move robots
	x_var_mean := float64(0)
	x_var_var_sum := float64(0)
	x_var_prev := 0

	y_var_mean := float64(0)
	y_var_var_sum := float64(0)
	y_var_prev := 0
	n := 0
	for t := 1; t < 100000; t++ {
		for i, robot := range robots {
			robot.move(1, x_max, y_max)
			robots[i] = robot
		}
		x_var, y_var := get_variance(robots)

		// Calc running variance of the variance!
		// Uses Welford's Algorithm
		n += 1
		x_var_diff := x_var - x_var_prev
		x_var_mean += float64(x_var_diff) / float64(n)
		x_var_var_sum += float64(x_var_diff) * (float64(x_var) - x_var_mean)

		y_var_diff := y_var - y_var_prev
		y_var_mean += float64(y_var_diff) / float64(n)
		y_var_var_sum += float64(y_var_diff) * (float64(y_var) - y_var_mean)

		// Check if it is a big outlier in both x and y
		// by calculating Z score
		x_var_var := x_var_var_sum / float64(n)
		x_var_sd := math.Sqrt(float64(x_var_var))
		x_var_z_score := (float64(x_var) - x_var_mean) / x_var_sd

		y_var_var := y_var_var_sum / float64(n)
		y_var_sd := math.Sqrt(float64(y_var_var))
		y_var_z_score := (float64(y_var) - y_var_mean) / y_var_sd

		// println("Time passed: ", t)
		// println("X var Z score is ", x_var_z)
		// println("Y var Z score is ", y_var_z)

		if math.Abs(x_var_z_score) > 3 && math.Abs(y_var_z_score) > 3 {
			return t
		}

		// draw_board(robots)
		x_var_prev, y_var_prev = x_var, y_var
	}
	return 0
}

func get_variance(robots []robot) (int, int) {
	// Calc mean
	x_sum := 0
	y_sum := 0
	for _, robot := range robots {
		x_sum += robot.x
		y_sum += robot.y
	}
	x_mean := x_sum / len(robots)
	y_mean := y_sum / len(robots)

	// Now calc (population) variance
	x_variance := 0
	y_variance := 0
	for _, robot := range robots {
		x_variance += (robot.x - x_mean) * (robot.x - x_mean)
		y_variance += (robot.y - y_mean) * (robot.y - y_mean)
	}
	x_variance /= len(robots)
	y_variance /= len(robots)

	return x_variance, y_variance
}

func draw_board(robots []robot) {
	x_max := 101
	y_max := 103

	var grid [103][101]int

	// Initialize the array with '.'
	for x := 0; x < x_max; x++ {
		for y := 0; y < y_max; y++ {
			grid[y][x] = 0
		}
	}

	for _, robot := range robots {
		// val := grid[robot.y][robot.x]
		grid[robot.y][robot.x] += 1
	}
	for _, row := range grid {
		for _, num := range row {

			if num == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%d", num)
			}
		}
		fmt.Printf("\n")
	}
	println(" ")
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
