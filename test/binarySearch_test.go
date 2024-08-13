package test

import (
	"Algo/algorithms"
	"fmt"
	"testing"
)

func check(arr []int, c int, expected bool, t *testing.T) {
	if res := algorithms.BS_list(arr, c); res != expected {
		t.Fatalf("Failed for %d", c)
	}
}

func TestBS(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	// Test cases
	check(foo, 69, true, t)
	check(foo, 1336, false, t)
	check(foo, 69420, true, t)
	check(foo, 69421, false, t)
	check(foo, 1, true, t)
	check(foo, 0, false, t)
	fmt.Println("BinarySearch Success...")
}
