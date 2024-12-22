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
2
3
2024`
	result := Solve2(example)
	expect := 23

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExample3P2(t *testing.T) {

	result := Get_num_bananas(123, [4]int{-1, -1, 0, 2})
	expect := 6

	if result != expect {
		t.Errorf("Failed!\nEquals: %d; \nWanted: %d", result, expect)
	}
}

func TestDayExample4P2(t *testing.T) {

	result := Get_num_bananas(1, [4]int{-2, 1, -1, 3})
	expect := 7

	if result != expect {
		t.Errorf("Failed!\nEquals: %d; \nWanted: %d", result, expect)
	}
}
func TestDayExample5P2(t *testing.T) {

	result := Get_num_bananas(2, [4]int{-2, 1, -1, 3})
	expect := 7

	if result != expect {
		t.Errorf("Failed!\nEquals: %d; \nWanted: %d", result, expect)
	}
}

func TestDayExample6P2(t *testing.T) {

	result := Get_num_bananas(3, [4]int{-2, 1, -1, 3})
	expect := 0

	if result != expect {
		t.Errorf("Failed!\nEquals: %d; \nWanted: %d", result, expect)
	}
}

func TestDayExample7P2(t *testing.T) {

	result := Get_num_bananas(2024, [4]int{-2, 1, -1, 3})
	expect := 9

	if result != expect {
		t.Errorf("Failed!\nEquals: %d; \nWanted: %d", result, expect)
	}
}

func TestDayExample10P2(t *testing.T) {

	result := Get_num_bananas(1, [4]int{-9, 9, -1, 0})
	expect := 7

	if result != expect {
		t.Errorf("Failed!\nEquals: %d; \nWanted: %d", result, expect)
	}
}
func TestDayExample11P2(t *testing.T) {

	result := Get_num_bananas(2, [4]int{-9, 9, -1, 0})
	expect := 7

	if result != expect {
		t.Errorf("Failed!\nEquals: %d; \nWanted: %d", result, expect)
	}
}

func TestDayExample12P2(t *testing.T) {

	result := Get_num_bananas(3, [4]int{-9, 9, -1, 0})
	expect := 0

	if result != expect {
		t.Errorf("Failed!\nEquals: %d; \nWanted: %d", result, expect)
	}
}

func TestDayExample13P2(t *testing.T) {

	result := Get_num_bananas(2024, [4]int{-9, 9, -1, 0})
	expect := 9

	if result != expect {
		t.Errorf("Failed!\nEquals: %d; \nWanted: %d", result, expect)
	}
}
