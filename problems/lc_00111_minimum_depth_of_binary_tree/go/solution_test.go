package lc_00111_minimum_depth_of_binary_tree

import (
	"testing"
)

func TestMinDepth(t *testing.T) {
	// Test case 1: Empty tree
	var root1 *TreeNode
	expected1 := 0
	if result := minDepth(root1); result != expected1 {
		t.Errorf("Expected %d, got %d for empty tree", expected1, result)
	}

	// Test case 2: Single node
	root2 := &TreeNode{
		Val: 1,
	}
	expected2 := 1
	if result := minDepth(root2); result != expected2 {
		t.Errorf("Expected %d, got %d for single node", expected2, result)
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
	expected3 := 2
	if result := minDepth(root3); result != expected3 {
		t.Errorf("Expected %d, got %d for balanced tree", expected3, result)
	}

	// Test case 4: Right-skewed tree with single child paths
	//   1
	//    \
	//     2
	//      \
	//       3
	//        \
	//         4
	root4 := &TreeNode{
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
	expected4 := 4
	if result := minDepth(root4); result != expected4 {
		t.Errorf("Expected %d, got %d for right-skewed tree", expected4, result)
	}

	// Test case 5: Left-skewed tree
	//       1
	//      /
	//     2
	//    /
	//   3
	root5 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	expected5 := 3
	if result := minDepth(root5); result != expected5 {
		t.Errorf("Expected %d, got %d for left-skewed tree", expected5, result)
	}

	// Test case 6: Tree with one child being null
	//       1
	//      / \
	//     2   null
	root6 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	expected6 := 2
	if result := minDepth(root6); result != expected6 {
		t.Errorf("Expected %d, got %d for tree with one null child", expected6, result)
	}
}