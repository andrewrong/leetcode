package lc_00101_symmetric_tree

// TreeNode 定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 核心就是判断根节点的左子树和右子树是否互为镜像
	return isMirror(root.Left, root.Right)
}

// isMirror 是一个辅助函数，用于判断两个节点 t1 和 t2 是否互为镜像
func isMirror(t1 *TreeNode, t2 *TreeNode) bool {
	// 基本情况 (Base Case):
	// 1. 如果两个节点都是 nil，那么它们是镜像对称的
	if t1 == nil && t2 == nil {
		return true
	}
	// 2. 如果只有一个是 nil，或者它们的值不相等，那么肯定不对称
	if t1 == nil || t2 == nil || t1.Val != t2.Val {
		return false
	}

	// 递归步骤 (Recursive Step):
	// 必须同时满足两个条件：
	// 1. t1的左子树 和 t2的右子树 互为镜像
	// 2. t1的右子树 和 t2的左子树 互为镜像
	return isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}
