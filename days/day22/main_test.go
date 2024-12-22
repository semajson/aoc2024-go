package day22

import "testing"

func TestDayExampleP1(t *testing.T) {
	example := `1
10
100
2024`
	result := Solve1(example)
	expect := 37327623

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExampleP2(t *testing.T) {
	example := `1
10
100
2024`
	result := Solve2(example)
	expect := 31

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}
