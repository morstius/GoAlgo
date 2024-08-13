package test

import (
	"Algo/algorithms"
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var size int = 10
	arr := rndArray(size)
	check := sortedArray(arr, 10)
	algorithms.MergeSort(arr)

	if reflect.DeepEqual(arr, check) == false {
		t.Fatal("MergeSort failed...")
	}
	fmt.Println("MergeSort Success...")
}

func BenchmarkMergeSort(b *testing.B) {
	arr := rndArray(10000)

	for i := 0; i < b.N; i++ {
		algorithms.MergeSort(arr)
	}
}
