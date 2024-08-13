package test

import (
	"Algo/algorithms"
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var size int = 10000
	arr := rndArray(size)
	check := sortedArray(arr, size)
	algorithms.QuickSort(arr)

	if reflect.DeepEqual(arr, check) == false {
		t.Fatal("QuickSort failed...")
	}
	fmt.Println("QuickSort Success...")
}

func BenchmarkQuickSort(b *testing.B) {
	arr := rndArray(10000)

	for i := 0; i < b.N; i++ {
		algorithms.QuickSort(arr)
	}
}
