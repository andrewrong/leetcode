package solution

// Partition returns all possible palindrome partitioning of s
func Partition(s string) [][]string {
	one := make([]string, 0)
	all := make([][]string, 0)
	partition1(s, &one, &all)

	return all
}

func partition1(s string, one *[]string, answers *[][]string) {
	if len(s) == 0 {
		if len(*one) != 0 {
			tmp := make([]string, len(*one))
			copy(tmp, *one)
			*answers = append(*answers, tmp)
		}

		return
	}

	for i := 1; i <= len(s); i += 1 {
		part1 := s[:i]
		if !isParlindrom(part1) {
			continue
		}

		*one = append(*one, part1)
		if i == len(s) {
			partition1("", one, answers)
		} else {
			partition1(s[i:], one, answers)
		}
		*one = (*one)[:len(*one)-1]
	}
}

func isParlindrom(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] == s[j] {
			continue
		}
		return false
	}

	return true
}
