package day12

import "testing"

func TestDayExampleP1(t *testing.T) {
	example := `AAAA
BBCD
BBCC
EEEC`
	result := Solve1(example)
	expect := 140

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExample2P1(t *testing.T) {
	example := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	result := Solve1(example)
	expect := 772

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExample1P1(t *testing.T) {
	example := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	result := Solve1(example)
	expect := 1930

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExampleP2(t *testing.T) {
	example := `AAAA
BBCD
BBCC
EEEC`
	result := Solve2(example)
	expect := 31

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}
