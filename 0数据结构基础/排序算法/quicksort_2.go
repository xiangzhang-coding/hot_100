// partition function
// return the index of pivot
func partition(arr []int, low, high int) int {
	pivot := arr[high] // choose the last element as pivot
	i := low - 1

	// iterate from low to high-1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // swap
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low, high int) {
	if low < high { // don't forget the exit condition
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}