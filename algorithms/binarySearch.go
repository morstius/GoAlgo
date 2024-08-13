package algorithms

func BS_list(haystack []int, needle int) bool {
	var lo = 0
	var hi = len(haystack)

	for lo < hi {
		var m int = lo + (hi-lo)/2
		v := haystack[m]

		if v == needle {
			return true
		} else if needle < v {
			hi = m
		} else {
			lo = m + 1
		}
	}

	return false
}
