package day17

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Solve1(input_lines string) string {
	register_a, register_b, register_c, program := parse_input(input_lines)

	// Calc difference when sorted
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

	// Crack it like a safe
	// register_a := 23
	// for power := 0; power < len(program); power++ {
	// 	for j := 0; j < 8; j++ {
	// 		register_a_guess := register_a + j*int(math.Pow(8, float64(power)))
	// 		computer := Computer{
	// 			register_a: register_a_guess,
	// 			register_b: register_b,
	// 			register_c: register_c,
	// 			program:    program}

	// 		computer.run_program()

	// 		output := computer.get_output_ints()

	// 		if reflect.DeepEqual(output, program[:power+1]) {
	// 			register_a = register_a_guess
	// 			break
	// 		}
	// 	}
	// }

	// Calc initial program
	output_str := []string{}
	for _, num := range program {
		num_str := strconv.Itoa(num)
		output_str = append(output_str, num_str)
	}
	initial_program := strings.Join(output_str, ",")
	// initial_program_len := len(initial_program)

	register_a := 0
	for power := len(program) - 1; power >= 0; power-- {
		for j := 7; j >= 0; j-- {
			register_a_guess := register_a + j*int(math.Pow(8, float64(power)))

			if register_a_guess == 0 {
				continue
			}
			println("register_a_guess is ", register_a_guess)
			computer := Computer{
				register_a: register_a_guess,
				register_b: register_b,
				register_c: register_c,
				program:    program}

			computer.run_program()

			output := computer.get_output_ints()
			output_str := computer.get_output()

			println("Power: ", power, ", j is: ", j, ", output:", output_str)

			if initial_program == output_str {
				return register_a_guess
			}

			if power-1 >= 0 {
				if output[power-1] == program[power-1] {
					register_a = register_a_guess
				}

			}

		}
	}
	return 0

	// register_a := 0
	// for power := 0; power < len(program); power++ {
	// 	for j := 1; j < 8; j++ {
	// 		register_a_guess := register_a + j*int(math.Pow(8, float64(power)))
	// 		println("register_a_guess is ", register_a_guess)
	// 		computer := Computer{
	// 			register_a: register_a_guess,
	// 			register_b: register_b,
	// 			register_c: register_c,
	// 			program:    program}

	// 		computer.run_program()

	// 		output := computer.get_output_ints()
	// 		output_str := computer.get_output()

	// 		println("Power: ", power, ", j is: ", j, ", output:", output_str)

	// 		if power > 0 && output[power-1] == program[power-1] {
	// 			register_a = register_a_guess
	// 		}

	// 		if initial_program == output_str {
	// 			return register_a
	// 		}
	// 	}
	// }

	return register_a
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
