package lc_00059_spiral_matrix_2

import "fmt"

func Solve(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	// // start position
	// stepX := 0
	// stepY := 0

	// leftY := 0
	// rightY := n - 1
	// upperX := 0
	// bottomX := n - 1

	// direction := "right"

	// for i := 1; i <= n*n; i++ {
	// 	/**
	// 	 * 1. 先往 X 正方向走
	// 	 * 2. 遇到最右边，就往 Y 正方向走
	// 	 * 3. 遇到了 Y 的最下面，并且 X 已经处于最右边，就往 X 反方向走
	// 	 * 4. 遇到 X 最为左边，并且处于 Y 的最下面，就再往 Yf 反向走
	// 	 * 边界会在走动的变化过程中
	// 	 */
	// 	switch direction {
	// 	case "right":
	// 		{
	// 			result[stepX][stepY] = i
	// 			fmt.Printf("%s-(%v, %v) = %v\n", direction, stepY, stepX, i)

	// 			tmpY := stepY + 1
	// 			if tmpY <= rightY {
	// 				stepY = tmpY
	// 			} else {
	// 				upperX += 1
	// 				stepX += 1
	// 				direction = "down"
	// 			}
	// 		}
	// 	case "down":
	// 		{
	// 			result[stepX][stepY] = i
	// 			fmt.Printf("(%v, %v) = %v\n", stepY, stepX, i)

	// 			tmpX := stepX + 1
	// 			if tmpX <= bottomX {
	// 				stepX = tmpX
	// 			} else {
	// 				rightY -= 1
	// 				stepY -= 1
	// 				direction = "left"
	// 			}
	// 		}
	// 	case "left":
	// 		{
	// 			result[stepX][stepY] = i
	// 			fmt.Printf("(%d, %d) = %d\n", stepY, stepX, i)

	// 			tmpY := stepY - 1
	// 			if tmpY >= leftY {
	// 				stepY = tmpY
	// 			} else {
	// 				bottomX -= 1
	// 				stepX -= 1
	// 				direction = "upper"
	// 			}
	// 		}
	// 	case "upper":
	// 		{

	// 			result[stepX][stepY] = i
	// 			fmt.Printf("(%d, %d) = %d\n", stepY, stepX, i)

	// 			tmpX := stepX - 1
	// 			if tmpX >= upperX {
	// 				stepX = tmpX
	// 			} else {
	// 				leftY += 1
	// 				stepY += 1
	// 				direction = "right"

	// 			}
	// 		}
	// 	}
	// }
	robot := NewRobot(n)
	for i := 1; i <= n*n; i++ {
		x, y := robot.Move()
		result[x][y] = i
	}
	return result
}

// 可以作为状态机的方式来编写
type Robot interface {
	Move() (int, int)
	Reset()
}

type MatrixRobot struct {
	n         int
	x         int
	y         int
	currStep  int
	direction string

	leftY, rightY, upperX, bottomX int //边界
}

func NewRobot(n int) Robot {
	if n == 0 {
		return nil
	}

	fmt.Println("n = ", n)

	return &MatrixRobot{
		n:         n,
		x:         0,
		y:         0,
		direction: "right",
		leftY:     0,
		rightY:    n - 1,
		upperX:    0,
		bottomX:   n - 1,
		currStep:  0,
	}
}

func (r *MatrixRobot) Move() (x int, y int) {
	if r.currStep > r.n*r.n {
		fmt.Println("this robot have already complete moved.")
		return -1, -1
	}

	x = r.x
	y = r.y

	fmt.Println("x:", x, ";y:", y)
	r.currStep += 1

	switch r.direction {
	case "right":
		{
			tmpY := r.y + 1
			if tmpY <= r.rightY {
				r.y += 1
			} else {
				// 边界变化
				r.upperX += 1
				// step 变化
				r.x += 1
				// direction 变化
				r.direction = "down"
			}
		}
	case "down":
		{
			tmpX := r.x + 1
			if tmpX <= r.bottomX {
				r.x = tmpX
			} else {
				// 边界变化
				r.rightY -= 1
				r.y -= 1
				r.direction = "left"
			}
		}
	case "left":
		{
			tmpY := r.y - 1
			if tmpY >= r.leftY {
				r.y = tmpY
			} else {
				r.bottomX -= 1
				r.x -= 1
				r.direction = "upper"
			}
		}
	case "upper":
		{
			tmpX := r.x - 1
			if tmpX >= r.upperX {
				r.x = tmpX
			} else {
				r.leftY += 1
				r.y += 1
				r.direction = "right"
			}
		}
	}

	return
}

func (r *MatrixRobot) Reset() {
	r.x = 0
	r.y = 0
	r.direction = "right"
	r.currStep = 0
	r.leftY = 0
	r.rightY = r.n - 1
	r.upperX = 0
	r.bottomX = r.n - 1

	fmt.Println("matricRobot reset successfully")
}
