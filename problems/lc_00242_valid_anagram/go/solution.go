package solution

// IsAnagram checks if two strings are anagrams of each other
func Solve(s string, t string) bool {
	// Implementation will be added later

	if len(s) != len(t) {
		return false
	}

	var record [26]int
	for i := 0; i < len(s); i += 1 {
		record[s[i]-'a'] += 1
		record[t[i]-'a'] -= 1
	}

	for _, element := range record {
		if element != 0 {
			return false
		}
	}

	return true
}
