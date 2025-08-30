package lc_00111_minimum_depth_of_binary_tree

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// minDepth 计算二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	minLevel := 1
	if left != 0 && right != 0 {
		minLevel += min(left, right)
	} else if left == 0 {
		minLevel += right
	} else if right == 0 {
		minLevel += left
	}
	return minLevel
}
