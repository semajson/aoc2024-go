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

func TestDayExample2P1(t *testing.T) {
	example := `Register A: 30553366
Register B: 0
Register C: 0

Program: 2,4,1,1,7,5,4,7,1,4,0,3,5,5,3,0`
	result := Solve1(example)
	expect := "1,3,7,4,6,4,2,3,5"

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %s; \nWanted: %s", example, result, expect)
	}
}

func TestDayExampleP2(t *testing.T) {
	example := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`
	result := Solve2(example)
	expect := 117440

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExample2P2(t *testing.T) {
	example := `Register A: 30553366
Register B: 0
Register C: 0

Program: 2,4,1,1,7,5,4,7,1,4,0,3,5,5,3,0`
	result := Solve2(example)
	expect := 117440

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}
