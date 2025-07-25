package lc_00059_spiral_matrix_2

func Solve(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	stepX := 0
	stepY := 0

	leftX := 0
	rightX := n - 1
	upperY := 0
	bottomY := n - 1

	direction := "right"

	for i := 1; i <= n*n; i++ {
		/**
		 * 1. 先往X正方向走
		 * 2. 遇到最右边，就往Y正方向走
		 * 3. 遇到了Y的最下面，并且X已经处于最右边，就往X反方向走
		 * 4. 遇到X最为左边，并且处于Y的最下面，就再往Yf反向走
		 * 边界会在走动的变化过程中
		 */
		switch direction {
		case "right":
			{
				result[stepX][stepY] = i

				tmpX := stepX + 1
				if tmpX <= rightX {
					stepX = tmpX
				} else {
					upperY += 1
					stepY += 1
					direction = "down"
				}
			}
		case "down":
			{
				result[stepX][stepY] = i

				tmpY := stepY + 1
				if tmpY <= bottomY {
					stepY = tmpY
				} else {
					rightX -= 1
					direction = "left"
					stepX -= 1
				}
			}
		case "left":
			{
				tmpX := stepX - 1
				if tmpX >= leftX {
					stepX = tmpX
					result[stepX][stepY] = i
				} else {
					bottomY -= 1
					direction = "upper"
					goto Move
				}
			}
		case "upper":
			{
				tmpY := stepY - 1
				if tmpY >= upperY {
					stepY = tmpY
					result[stepX][stepY] = i
				} else {
					leftX += 1
					direction = "right"
					goto Move
				}
			}
		}
	}
	return result
}
