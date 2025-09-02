package lc_00669_trim_a_binary_search_tree

import (
	"testing"
)

// isSameTree checks if two binary trees are structurally identical and have the same node values
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Both nodes are nil
	if p == nil && q == nil {
		return true
	}

	// One node is nil, the other is not
	if p == nil || q == nil {
		return false
	}

	// Values are different
	if p.Val != q.Val {
		return false
	}

	// Recursively check left and right subtrees
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestTrimBST(t *testing.T) {
	// Test case 1: Normal case
	// Tree:     1
	//          /
	//         0   2
	//       low=1, high=2
	// Expected:  1
	//
	//              2
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	expected1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	result1 := TrimBST(root1, 1, 2)
	if !isSameTree(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected tree structure does not match.")
	}

	// Test case 2: More complex case
	// Tree:       3
	//            /
	//           0   4
	//
	//             2
	//            /
	//           1
	//       low=1, high=3
	// Expected:   3
	//            /
	//           1
	//
	//             2
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 0,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	expected2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	result2 := TrimBST(root2, 1, 3)
	if !isSameTree(result2, expected2) {
		t.Errorf("Test case 2 failed. Expected tree structure does not match.")
	}

	// Test case 3: All nodes within range
	// Tree:   2
	//        /
	//       1   3
	//     low=0, high=4
	// Expected:  2
	//           /
	//          1   3
	root3 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	result3 := TrimBST(root3, 0, 4)
	if !isSameTree(result3, root3) {
		t.Errorf("Test case 3 failed. Expected tree structure does not match.")
	}

	// Test case 4: Empty tree
	var root4 *TreeNode
	result4 := TrimBST(root4, 1, 2)
	if result4 != nil {
		t.Errorf("Test case 4 failed. Expected nil for empty tree.")
	}
}
