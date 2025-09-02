package lc_00098_validate_binary_search_tree

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBST checks if a binary tree is a valid binary search tree
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	res := make([]int, 0)
	inorder(root, &res)

	for i := 0; i < len(res)-1; i += 1 {
		if res[i] < res[i+1] {
			continue
		}
		return false
	}

	return true
}

func inorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, result)
	*result = append(*result, root.Val)
	inorder(root.Right, result)
}
