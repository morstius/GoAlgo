package test

import (
	"Algo/algorithms"
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	size := 10000
	arr := rndArray(size)
	check := sortedArray(arr, size)
	algorithms.BubbleSort(arr)

	if reflect.DeepEqual(arr, check) == false {
		t.Fatal("BubbleSort failed...")
	}
	fmt.Println("BubbleSort Success...")
}

func BenchmarkBubbleSort(b *testing.B) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	for i := 0; i < b.N; i++ {
		algorithms.BubbleSort(arr)
	}
}
