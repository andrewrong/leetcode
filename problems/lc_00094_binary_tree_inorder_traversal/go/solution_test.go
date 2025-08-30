package lc_00094_binary_tree_inorder_traversal

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	// Test case 1: Empty tree
	root1 := (*TreeNode)(nil)
	expected1 := []int{}
	result1 := inorderTraversal(root1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected1, result1)
	}

	// Test case 2: Single node
	root2 := &TreeNode{Val: 1}
	expected2 := []int{1}
	result2 := inorderTraversal(root2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, got %v", expected2, result2)
	}

	// Test case 3: Complete binary tree
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	root3 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	// 中序遍历结果应该是 [4,2,5,1,3]
	expected3 := []int{4, 2, 5, 1, 3}
	result3 := inorderTraversal(root3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, got %v", expected3, result3)
	}

	// Test case 4: Left-skewed tree
	//   1
	//  /
	// 2
	///
	//3
	root4 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	expected4 := []int{3, 2, 1}
	result4 := inorderTraversal(root4)
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Test case 4 failed: expected %v, got %v", expected4, result4)
	}

	// Test case 5: Right-skewed tree
	// 1
	//  \
	//   2
	//    \
	//     3
	root5 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	expected5 := []int{1, 2, 3}
	result5 := inorderTraversal(root5)
	if !reflect.DeepEqual(result5, expected5) {
		t.Errorf("Test case 5 failed: expected %v, got %v", expected5, result5)
	}
}