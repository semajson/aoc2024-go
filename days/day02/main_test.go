package day02

import (
	"fmt"
	"testing"
)

func TestDayExampleP1(t *testing.T) {
	example := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	result := Solve1(example)
	expect := 2

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestIsSafe(t *testing.T) {
	example := []int{7, 6, 4, 2, 1}
	result := is_safe(example)
	expect := true

	if result != expect {
		fmt.Println("Failed!, input:", example)
		fmt.Println("Expected", expect)
		fmt.Println("got:", result)
		t.Errorf("Failure")
	}
}
func TestIsSafe2(t *testing.T) {
	example := []int{1, 3, 6, 7, 9}
	result := is_safe(example)
	expect := true

	if result != expect {
		fmt.Println("Failed!, input:", example)
		fmt.Println("Expected", expect)
		fmt.Println("got:", result)
		t.Errorf("Failure")
	}
}

func TestDayExampleP2(t *testing.T) {
	example := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	result := Solve2(example)
	expect := 4

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestIsSafe3(t *testing.T) {
	example := []int{34, 1, 2, 3}
	result := is_safe_with_dampener(example)
	expect := true

	if result != expect {
		fmt.Println("Failed!, input:", example)
		fmt.Println("Expected", expect)
		fmt.Println("got:", result)
		t.Errorf("Failure")
	}
}

func TestIsSafe4(t *testing.T) {
	example := []int{4, 3, 199, 2, 1}
	result := is_safe_with_dampener(example)
	expect := true

	if result != expect {
		fmt.Println("Failed!, input:", example)
		fmt.Println("Expected", expect)
		fmt.Println("got:", result)
		t.Errorf("Failure")
	}
}

func TestIsSafe45(t *testing.T) {
	example := []int{1, 2, 4, 5, 4}
	result := is_safe_with_dampener(example)
	expect := true

	if result != expect {
		fmt.Println("Failed!, input:", example)
		fmt.Println("Expected", expect)
		fmt.Println("got:", result)
		t.Errorf("Failure")
	}
}

func TestIsSafe46(t *testing.T) {
	example := []int{5, 1, 6, 7, 8}
	result := is_safe_with_dampener(example)
	expect := true

	if result != expect {
		fmt.Println("Failed!, input:", example)
		fmt.Println("Expected", expect)
		fmt.Println("got:", result)
		t.Errorf("Failure")
	}
}

func TestIsSafe47(t *testing.T) {
	example := []int{5, 2, 6, 7, 8}
	result := is_safe_with_dampener(example)
	expect := true

	if result != expect {
		fmt.Println("Failed!, input:", example)
		fmt.Println("Expected", expect)
		fmt.Println("got:", result)
		t.Errorf("Failure")
	}
}
