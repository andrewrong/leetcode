package solution

import "sort"

// FindContentChildren finds the maximum number of children that can be satisfied with cookies
// g is the array of children's appetite values
// s is the array of cookie sizes
func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	sum := 0
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] <= s[j] {
			sum += 1
			i += 1
			j += 1
		} else {
			j += 1
		}
	}
	return sum
}
