package algorithms

func BubbleSort(arr []int) {
	for i := len(arr) - 1; i > 1; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}
