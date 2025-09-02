package solution

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DeleteNode deletes the node with the given key from the BST
func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = DeleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}

		root.Val = successor.Val
		root.Right = DeleteNode(root.Right, successor.Val)
	}

	return root
}
