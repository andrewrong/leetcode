package lc_00701_insert_into_a_bst

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// insertIntoBST inserts a value into the BST and returns the root node.
// You need to implement this function.
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	if root.Val > val {
		n := insertIntoBST(root.Left, val)
		if n != nil {
			root.Left = n
			return root
		}
	} else if root.Val < val {
		n := insertIntoBST(root.Right, val)
		if n != nil {
			root.Right = n
			return root
		}
	}

	return root
}
