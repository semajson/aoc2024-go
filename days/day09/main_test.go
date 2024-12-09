package day09

import "testing"

func TestDayExampleP1(t *testing.T) {
	example := `2333133121414131402`
	result := Solve1(example)
	expect := 1928

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExample2P1(t *testing.T) {
	example := `2222222222222222222222222222222222222222222222`
	result := Solve1(example)
	expect := 1928

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExampleP2(t *testing.T) {
	example := `2333133121414131402`
	// example := `2333133121414131404`
	result := Solve2(example)
	expect := 2858

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayFreeSpace(t *testing.T) {
	example := []int{2, -1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	result := First_free_space(example, 4, 3)
	expect := 1

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}
