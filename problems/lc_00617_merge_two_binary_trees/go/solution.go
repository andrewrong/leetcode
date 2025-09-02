package solution

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MergeTrees merges two binary trees according to the specified rules
func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	node := &TreeNode{}
	val := 0

	var leftNode1 *TreeNode = nil
	var rightNode1 *TreeNode = nil

	if root1 != nil {
		val += root1.Val
		leftNode1 = root1.Left
		rightNode1 = root1.Right
	}

	var leftNode2 *TreeNode = nil
	var rightNode2 *TreeNode = nil
	if root2 != nil {
		val += root2.Val
		leftNode2 = root2.Left
		rightNode2 = root2.Right
	}
	node.Val = val
	node.Left = MergeTrees(leftNode1, leftNode2)
	node.Right = MergeTrees(rightNode1, rightNode2)

	return node
}
