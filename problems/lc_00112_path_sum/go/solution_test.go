package lc_00112_path_sum

import (
	"testing"
)

func TestHasPathSum(t *testing.T) {
	// Test case 1: Empty tree
	var root1 *TreeNode
	target1 := 0
	expected1 := false
	if result := hasPathSum(root1, target1); result != expected1 {
		t.Errorf("Expected %t, got %t for empty tree", expected1, result)
	}

	// Test case 2: Single node matching target
	root2 := &TreeNode{
		Val: 1,
	}
	target2 := 1
	expected2 := true
	if result := hasPathSum(root2, target2); result != expected2 {
		t.Errorf("Expected %t, got %t for single node matching target", expected2, result)
	}

	// Test case 3: Single node not matching target
	root3 := &TreeNode{
		Val: 1,
	}
	target3 := 2
	expected3 := false
	if result := hasPathSum(root3, target3); result != expected3 {
		t.Errorf("Expected %t, got %t for single node not matching target", expected3, result)
	}

	// Test case 4: Path exists
	//       5
	//      / \
	//     4   8
	//    /   / \
	//   11  13  4
	//  / \       \
	// 7   2       1
	// Path: 5->4->11->2 = 22
	root4 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	target4 := 22
	expected4 := true
	if result := hasPathSum(root4, target4); result != expected4 {
		t.Errorf("Expected %t, got %t for tree with valid path", expected4, result)
	}

	// Test case 5: Path does not exist
	// Same tree as above, but target sum is different
	target5 := 23
	expected5 := false
	if result := hasPathSum(root4, target5); result != expected5 {
		t.Errorf("Expected %t, got %t for tree without valid path", expected5, result)
	}

	// Test case 6: All negative values
	//    -3
	//    /
	//  -9
	//  /
	// -2
	root6 := &TreeNode{
		Val: -3,
		Left: &TreeNode{
			Val: -9,
			Left: &TreeNode{
				Val: -2,
			},
		},
	}
	target6 := -14
	expected6 := true // -3 + (-9) + (-2) = -14
	if result := hasPathSum(root6, target6); result != expected6 {
		t.Errorf("Expected %t, got %t for tree with negative values", expected6, result)
	}

	// Test case 7: Zero values
	//   0
	//  / \
	// 1   0
	root7 := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	target7 := 1
	expected7 := true // 0 -> 1 path
	if result := hasPathSum(root7, target7); result != expected7 {
		t.Errorf("Expected %t, got %t for tree with zero values", expected7, result)
	}
}