package task3LT

func hasDuplicatesInString(s string, sub string) bool {
	for i := 0; i < len(s); i++ {
		if sub == string(s[i]) {
			return true
		}
	}

	return false
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	if len(s) == 2 {
		if string(s[0]) == string(s[1]) {
			return 1
		} else {
			return 2
		}
	}

	max_len := -1
	for i := 0; i < len(s)-1; i++ {
		substr := string(s[i])
		for j := i + 1; j < len(s); j++ {
			if !hasDuplicatesInString(string(substr), string(s[j])) {
				substr += string(s[j])
			} else {
				max_len = Max(len(string(substr)), max_len)
				break
			}
		}
		max_len = Max(len(string(substr)), max_len)
	}
	return max_len
}
