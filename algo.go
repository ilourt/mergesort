package mergesort

// Sort list of int using merge algorithm
func Sort(list []int) []int {
	ln := len(list)
	if ln <= 1 {
		return list
	}

	left := Sort(list[:ln/2])
	right := Sort(list[ln/2:])

	return merge(left, right)
}

// merge right and left part
func merge(left, right []int) []int {
	list := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(list, right...)
		}

		if len(right) == 0 {
			return append(list, left...)
		}

		if right[0] < left[0] {
			list = append(list, right[0])
			right = right[1:]
		} else {
			list = append(list, left[0])
			left = left[1:]
		}
	}

	return list
}
