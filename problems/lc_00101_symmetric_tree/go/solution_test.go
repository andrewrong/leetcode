package lc_00101_symmetric_tree

import (
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	// Test case 1: Symmetric tree
	//       1
	//      / \
	//     2   2
	//    / \ / \
	//   3  4 4  3
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	if !isSymmetric(root1) {
		t.Errorf("Expected true for symmetric tree, got false")
	}

	// Test case 2: Not symmetric tree
	//       1
	//      / \
	//     2   2
	//      \   \
	//       3   3
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	if isSymmetric(root2) {
		t.Errorf("Expected false for non-symmetric tree, got true")
	}

	// Test case 3: Empty tree
	var root3 *TreeNode

	if !isSymmetric(root3) {
		t.Errorf("Expected true for empty tree, got false")
	}

	// Test case 4: Single node
	root4 := &TreeNode{
		Val: 1,
	}

	if !isSymmetric(root4) {
		t.Errorf("Expected true for single node tree, got false")
	}

	// Test case 5: [1,2,2,2,null,2]
	//       1
	//      / \
	//     2   2
	//    /   /
	//   2   2
	root5 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}

	if isSymmetric(root5) {
		t.Errorf("Expected false for tree [1,2,2,2,null,2], got true")
	}

	// Test case 6: [5,2,2,4,null,null,1,null,1,null,4,2,null,2,null]
	//               5
	//             /   \
	//            2     2
	//           / \   / \
	//          4  null null 1
	//         / \     / \
	//       null 1  null 4
	//           / \   / \
	//          2 null null 2
	//         / \
	//       null null
	root6 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		},
	}

	if isSymmetric(root6) {
		t.Errorf("Expected false for tree [5,2,2,4,null,null,1,null,1,null,4,2,null,2,null], got true")
	}
}
