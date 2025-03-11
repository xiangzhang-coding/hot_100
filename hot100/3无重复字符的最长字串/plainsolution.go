func lengthOfLongestSubstring(s string) int {
	// res is the length of current substring with no repeating characters
	var res int = 0
	// maxValue is the maximum length of substring with no repeating characters
	var maxValue int = 0
	length := len(s)
	// O(n) space complexity to store the characters in another formation
	runes := []rune(s)
	if length == 1 {
		return 1
	}
	if length == 0 {
		return 0
	}
	for index, _ := range runes {
		// O(n^2) space complexity to store the characters in a map
		// use a map to store that if the character is already in the substring
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
