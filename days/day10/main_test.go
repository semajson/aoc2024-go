package day10

import "testing"

func TestDayExampleP1(t *testing.T) {
	example := `0123
7654
89..
...`
	result := Solve1(example)
	expect := 1

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExample4P1(t *testing.T) {
	example := `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`
	result := Solve1(example)
	expect := 2

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExample2P1(t *testing.T) {
	example := `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`
	result := Solve1(example)
	expect := 4

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExample3P1(t *testing.T) {
	example := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	result := Solve1(example)
	expect := 36

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExampleP2(t *testing.T) {
	example := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	result := Solve2(example)
	expect := 81

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}
