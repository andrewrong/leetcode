package lc_00104_maximum_depth_of_binary_tree

import "math"

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth 计算二叉树的最大深度
func maxDepth(root *TreeNode) int {
	// 终止条件
	if root == nil {
		return 0
	}
	// 单层处理
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}
