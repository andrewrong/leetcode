package lc_00257_binary_tree_paths

import (
	"reflect"
	"sort"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	// Test case 1: Normal binary tree
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	expected1 := []string{"1->2->5", "1->3"}
	result1 := binaryTreePaths(root1)
	sort.Strings(result1)
	sort.Strings(expected1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected1, result1)
	}

	// Test case 2: Single node tree
	root2 := &TreeNode{Val: 1}
	expected2 := []string{"1"}
	result2 := binaryTreePaths(root2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, got %v", expected2, result2)
	}

	// Test case 3: Empty tree
	var root3 *TreeNode
	expected3 := []string{}
	result3 := binaryTreePaths(root3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, got %v", expected3, result3)
	}
}