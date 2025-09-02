package solution

import "math"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// GetMinimumDifference returns the minimum absolute difference between values of any two different nodes in a BST
func GetMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := make([]int, 0)
	inorder(root, &res)

	minV := math.MaxInt
	for i := 0; i < len(res)-1; i += 1 {
		v := math.Abs(float64(res[i] - res[i+1]))
		if int(v) < minV {
			minV = int(v)
		}
	}

	return minV
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}
