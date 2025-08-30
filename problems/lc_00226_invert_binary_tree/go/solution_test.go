package lc_00226_invert_binary_tree

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	// Test case 1: [4,2,7,1,3,6,9] => [4,7,2,9,6,3,1]
	root1 := &TreeNode{
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
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	
	expected1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	
	result1 := invertTree(root1)
	if !treesEqual(result1, expected1) {
		t.Errorf("invertTree failed for test case 1")
	}
	
	// Test case 2: [2,1,3] => [2,3,1]
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	
	expected2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	
	result2 := invertTree(root2)
	if !treesEqual(result2, expected2) {
		t.Errorf("invertTree failed for test case 2")
	}
	
	// Test case 3: [] => []
	var root3 *TreeNode
	var expected3 *TreeNode
	
	result3 := invertTree(root3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("invertTree failed for test case 3 (empty tree)")
	}
}

// treesEqual checks if two trees are structurally identical
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