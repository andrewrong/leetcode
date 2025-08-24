// leetcode/problems/lc_00459_repeated_substring_pattern/go/solution.go
package main

func repeatedSubstringPattern(str string) bool {
	if len(str) == 0 || len(str) == 1 {
		return false
	}

	next := make([]int, len(str))

	for i := 1; i < len(str); i += 1 {
		// 前一个字符串的LPS
		j := next[i-1]
		for j > 0 && str[j] != str[i] {
			j = next[j-1]
		}

		if str[j] == str[i] {
			next[i] = j + 1
		} else {
			next[i] = 0
		}
	}

	if next[len(str)-1] == 0 {
		return false
	}

	k := len(str) - next[len(str)-1]

	if len(str)%k == 0 {
		return true
	}
	return false
}
