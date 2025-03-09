type NumWithIndex struct {
	value int
	index int
}

func twoSum(nums []int, target int) []int {
	s := make([]NumWithIndex, len(nums))
	for i := range nums {
		s[i] = NumWithIndex{nums[i], i}
	}
	// sort.Slice第一个参数是待排序的切片，第二个参数是一个函数，该函数接收两个参数，返回true表示前一个参数应该排在后一个参数前面
	sort.Slice(s, func(i, j int) bool {
		return s[i].value < s[j].value
	})
	left, right := 0, len(s)-1
	for left < right {
		sum := s[left].value + s[right].value
		if sum == target {
			return []int{s[left].index, s[right].index}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}