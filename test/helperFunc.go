package test

import (
	"math/rand"
	"sort"
)

func rndArray(size int) []int {
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = rand.Int()
	}
	return arr
}

func sortedArray(arr []int, size int) []int {
	check := make([]int, size)
	copy(check, arr)
	sort.Ints(check)
	return check
}
