package solution

// TreeNode 定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var base, count, maxCount int
	var result []int

	// update 函数负责在遇到新数字或遍历结束时，更新 result 列表
	update := func() {
		if count > maxCount {
			maxCount = count
			result = []int{base} // 发现了新的最高频率，清空结果，只放当前的 base
		} else if count == maxCount {
			result = append(result, base) // 频率与最高相同，追加到结果中
		}
	}

	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrder(node.Left)

		// --- 核心处理逻辑 ---
		if node.Val == base {
			count++
		} else {
			update() // 遇到了新数字，先结算上一个数字
			base = node.Val
			count = 1
		}
		// --- 核心处理逻辑结束 ---

		inOrder(node.Right)
	}

	inOrder(root)
	update() // 别忘了对遍历的最后一个数字序列进行结算！

	return result
}
