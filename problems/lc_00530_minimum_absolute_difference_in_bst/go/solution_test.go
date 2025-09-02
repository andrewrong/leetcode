package solution

import (
	"testing"
)

func TestGetMinimumDifference(t *testing.T) {
	// Test case 1: [4,2,6,1,3]
	//       4
	//      / \
	//     2   6
	//    / \
	//   1   3
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
			Val: 6,
		},
	}

	// Test case 2: [1,0,48,null,null,12,49]
	//         1
	//       /   \
	//      0     48
	//          /   \
	//         12   49
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 48,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 49,
			},
		},
	}

	// Test case 3: Two nodes only
	//   1
	//    \
	//     3
	root3 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
		},
	}

	// Test case 4: [236,104,701,null,227,null,911]
	//       236
	//      /   \
	//    104   701
	//      \     \
	//      227   911
	root4 := &TreeNode{
		Val: 236,
		Left: &TreeNode{
			Val: 104,
			Right: &TreeNode{
				Val: 227,
			},
		},
		Right: &TreeNode{
			Val: 701,
			Right: &TreeNode{
				Val: 911,
			},
		},
	}

	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{"Test case 1", root1, 1},
		{"Test case 2", root2, 1},
		{"Two nodes", root3, 2},
		{"Test case 4", root4, 9}, // |236-227| = 9
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetMinimumDifference(test.root)
			if result != test.expected {
				t.Errorf("GetMinimumDifference() = %d; expected %d", result, test.expected)
			}
		})
	}
}
