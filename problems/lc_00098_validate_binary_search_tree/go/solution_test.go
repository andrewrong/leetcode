package lc_00098_validate_binary_search_tree

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	// Test case 1: Valid BST
	//    2
	//   / \
	//  1   3
	root1 := &TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}
	
	if !isValidBST(root1) {
		t.Errorf("Expected true for valid BST, got false")
	}
	
	// Test case 2: Invalid BST
	//    5
	//   / \
	//  1   4
	//     / \
	//    3   6
	root2 := &TreeNode{Val: 5}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 4}
	root2.Right.Left = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 6}
	
	if isValidBST(root2) {
		t.Errorf("Expected false for invalid BST, got true")
	}
	
	// Test case 3: Single node
	root3 := &TreeNode{Val: 1}
	if !isValidBST(root3) {
		t.Errorf("Expected true for single node, got false")
	}
	
	// Test case 4: Empty tree
	if !isValidBST(nil) {
		t.Errorf("Expected true for empty tree, got false")
	}
	
	// Test case 5: All left children
	// 3
	//  \
	//   2
	//    \
	//     1
	root5 := &TreeNode{Val: 3}
	root5.Right = &TreeNode{Val: 2}
	root5.Right.Right = &TreeNode{Val: 1}
	
	if isValidBST(root5) {
		t.Errorf("Expected false for invalid BST, got true")
	}
	
	// Test case 6: Invalid BST from example [5,4,6,null,null,3,7]
	//    5
	//   / \
	//  4   6
	//     / \
	//    3   7
	root6 := &TreeNode{Val: 5}
	root6.Left = &TreeNode{Val: 4}
	root6.Right = &TreeNode{Val: 6}
	root6.Right.Left = &TreeNode{Val: 3}
	root6.Right.Right = &TreeNode{Val: 7}
	
	if isValidBST(root6) {
		t.Errorf("Expected false for invalid BST [5,4,6,null,null,3,7], got true")
	}
}