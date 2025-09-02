package lc_00669_trim_a_binary_search_tree

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TrimBST trims the binary search tree such that all its elements lies in [low, high]
func TrimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > high {
		return TrimBST(root.Left, low, high)
	} else if root.Val < low {
		return TrimBST(root.Right, low, high)
	} else {
		root.Left = TrimBST(root.Left, low, root.Val)
		root.Right = TrimBST(root.Right, root.Val, high)
	}

	return root
}
