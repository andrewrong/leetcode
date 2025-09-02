package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SortedArrayToBST converts a sorted array to a height-balanced binary search tree.
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	// root 节点
	midNode := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}

	if mid == 0 {
		return midNode
	}

	midNode.Left = SortedArrayToBST(nums[:mid])
	midNode.Right = SortedArrayToBST(nums[mid+1:])

	return nil
}
