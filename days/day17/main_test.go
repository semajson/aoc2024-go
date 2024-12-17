package day17

import "testing"

func TestDayExampleP1(t *testing.T) {
	example := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	result := Solve1(example)
	expect := "4,6,3,5,6,3,5,2,1,0"

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %s; \nWanted: %s", example, result, expect)
	}
}

func TestDayExampleP2(t *testing.T) {
	example := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	result := Solve2(example)
	expect := "4,6,3,5,6,3,5,2,1,0"

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %s; \nWanted: %s", example, result, expect)
	}
}
