func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	r1 := (m + n + 1) / 2
	r2 := (m + n + 2) / 2

	// divide the problem into 2 subquestion.
	r1 = findKth(nums1, 0, nums2, 0, r1)
	r2 = findKth(nums1, 0, nums2, 0, r2)

	return 1. * float64(r1+r2) / 2
}

// sta is the initial index of nums1,
// stb is the initial index of nums2.
// k is the middle of [len(nums1)+len(nums2)]/2
func findKth(nums1 []int, sta int, nums2 []int, stb int, k int) int {

	lena := len(nums1)
	lenb := len(nums2)

	if sta >= lena {
		return nums2[stb+k-1]
	}
	if stb >= lenb {
		return nums1[sta+k-1]
	}
	if k == 1 {
		if nums1[sta] <= nums2[stb] {
			return nums1[sta]
		} else {
			return nums2[stb]
		}
	}
	var (
		aaa int
		bbb int
	)
	half := k / 2
	if sta+half-1 >= lena {
		aaa = lena - 1
	} else {
		aaa = sta + half - 1
	}
	if stb+half-1 >= lenb {
		bbb = lenb - 1
	} else {
		bbb = stb + half - 1
	}

	if nums1[aaa] <= nums2[bbb] {
		excluded := aaa - sta + 1
		return findKth(nums1, aaa+1, nums2, stb, k-excluded)
	} else {
		excluded := bbb - stb + 1
		return findKth(nums1, sta, nums2, bbb+1, k-excluded)
	}

}