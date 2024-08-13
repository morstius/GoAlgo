package algorithms

import "fmt"

func print(arr []int, lo int, hi int) {
	for i := lo; i < hi; i++ {
		fmt.Println(arr[i])
	}
}

func merge(arr []int, lo int, hi int, m int) {
	i := lo
	j := lo + m
	for i <= j && j < hi {
		if arr[j] < arr[i] {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i++
		} else {
			i++
		}
	}
}

func mSort(arr []int, lo int, hi int) {
	m := (hi - lo) / 2

	if (hi - lo) <= 1 {
		return
	} else {
		mSort(arr, lo, lo+m)
		mSort(arr, lo+m, hi)
		merge(arr, lo, hi, m)
	}
}

func MergeSort(arr []int) {
	mSort(arr, 0, len(arr))
}
