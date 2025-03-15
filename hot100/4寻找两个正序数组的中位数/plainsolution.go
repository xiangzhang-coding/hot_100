func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	ll := make([]int, 0)
	for _, v := range nums1 {
		ll = append(ll, v)
	}
	for _, v := range nums2 {
		ll = append(ll, v)

	}
	quicksort(ll, 0, m+n-1)

	if (m+n)%2 == 1 {
		return float64(ll[(m+n-1)/2])
	} else {
		return float64((ll[(m+n-1)/2] + ll[(m+n)/2])) / 2
	}
}

// use quick sort , heap sort or merged sort to make time complexity O(nlogn)
// here is a quick sort implementation
func quicksort(num []int, low int, high int) {
	if low < high {
		pivot := partition(num, low, high)
		quicksort(num, low, pivot-1)
		quicksort(num, pivot+1, high)
	}
}

// auxiliary function
// swap function
// arr[i], arr[j] = arr[j], arr[i] // equivalent form
// a, b = b, a
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// partition function
func partition(num []int, low int, high int) int {
	// a pivot is needed in quicksort
	pivot := num[low]
	for low < high {
		for low < high && num[high] >= pivot {
			high--
		}
		swap(&num[low], &num[high])
		for low < high && num[low] <= pivot {
			low++
		}
		swap(&num[low], &num[high])
	}
	return low

}
