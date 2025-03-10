func lengthOfLongestSubstring(s string) int {
	var res int = 0
	var maxValue int = 0
	length := len(s)
	runes := []rune(s)
	if length == 1 {
		return 1
	}
	if length == 0 {
		return 0
	}
	for index, _ := range runes {
		charMap := make(map[rune]bool)
		for j := index; j < length; j++ {

			if _, exists := charMap[runes[j]]; exists {

				break
			} else {
				charMap[runes[j]] = true
				res = j - index + 1
				if res > maxValue {
					maxValue = res
				}

			}
		}

	}
	return maxValue
}
