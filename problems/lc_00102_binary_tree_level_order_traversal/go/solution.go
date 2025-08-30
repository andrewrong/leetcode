package lc_00102_binary_tree_level_order_traversal

import kit "leetcode-go/kit/go"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := kit.NewQueue[*TreeNode]()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		size := queue.Size()
		tmp := make([]int, 0)

		for i := 0; i < size; i += 1 {
			node, res := queue.Dequeue()
			if !res {
				return result
			}

			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
		result = append(result, tmp)
	}

	return result
}
