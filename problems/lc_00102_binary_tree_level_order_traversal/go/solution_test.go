package lc_00102_binary_tree_level_order_traversal

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	// Test case 1: Empty tree
	var root1 *TreeNode
	expected1 := [][]int{}
	result1 := levelOrder(root1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected1, result1)
	}

	// Test case 2: Single node
	root2 := &TreeNode{Val: 1}
	expected2 := [][]int{{1}}
	result2 := levelOrder(root2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, got %v", expected2, result2)
	}

	// Test case 3: Complete binary tree
	//       3
	//      / \
	//     9   20
	//        /  \
	//       15   7
	root3 := &TreeNode{
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
	expected3 := [][]int{{3}, {9, 20}, {15, 7}}
	result3 := levelOrder(root3)
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
	expected4 := [][]int{{1}, {2}, {3}}
	result4 := levelOrder(root4)
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Test case 4 failed: expected %v, got %v", expected4, result4)
	}

	// Test case 5: Right-skewed tree
	// 1
	//  \
	//   2
	//    \
	//     3
	//      \
	//       4
	root5 := &TreeNode{
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
	expected5 := [][]int{{1}, {2}, {3}, {4}}
	result5 := levelOrder(root5)
	if !reflect.DeepEqual(result5, expected5) {
		t.Errorf("Test case 5 failed: expected %v, got %v", expected5, result5)
	}

	// Test case 6: Full binary tree with more levels
	//         1
	//       /   \
	//      2     3
	//     / \   / \
	//    4   5 6   7
	root6 := &TreeNode{
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
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	expected6 := [][]int{{1}, {2, 3}, {4, 5, 6, 7}}
	result6 := levelOrder(root6)
	if !reflect.DeepEqual(result6, expected6) {
		t.Errorf("Test case 6 failed: expected %v, got %v", expected6, result6)
	}
}
