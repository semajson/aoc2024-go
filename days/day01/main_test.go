package day01

import "testing"

func TestDay01ExampleP1(t *testing.T) {
	example := `3   4
4   3
2   5
1   3
3   9
3   3`
	result := Solve1(example)
	expect := 11

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDay01ExampleP2(t *testing.T) {
	example := `3   4
4   3
2   5
1   3
3   9
3   3`
	result := Solve2(example)
	expect := 31

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}
