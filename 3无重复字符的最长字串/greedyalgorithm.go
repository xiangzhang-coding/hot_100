func lengthOfLongestSubstring(s string) int {
	var result int = 0
	lenOfS := len(s)
	charMap := make(map[byte]int)
	// l and r are two ends in a substring
	// if the s[r]'s count is equal to 2 ,
	// s[l]'s count should minus one and move right over and over again
	// so that the substring has no repeating character.
	l := 0
	for r := 0; r < lenOfS; r++ {
		if _, exist := charMap[s[r]]; exist {
			charMap[s[r]]++
		} else {
			charMap[s[r]] = 1
		}
		// use for not if
		for charMap[s[r]] == 2 {
			charMap[s[l]]--
			l++
		}
		// don't forget '+1'
		if r-l+1 > result {
			result = r - l + 1
		}
	}

	return result
}