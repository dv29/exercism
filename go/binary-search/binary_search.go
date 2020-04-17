package binarysearch

func SearchInts(sl []int, q int) int {
	if len(sl) == 0 {
		return -1
	}

	start, end := 0, len(sl)-1

	for start < end {
		mid := start + ((end - start) / 2)
		if q <= sl[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}

	if start < len(sl) && sl[start] == q {
		return start
	}
	return -1
}
