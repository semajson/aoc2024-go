package day24

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Solve1(input_lines string) int {
	wire_vals, wire_gates := parse_input(input_lines)

	// Calc difference when sorted

	return get_z_output(wire_vals, wire_gates)
}

var max_swaps int = 5

func Solve2(input_lines string) string {
	wire_vals, wire_gates := parse_input(input_lines)
	println(len(wire_gates), len(wire_vals))

	swaps_list := [][][2]string{[][2]string{}}
	// potential_swaps = get_possible_swaps(wire_gates, potential_swaps)
	// for _, swap := range get_possible_swaps(wire_gates, [][2]string{}) {
	// 	swaps_list = append(swaps_list, [][2]string{swap})
	// }

	for i := 0; i < 44; i++ {
		println("Checking i", i)
		println("Swap list len is", len(swaps_list))
		for _, x := range swaps_list {
			for _, y := range x {
				fmt.Printf("%s,", y)
			}
			fmt.Printf("\n")
		}
		println("")

		new_swaps_list := [][][2]string{}
		seen_here := map[string]struct{}{}
		for _, swaps := range swaps_list {
			// if slices.Contains(seen_here, swaps) {
			// 	continue
			// }
			println("in here")

			// recursive
			// new_swaps_list = append(new_swaps_list, check_and_maybe_swap(wire_gates, swaps, i, 1)...)
			key := swaps_to_string(swaps)
			_, seen := seen_here[key]
			if seen {
				continue
			}
			seen_here[key] = struct{}{}

			if valid_up_plus_swap(i, wire_gates, swaps) {
				new_swaps_list = append(new_swaps_list, swaps)
			} else {
				if len(swaps) >= max_swaps {
					// Too many swaps already, ignore

				} else {
					// Try another swap
					for _, new_swap := range get_possible_swaps(wire_gates, swaps) {
						new_swaps := append(swaps, new_swap)

						if valid_up_plus_swap(i, wire_gates, new_swaps) {
							new_swaps_list = append(new_swaps_list, new_swaps)
						}
						// new_swaps_list = append(new_swaps_list, check_and_maybe_swap(wire_gates, new_swaps, z_test, depth-1)...)
					}
				}
			}

		}
		swaps_list = new_swaps_list
	}

	// answer := swaps_list[0]
	// gates_swapped := []string{}
	// for _, swap := range answer {
	// 	gates_swapped = append(gates_swapped, swap[0])
	// 	gates_swapped = append(gates_swapped, swap[1])
	// }

	// sort.Strings(gates_swapped)

	return swaps_to_string(swaps_list[0])
}

func swaps_to_string(swaps [][2]string) string {
	gates_swapped := []string{}
	for _, swap := range swaps {
		gates_swapped = append(gates_swapped, swap[0])
		gates_swapped = append(gates_swapped, swap[1])
	}

	sort.Strings(gates_swapped)
	return strings.Join(gates_swapped, ",")
}

func check_and_maybe_swap(wire_gates map[string]gate, current_swaps [][2]string, z_test int, depth int) [][][2]string {
	new_swaps_list := [][][2]string{}

	if valid_up_plus_swap(z_test, wire_gates, current_swaps) {
		new_swaps_list = append(new_swaps_list, current_swaps)
	} else {
		if len(current_swaps) >= max_swaps || depth == 0 {
			// Too many swaps already, ignore

		} else {
			// Try another swap
			for _, new_swap := range get_possible_swaps(wire_gates, current_swaps) {
				new_swaps := append(current_swaps, new_swap)
				new_swaps_list = append(new_swaps_list, check_and_maybe_swap(wire_gates, new_swaps, z_test, depth-1)...)
			}
		}
	}
	return new_swaps_list
}

func get_z_output(wire_vals map[string]int, wire_gates map[string]gate) int {
	// Calc difference when sorted
	z_index := 0
	output := 0
	for {
		// Build z wire
		z_wire := "z"
		z_wire += get_padded_num(z_index)

		// Exit check
		_, exists := wire_gates[z_wire]
		if !exists {
			break
		}

		loop_detection := make(map[string]struct{})
		output += get_wire_val(z_wire, wire_vals, wire_gates, loop_detection) * int(math.Pow(2, float64(z_index)))
		z_index += 1
	}

	return output
}

func get_padded_num(num int) string {
	output := ""
	if num < 10 {
		// add padding
		output += "0"
	}
	output += strconv.Itoa(num)

	return output
}

func valid_up_plus_swap(z_end int, wire_gates map[string]gate, swaps [][2]string) bool {
	// println("in valid_up_plus_swap")
	// Apply swaps
	for _, swaps := range swaps {
		a := swaps[0]
		b := swaps[1]
		a_gate := wire_gates[a]
		b_gate := wire_gates[b]
		wire_gates[a] = b_gate
		wire_gates[b] = a_gate
	}

	if len(swaps) > 1 && len(swaps[0]) > 0 {
		if swaps[0][0] == "z00" && swaps[0][1] == "z05" {
			if swaps[1][0] == "z01" && swaps[1][1] == "z02" {

				// println("in here")
			}
		}
	}
	valid := valid_up_to_test(z_end, wire_gates)

	// Undo swaps
	for _, swaps := range swaps {
		a := swaps[0]
		b := swaps[1]
		a_gate := wire_gates[a]
		b_gate := wire_gates[b]
		wire_gates[a] = b_gate
		wire_gates[b] = a_gate
	}

	return valid
}

func valid_up_to(z_end int, wire_gates map[string]gate) bool {
	// println("in valid_up_to")

	x := 613928449
	y := 104366443
	wire_vals := map[string]int{}
	Set_x_y(wire_vals, x, y)

	z_out := get_z_output(wire_vals, wire_gates)

	z_expected := x + y
	for i := 0; i <= z_end; i++ {
		power := int(math.Pow(2, float64(i)))
		if (power & z_out) != (power & z_expected) {
			return false
		}
	}
	return true
	// if z_out != (x + y) {
	// 	return false
	// }

	// for x := 0; x < z_end; x++ {
	// 	for y := 0; y < z_end; y++ {
	// 		wire_vals := map[string]int{}
	// 		Set_x_y(wire_vals, x, y)

	// 		z_out := get_z_output(wire_vals, wire_gates)

	// 		if z_out != (x + y) {
	// 			return false
	// 		}
	// 	}

	// }
	// return true
}

func valid_up_to_test(z_end int, wire_gates map[string]gate) bool {
	if z_end == 0 {
		return true
	}

	sum := 0
	for i := 0; i < z_end; i++ {
		sum += int(math.Pow(2, float64(i)))
	}

	var x int
	var y int

	// Check 1+1
	x = int(math.Pow(2, float64(z_end)))
	y = int(math.Pow(2, float64(z_end)))
	wire_vals := map[string]int{}
	Set_x_y(wire_vals, x, y)
	z_out := get_z_output(wire_vals, wire_gates)
	z_expected := x + y
	for i := 0; i <= z_end; i++ {
		power := int(math.Pow(2, float64(i)))
		if (power & z_out) != (power & z_expected) {
			return false
		}
	}

	x = int(math.Pow(2, float64(z_end-1)))
	y = int(math.Pow(2, float64(z_end-1)))
	wire_vals = map[string]int{}
	Set_x_y(wire_vals, x, y)
	z_out = get_z_output(wire_vals, wire_gates)
	z_expected = x + y
	for i := 0; i <= z_end; i++ {
		power := int(math.Pow(2, float64(i)))
		if (power & z_out) != (power & z_expected) {
			return false
		}
	}

	// check 0+1 and 1+0
	x = int(math.Pow(2, float64(z_end)))
	y = 1
	wire_vals = map[string]int{}
	Set_x_y(wire_vals, x, y)
	z_out = get_z_output(wire_vals, wire_gates)
	z_expected = x + y
	for i := 0; i <= z_end; i++ {
		power := int(math.Pow(2, float64(i)))
		if (power & z_out) != (power & z_expected) {
			return false
		}
	}

	x = int(math.Pow(2, float64(z_end)))
	y = 1
	wire_vals = map[string]int{}
	Set_x_y(wire_vals, x, y)
	z_out = get_z_output(wire_vals, wire_gates)
	z_expected = x + y
	for i := 0; i <= z_end; i++ {
		power := int(math.Pow(2, float64(i)))
		if (power & z_out) != (power & z_expected) {
			return false
		}
	}

	// check 0 + 0
	x = int(math.Pow(2, float64(z_end-1)))
	y = 1
	wire_vals = map[string]int{}
	Set_x_y(wire_vals, x, y)
	z_out = get_z_output(wire_vals, wire_gates)
	z_expected = x + y
	for i := 0; i <= z_end; i++ {
		power := int(math.Pow(2, float64(i)))
		if (power & z_out) != (power & z_expected) {
			return false
		}
	}

	// Undo swap

	return true
}

func Set_x_y(wire_vals map[string]int, x int, y int) {
	for i := 0; i <= 44; i++ {
		x_bit := 0
		if x&int(math.Pow(2, float64(i))) > 0 {
			x_bit = 1
		}

		wire := "x" + get_padded_num(i)
		wire_vals[wire] = x_bit
	}

	for j := 0; j <= 44; j++ {
		j_bit := 0
		if y&int(math.Pow(2, float64(j))) > 0 {
			j_bit = 1
		}

		wire := "y" + get_padded_num(j)
		wire_vals[wire] = j_bit
	}
}

func get_possible_swaps(wire_gates map[string]gate, current_swaps [][2]string) [][2]string {
	possible_swaps := [][2]string{}

	wires := []string{}
	for wire := range wire_gates {
		wires = append(wires, wire)
	}

	skip := map[string]struct{}{}
	for _, current_swap := range current_swaps {
		skip[current_swap[0]] = struct{}{}
		skip[current_swap[1]] = struct{}{}
	}

	for i := range wires {
		for j := i + 1; j < len(wires); j++ {
			i_wire := wires[i]
			_, skip_i := skip[i_wire]
			if skip_i {
				continue
			}

			j_wire := wires[j]
			_, skip_j := skip[j_wire]
			if skip_j {
				continue
			}
			possible_swap := [2]string{i_wire, j_wire}

			tmp := possible_swap[:]
			sort.Strings(tmp)
			copy(tmp, possible_swap[:])

			possible_swaps = append(possible_swaps, possible_swap)
		}
	}

	return possible_swaps
}

func get_wire_val(wire string, wire_vals map[string]int, wire_gates map[string]gate, loop_detection map[string]struct{}) int {
	// println("in get_wire_val")
	val, exists := wire_vals[wire]
	if exists {
		return val
	}

	_, seen := loop_detection[wire]
	if seen {
		return -1
	}

	loop_detection[wire] = struct{}{}

	gate, gate_exists := wire_gates[wire]
	if !gate_exists {
		panic("Unexpected wire")
	}

	left_val := get_wire_val(gate.left, wire_vals, wire_gates, loop_detection)
	right_val := get_wire_val(gate.right, wire_vals, wire_gates, loop_detection)

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
