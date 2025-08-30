package lc_00094_binary_tree_inorder_traversal

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderTraversal returns the inorder traversal of a binary tree's nodes' values
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	inOrderTraversal(root, &res)
	return res
}

// inOrderTraversal performs the actual recursive inorder traversal
// Inorder traversal: Left -> Root -> Right
func inOrderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	// Traverse left subtree
	inOrderTraversal(root.Left, res)

	// Visit root node
	*res = append(*res, root.Val)

	// Traverse right subtree
	inOrderTraversal(root.Right, res)
}
