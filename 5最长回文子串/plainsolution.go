func longestPalindrome(s string) string {
	longestP := 1
	var res1 int
	var res2 int
	for i := range s {
		r1, r2 := single(i, s)
		if r2-r1+1 > longestP {
			longestP = r2 - r1 + 1
			res1 = r1
			res2 = r2
		}
	}
	for i := 0; i < len(s)-1; i++ {
		r1, r2 := dual(i, i+1, s)
		if r2-r1+1 > longestP {
			longestP = r2 - r1 + 1
			res1 = r1
			res2 = r2
		}
	}
	return s[res1 : res2+1]
}

func single(index int, s string) (int, int) {
	n := len(s)
	longest := 1
	var r1 int
	var r2 int
	var left int = index - 1
	var right int = index + 1
	for left >= 0 && right < n {
		if s[left] == s[right] {

			if right-left+1 > longest {
				longest = right - left + 1
				r1 = left
				r2 = right
			}

			left--
			right++
		} else {
			break
		}

	}
	return r1, r2
}

func dual(left int, right int, s string) (int, int) {
	n := len(s)
	longest := 1
	var r1 int
	var r2 int
	for left >= 0 && right < n {
		if s[left] == s[right] {

			if right-left+1 > longest {
				longest = right - left + 1
				r1 = left
				r2 = right
			}

			left--
			right++
		} else {
			break
		}

	}
	return r1, r2
}