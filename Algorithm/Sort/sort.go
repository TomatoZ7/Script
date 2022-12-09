package sort

func BubbleSort(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		// Bubble Flag
		hasBubble := false
		for j := 0; j < length-1-i; j++ {
			if arr[j+1] < arr[j] {
				// Bubble
				arr[j+1], arr[j] = arr[j], arr[j+1]
				// Flag
				hasBubble = true
			}
		}

		if !hasBubble {
			break
		}
	}
}

func InsertionSort(arr []int) {
	length := len(arr)

	for i := 1; i < length; i++ {
		value := arr[i]
		j := i - 1

		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}

		arr[j+1] = value
	}
}

func SelectionSort(arr []int) {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		minIdx := i

		for j := i + 1; j < length; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}

		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
}

func MergeSort(arr []int) []int {
	length := len(arr)

	if length == 1 {
		return arr
	}

	// Separate
	mid := length / 2
	leftArr := MergeSort(arr[:mid])
	rightArr := MergeSort(arr[mid:])
	// Length
	leftLen, rightLen := len(leftArr), len(rightArr)
	// Index
	i, j, resIdx := 0, 0, 0
	// Result
	res := make([]int, leftLen+rightLen)

	// Merge
	for i < leftLen && j < rightLen {
		if leftArr[i] < rightArr[j] {
			res[resIdx] = leftArr[i]
			resIdx++
			i++
		} else {
			res[resIdx] = rightArr[j]
			resIdx++
			j++
		}
	}

	for ; i < leftLen; i++ {
		res[resIdx] = leftArr[i]
		resIdx++
	}

	for ; j < rightLen; j++ {
		res[resIdx] = rightArr[j]
		resIdx++
	}

	return res
}

func QuickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	pivot := arr[end]
	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]

	QuickSort(arr, start, i-1)
	QuickSort(arr, i+1, end)
}
