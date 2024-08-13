package algorithms

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	idx := lo - 1

	for j := lo; j < hi; j++ {
		if a[j] <= pivot {
			idx++
			temp := a[j]
			a[j] = a[idx]
			a[idx] = temp
		}
	}
	idx++
	a[hi] = a[idx]
	a[idx] = pivot
	return idx
}

func qSort(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	} else {
		mid := partition(arr, lo, hi)
		qSort(arr, lo, mid-1)
		qSort(arr, mid+1, hi)
	}
}

func QuickSort(arr []int) {
	qSort(arr, 0, len(arr)-1)
}
