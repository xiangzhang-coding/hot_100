func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if remain_index, exists := numMap[diff]; exists {
			return []int{remain_index, i}
		} else {
			numMap[num] = i
		}
	}
	return []int{}
}