package solution

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PostorderTraversal returns the postorder traversal of a binary tree's nodes' values.
func PostorderTraversal(root *TreeNode) []int {
	// Implementation will be added later
	if root == nil {
		return []int{}
	}
	res := []int{}
	postOrderTraversal(root, &res)

	return res
}

func postOrderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	postOrderTraversal(root.Left, res)
	postOrderTraversal(root.Right, res)
	*res = append(*res, root.Val)
}
