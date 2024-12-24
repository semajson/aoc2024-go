package day24

import (
	"math"
	"regexp"
	"strconv"
)

func Solve1(input_lines string) int {
	wire_vals, wire_gates := parse_input(input_lines)

	// Calc difference when sorted
	z_index := 0
	output := 0
	for {
		// Build z wire
		z_wire := "z"
		if z_index < 10 {
			// add padding
			z_wire += "0"
		}
		z_wire += strconv.Itoa(z_index)

		// Exit check
		_, exists := wire_gates[z_wire]
		if !exists {
			break
		}

		output += get_wire_val(z_wire, wire_vals, wire_gates) * int(math.Pow(2, float64(z_index)))
		z_index += 1
	}

	return output
}

func Solve2(input_lines string) int {
	wire_vals, wire_gates := parse_input(input_lines)

	// Calc difference when sorted
	println(len(wire_vals), len(wire_gates))

	return 1
}

func get_wire_val(wire string, wire_vals map[string]int, wire_gates map[string]gate) int {
	val, exists := wire_vals[wire]
	if exists {
		return val
	}

	gate, gate_exists := wire_gates[wire]
	if !gate_exists {
		panic("Unexpected wire")
	}

	left_val := get_wire_val(gate.left, wire_vals, wire_gates)
	right_val := get_wire_val(gate.right, wire_vals, wire_gates)

	wire_val := gate.do_operation(left_val, right_val)

	wire_vals[wire] = wire_val
	return wire_val
}

type gate struct {
	left      string
	right     string
	operation string
}

func (g gate) do_operation(left_val int, right_val int) int {
	switch g.operation {
	case "AND":
		return left_val & right_val
	case "OR":
		return left_val | right_val
	case "XOR":
		return left_val ^ right_val
	default:
		panic("Unexpected operation")
	}
}

func parse_input(input_lines string) (map[string]int, map[string]gate) {

	wire_re, _ := regexp.Compile(`([a-z\d]{3}): (\d)`)

	wire_matches := wire_re.FindAllStringSubmatch(input_lines, -1)

	wire_vals := make(map[string]int)
	for _, match := range wire_matches {
		wire := match[1]
		val, err := strconv.Atoi(match[2])
		if err != nil {
			panic("Error parsing")
		}
		wire_vals[wire] = val
	}

	gate_re, _ := regexp.Compile(`([a-z\d]{3}) ([ANDXOROR]+) ([a-z\d]{3}) -> ([a-z\d]{3})`)

	gate_matches := gate_re.FindAllStringSubmatch(input_lines, -1)

	wire_gates := make(map[string]gate)
	for _, match := range gate_matches {
		left := match[1]
		operation := match[2]
		right := match[3]

		gate := gate{left, right, operation}

		result := match[4]

		wire_gates[result] = gate
	}

	return wire_vals, wire_gates
}
