package day06

import "testing"

func TestDayExampleP1(t *testing.T) {
	example := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	result := Solve1(example)
	expect := 41

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExampleP2(t *testing.T) {
	example := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	result := Solve2(example)
	expect := 31

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}
