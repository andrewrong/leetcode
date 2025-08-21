package solution

func reverseString(s []byte) {
	// Implementation will go here
	size := len(s)

	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
