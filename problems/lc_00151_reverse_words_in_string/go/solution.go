package solution

// 前后一起遍历，分别找到对应的单词进行调转，不过这边有一个问题是：长度不一致的问题，不同的word之间的长度是不一样的，这块需要如何解决呢?
func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	size := len(runes)

	// step1: delete necessary spaces
	writable := 0
	start := -1
	end := -1
	for i := 0; i < size; i += 1 {
		if runes[i] != ' ' && start == -1 {
			start = i
			continue
		}

		if runes[i] == ' ' && start != -1 {
			end = i - 1

			for j := start; j <= end; j += 1 {
				runes[writable] = runes[j]
				writable += 1
			}
			runes[writable] = ' '
			writable += 1

			start = -1
			end = -1
		}
	}

	if start != -1 {
		for i := start; i < size; i++ {
			runes[writable] = runes[i]
			writable += 1
		}
	} else {
		if writable > 0 {
			writable -= 1
		}
	}

	runes = runes[:writable]
	// reverse word
	size = len(runes)
	for left, right := 0, size-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}

	// per word reverse
	start = 0
	for start < size {
		end := start
		for end < size && runes[end] != ' ' {
			end += 1
		}

		for i, j := start, end-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		start = end + 1
	}

	return string(runes)
}
