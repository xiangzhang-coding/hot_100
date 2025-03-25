func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	// two pointers
	i := 0
	j := 0
	// to handle mid <= 0
	if m == 0 && n == 0 {
		return float64(0)
	}
	if m == 0 && n == 1 {
		return float64(nums2[0])
	}
	if m == 1 && n == 0 {
		return float64(nums1[0])
	}
	if m == 1 && n == 1 {
		return float64(nums1[0]+nums2[0]) / 2
	}
	// set a flag
	var flag_odd bool
	if (m+n)%2 == 1 {
		flag_odd = true
	} else {
		flag_odd = false
	}
	// find the middle position
	mid := (m + n - 1) / 2
	for mid > 0 {
		if i < m && j < n {
			if nums1[i] <= nums2[j] {
				i++
			} else {
				j++
			}
			mid--
			continue
		}
		if i == m {
			j++
			mid--
		} else {
			i++
			mid--
		}
	}

	if flag_odd {
		if i == m {
			return float64(nums2[j])
		}
		if j == n {
			return float64(nums1[i])
		}
		if nums1[i] <= nums2[j] {
			return float64(nums1[i])
		} else {
			return float64(nums2[j])
		}
	} else {
		var r1 float64
		var r2 float64
		if i == m {
			r1 = float64(nums2[j])
			j++
			r2 = float64(nums2[j])
			return (r1 + r2) / 2
		}
		if j == n {
			r1 = float64(nums1[i])
			i++
			r2 = float64(nums1[i])
			return (r1 + r2) / 2
		}
		if nums1[i] <= nums2[j] {
			r1 = float64(nums1[i])
			i++
		} else {
			r1 = float64(nums2[j])
			j++
		}

		if i == m {
			r2 = float64(nums2[j])
			return (r1 + r2) / 2
		}
		if j == n {
			r2 = float64(nums1[i])
			return (r1 + r2) / 2
		}
		if nums1[i] <= nums2[j] {
			r2 = float64(nums1[i])
			i++
		} else {
			r2 = float64(nums2[j])
			j++
		}
		return (r1 + r2) / 2

	}
}