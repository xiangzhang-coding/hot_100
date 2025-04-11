func longestPalindrome(s string) string {
	n := len(s)
	table := make([][]bool, n)
	for i := range table {
		table[i] = make([]bool, n)
	}
	maxlen := 1
	start := 0

	for i := 0; i < n; i++ {
		table[i][i] = true
	}

	for len := 2; len <= n; len++ {
		for i := 0; i+len <= n; i++ {
			if s[i] != s[i+len-1] {
				continue
			}
			if len == 2 {
				table[i][i+1] = true
				if len > maxlen {
					maxlen = len
					start = i
				}
			}
			if len > 2 && table[i+1][i+len-2] == true {
				table[i][i+len-1] = true
				if len > maxlen {
					maxlen = len
					start = i
				}
			}
		}
	}

	return s[start : start+maxlen]

}