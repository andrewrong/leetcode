package lc_00700_search_in_a_bst

import (
	"testing"
)

func TestSearchBST(t *testing.T) {
	// Test case 1: Empty tree
	var root1 *TreeNode
	val1 := 5
	expected1 := (*TreeNode)(nil)
	result1 := searchBST(root1, val1)
	if result1 != expected1 {
		t.Errorf("Expected nil for empty tree, got %v", result1)
	}

	// Test case 2: Value exists in the tree
	// Tree structure:
	//       4
	//      / \
	//     2   7
	//    / \
	//   1   3
	root2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}
	
	// Search for existing value 2
	val2 := 2
	expected2 := root2.Left
	result2 := searchBST(root2, val2)
	if result2 != expected2 {
		t.Errorf("Expected node with value %d, got %v", val2, result2)
	}
	
	// Search for existing value 7
	val3 := 7
	expected3 := root2.Right
	result3 := searchBST(root2, val3)
	if result3 != expected3 {
		t.Errorf("Expected node with value %d, got %v", val3, result3)
	}
	
	// Test case 3: Value does not exist in the tree
	val4 := 5
	expected4 := (*TreeNode)(nil)
	result4 := searchBST(root2, val4)
	if result4 != expected4 {
		t.Errorf("Expected nil for non-existing value, got %v", result4)
	}
	
	// Test case 4: Value is at the root
	val5 := 4
	expected5 := root2
	result5 := searchBST(root2, val5)
	if result5 != expected5 {
		t.Errorf("Expected root node for value at root, got %v", result5)
	}
}