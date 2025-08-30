package lc_00112_path_sum

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// hasPathSum 判断是否存在从根节点到叶子节点的路径，使得路径上所有节点值之和等于目标和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}

	if root.Left != nil {
		res := hasPathSum(root.Left, targetSum)
		if res {
			return true
		}
	}
	if root.Right != nil {
		res := hasPathSum(root.Right, targetSum)
		if res {
			return true
		}
	}

	return false
}
