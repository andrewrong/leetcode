package lc_00042_trapping_rain_water

func trap(height []int) int {
	sum := 0
	for i := 0; i < len(height); i += 1 {
		if i == 0 || i == len(height)-1 {
			continue
		}

		lHeight := height[i]
		rHeight := height[i]

		for j := i + 1; j < len(height); j += 1 {
			if height[j] > rHeight {
				rHeight = height[j]
			}
		}

		for j := i - 1; j >= 0; j -= 1 {
			if height[j] > lHeight {
				lHeight = height[j]
			}
		}

		h := min(rHeight, lHeight) - height[i]
		if h > 0 {
			sum += h
		}
	}

	return sum
}

func trapDymanic(height []int) int {
	leftHeight := make([]int, len(height))
	rightHeight := make([]int, len(height))

	for i, v := range height {
		if i == 0 {
			leftHeight[i] = v
			continue
		}

		leftHeight[i] = max(v, leftHeight[i-1])
	}

	size := len(height)

	for i := size - 1; i >= 0; i -= 1 {
		if i == size-1 {
			rightHeight[i] = height[i]
			continue
		}
		rightHeight[i] = max(height[i], rightHeight[i+1])
	}

	sum := 0

	for i, v := range height {
		h := min(leftHeight[i], rightHeight[i]) - v
		if h > 0 {
			sum += h
		}
	}

	return sum
}
