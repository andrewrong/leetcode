package solution

func reverseStr(s string, k int) string {
	if k == 0 {
		return s
	}

	b := []byte(s)

	size := len(b)
	if size == 0 {
		return s
	}

	for i := 0; i < size; i += 2 * k {
		left := i
		right := i + k - 1

		if right >= size {
			right = size - 1
		}

		for left < right {
			b[left], b[right] = b[right], b[left]
			left += 1
			right -= 1
		}
	}

	return string(b)
}
