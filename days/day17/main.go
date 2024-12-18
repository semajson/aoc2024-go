package day17

import (
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func Solve1(input_lines string) string {
	register_a, register_b, register_c, program := parse_input(input_lines)

	computer := Computer{
		register_a: register_a,
		register_b: register_b,
		register_c: register_c,
		program:    program}

	computer.run_program()

	return computer.get_output()
}

func Solve2(input_lines string) int {
	_, register_b, register_c, program := parse_input(input_lines)

	// BFS guessing one 3-bit number at a time
	// Importantly, work backwards from last 3-bit number
	// This works as the input program is a particular format
	a_guesses := []int{0}
	for i := len(program) - 1; i >= 0; i-- {
		a_guesses_new := []int{}
		for j := 0; j < 8; j++ {
			for _, a_guess := range a_guesses {
				new_a_guess := a_guess*8 + j
				computer := Computer{
					register_a: new_a_guess,
					register_b: register_b,
					register_c: register_c,
					program:    program}

				computer.run_program()
				output := computer.get_output_ints()

				if reflect.DeepEqual(program[i:], output) {
					a_guesses_new = append(a_guesses_new, new_a_guess)
				}
			}
		}
		a_guesses = a_guesses_new
	}

	// Select minimum guess
	min := -1
	for _, guess := range a_guesses {
		if min == -1 || guess < min {
			min = guess
		}
	}

	return min
}

type Computer struct {
	register_a          int
	register_b          int
	register_c          int
	program             []int
	output              []int
	instruction_pointer int
	skip_next_jump      bool // Not super nice, maybe refactor
}

func (computer *Computer) run_program() {
	computer.instruction_pointer = 0
	computer.skip_next_jump = false

	for computer.instruction_pointer < len(computer.program) {
		index := computer.instruction_pointer
		instruction := computer.program[index]
		operand := computer.program[index+1]

		computer.do_instruction(instruction, operand)

		if !computer.skip_next_jump {
			computer.instruction_pointer += 2
		}
		computer.skip_next_jump = false
	}
}

func (computer *Computer) do_instruction(instruction int, operand int) {
	switch instruction {
	case 0:
		computer.do_adv(operand)
	case 1:
		computer.do_bxl(operand)
	case 2:
		computer.do_bst(operand)
	case 3:
		computer.do_jnz(operand)
	case 4:
		computer.do_bxc(operand)
	case 5:
		computer.do_out(operand)
	case 6:
		computer.do_bdv(operand)
	case 7:
		computer.do_cdv(operand)
	default:
		panic("Unexpected command!")
	}
}

func (computer *Computer) do_adv(operand int) {
	// Division
	numerator := float64(computer.register_a)
	denominator := math.Pow(2, float64(computer.combo_operand(operand)))

	result := int(numerator / denominator)
	computer.register_a = result
}
func (computer *Computer) do_bxl(operand int) {
	// Bitwise XOR
	result := computer.register_b ^ operand
	computer.register_b = result

}
func (computer *Computer) do_bst(operand int) {
	// Combo mod 8
	result := computer.combo_operand(operand) % 8
	computer.register_b = result
}
func (computer *Computer) do_jnz(operand int) {
	if computer.register_a != 0 {
		computer.instruction_pointer = operand
		computer.skip_next_jump = true
	}
}
func (computer *Computer) do_bxc(operand int) {
	if computer.register_c == 0 {
		// println("here")
	}
	result := computer.register_b ^ computer.register_c
	computer.register_b = result
}
func (computer *Computer) do_out(operand int) {
	result := computer.combo_operand(operand) % 8
	computer.output = append(computer.output, result)
}
func (computer *Computer) do_bdv(operand int) {
	// Division
	numerator := float64(computer.register_a)
	denominator := math.Pow(2, float64(computer.combo_operand(operand)))

	result := int(numerator / denominator)
	computer.register_b = result
}
func (computer *Computer) do_cdv(operand int) {
	// Division
	numerator := float64(computer.register_a)
	denominator := math.Pow(2, float64(computer.combo_operand(operand)))

	result := int(numerator / denominator)
	computer.register_c = result
}

func (computer Computer) get_output() string {
	output_str := []string{}

	for _, num := range computer.output {
		num_str := strconv.Itoa(num)
		output_str = append(output_str, num_str)
	}
	return strings.Join(output_str, ",")
}

func (computer Computer) get_output_ints() []int {
	return computer.output
}

func (computer Computer) combo_operand(operand int) int {
	switch operand {
	case 0:
		return operand
	case 1:
		return operand
	case 2:
		return operand
	case 3:
		return operand
	case 4:
		return computer.register_a
	case 5:
		return computer.register_b
	case 6:
		return computer.register_c
	case 7:
		panic("Invalid use of reserved operand 7")
	default:
		panic("Non base-3 operand")
	}
}

func parse_input(input_lines string) (int, int, int, []int) {
	re, _ := regexp.Compile(`Register A: (\d+)\nRegister B: (\d+)\nRegister C: (\d+)\n\nProgram: ([\d?,?]+)`)

	matches := re.FindAllStringSubmatch(input_lines, -1)
	match := matches[0]

	register_a, _ := strconv.Atoi(match[1])
	register_b, _ := strconv.Atoi(match[2])
	register_c, _ := strconv.Atoi(match[3])

	program := []int{}
	program_raw := match[4]
	for _, num_raw := range strings.Split(program_raw, ",") {
		num, _ := strconv.Atoi(num_raw)
		program = append(program, num)
	}

	return register_a, register_b, register_c, program
}
