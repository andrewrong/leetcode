package lc_00226_invert_binary_tree

import kit "leetcode-go/kit/go"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree inverts a binary tree and returns its root
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := kit.NewQueue[*TreeNode]()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		node, res := queue.Dequeue()
		if !res {
			return root
		}

		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
