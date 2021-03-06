package minimumSwaps

func minimumSwaps(arr []int32) int32 {
	if len(arr) == 0 {
		return 0
	}
	var swaps int32
	curMinimumIndex := minimumArrayItemIndex(arr[0:])
	if arr[0] > arr[curMinimumIndex] {
		arr[0], arr[curMinimumIndex] = arr[curMinimumIndex], arr[0]
		swaps++
	}
	return swaps + minimumSwaps(arr[1:])
}

func minimumArrayItemIndex(arr []int32) int {
	var minimum = 0
	for i, v := range arr {
		if v < arr[minimum] {
			minimum = i
		}
	}
	return minimum
}
