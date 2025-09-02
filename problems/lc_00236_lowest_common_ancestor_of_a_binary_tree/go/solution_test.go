package lc_00236_lowest_common_ancestor_of_a_binary_tree

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	// Test case 1: Basic case with a simple tree
	// Tree structure:
	//       3
	//      / \
	//     5   1
	//    / \ / \
	//   6  2 0  8
	//     / \
	//    7   4
	//
	// LCA(5, 1) = 3
	// LCA(5, 4) = 5
	// LCA(6, 2) = 5
	
	// Create tree nodes
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	
	// Define test cases
	tests := []struct {
		name     string
		root     *TreeNode
		p        *TreeNode
		q        *TreeNode
		expected *TreeNode
	}{
		{
			name:     "LCA of 5 and 1",
			root:     root,
			p:        root.Left,   // Node with value 5
			q:        root.Right,  // Node with value 1
			expected: root,        // Expected LCA is root (value 3)
		},
		{
			name:     "LCA of 5 and 4",
			root:     root,
			p:        root.Left,                    // Node with value 5
			q:        root.Left.Right.Right,        // Node with value 4
			expected: root.Left,                    // Expected LCA is node with value 5
		},
		{
			name:     "LCA of 6 and 2",
			root:     root,
			p:        root.Left.Left,               // Node with value 6
			q:        root.Left.Right,              // Node with value 2
			expected: root.Left,                    // Expected LCA is node with value 5
		},
	}
	
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lowestCommonAncestor(tt.root, tt.p, tt.q)
			if result != tt.expected {
				t.Errorf("lowestCommonAncestor() = %v, expected %v", result, tt.expected)
			}
		})
	}
}