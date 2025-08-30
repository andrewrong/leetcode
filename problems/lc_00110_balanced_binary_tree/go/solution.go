package lc_00110_balanced_binary_tree

import "math"

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced 判断二叉树是否是平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return GetHeight(root) != -1
}

func GetHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := GetHeight(node.Left)
	right := GetHeight(node.Right)

	if left == -1 {
		return -1
	}

	if right == -1 {
		return -1
	}

	if math.Abs(float64(left)-float64(right)) > 1 {
		return -1
	}

	return max(left, right) + 1
}
