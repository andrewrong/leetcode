package lc_00110_balanced_binary_tree

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	// Test case 1: Empty tree
	var root1 *TreeNode
	expected1 := true
	if result := isBalanced(root1); result != expected1 {
		t.Errorf("Expected %t, got %t for empty tree", expected1, result)
	}

	// Test case 2: Single node
	root2 := &TreeNode{
		Val: 1,
	}
	expected2 := true
	if result := isBalanced(root2); result != expected2 {
		t.Errorf("Expected %t, got %t for single node", expected2, result)
	}

	// Test case 3: Balanced tree
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
	expected3 := true
	if result := isBalanced(root3); result != expected3 {
		t.Errorf("Expected %t, got %t for balanced tree", expected3, result)
	}

	// Test case 4: Unbalanced tree (right subtree deeper)
	//       1
	//      / \
	//     2   2
	//    / \
	//   3   3
	//  / \
	// 4   4
	node4Left := &TreeNode{Val: 4}
	node4Right := &TreeNode{Val: 4}
	node3Left := &TreeNode{
		Val:   3,
		Left:  node4Left,
		Right: node4Right,
	}
	node3Right := &TreeNode{Val: 3}
	node2Left := &TreeNode{
		Val:   2,
		Left:  node3Left,
		Right: node3Right,
	}
	node2Right := &TreeNode{Val: 2}
	root4 := &TreeNode{
		Val:   1,
		Left:  node2Left,
		Right: node2Right,
	}
	expected4 := false
	if result := isBalanced(root4); result != expected4 {
		t.Errorf("Expected %t, got %t for unbalanced tree", expected4, result)
	}

	// Test case 5: Another balanced tree
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	root5 := &TreeNode{
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
	expected5 := true
	if result := isBalanced(root5); result != expected5 {
		t.Errorf("Expected %t, got %t for another balanced tree", expected5, result)
	}

	// Test case 6: Unbalanced tree from problem example
	//       1
	//      / \
	//     2   2
	//    /   /
	//   3   3
	//  /   /
	// 4   4
	// This represents [1,2,2,3,null,null,3,4,null,null,4]
	root6 := &TreeNode{
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
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	expected6 := false
	if result := isBalanced(root6); result != expected6 {
		t.Errorf("Expected %t, got %t for unbalanced tree from problem example", expected6, result)
	}
}
