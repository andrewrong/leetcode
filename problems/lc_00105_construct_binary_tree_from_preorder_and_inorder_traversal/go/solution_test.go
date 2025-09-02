package lc_00105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	// Test case 1: Normal case
	preorder1 := []int{3, 9, 20, 15, 7}
	inorder1 := []int{9, 3, 15, 20, 7}
	expected1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	result1 := buildTree(preorder1, inorder1)
	if !treesEqual(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected tree structure does not match result.")
	}

	// Test case 2: Single node
	preorder2 := []int{1}
	inorder2 := []int{1}
	expected2 := &TreeNode{
		Val: 1,
	}

	result2 := buildTree(preorder2, inorder2)
	if !treesEqual(result2, expected2) {
		t.Errorf("Test case 2 failed. Expected tree structure does not match result.")
	}

	// Test case 3: Empty tree
	preorder3 := []int{}
	inorder3 := []int{}
	var expected3 *TreeNode

	result3 := buildTree(preorder3, inorder3)
	if !treesEqual(result3, expected3) {
		t.Errorf("Test case 3 failed. Expected tree structure does not match result.")
	}

	// Test case 4: Left-skewed tree
	preorder4 := []int{1, 2, 3, 4}
	inorder4 := []int{4, 3, 2, 1}
	expected4 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
			},
		},
	}

	result4 := buildTree(preorder4, inorder4)
	if !treesEqual(result4, expected4) {
		t.Errorf("Test case 4 failed. Expected tree structure does not match result.")
	}

	// Test case 5: Right-skewed tree
	preorder5 := []int{1, 2, 3, 4}
	inorder5 := []int{1, 2, 3, 4}
	expected5 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}

	result5 := buildTree(preorder5, inorder5)
	if !treesEqual(result5, expected5) {
		t.Errorf("Test case 5 failed. Expected tree structure does not match result.")
	}
}

// treesEqual 比较两个二叉树是否相等
func treesEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return treesEqual(t1.Left, t2.Left) && treesEqual(t1.Right, t2.Right)
}