package day21

import "testing"

func TestDayExampleP1(t *testing.T) {
	example := `029A
980A
179A
456A
379A`
	result := Solve1(example)
	expect := 126384

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExample2P1(t *testing.T) {
	example := `179A`
	result := Solve1(example)
	expect := 68 * 179

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExample3P1(t *testing.T) {
	example := `286A
974A
189A
802A
805A`
	result := Solve1(example)
	expect := 68 * 179

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExampleP2(t *testing.T) {
	example := `029A
980A
179A
456A
379A`
	result := Solve2(example)
	expect := 31

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}
