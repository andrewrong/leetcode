package solution

import (
	"testing"
)

func TestFindMode(t *testing.T) {
	// Helper function to create a tree node
	newTreeNode := func(val int, left, right *TreeNode) *TreeNode {
		return &TreeNode{
			Val:   val,
			Left:  left,
			Right: right,
		}
	}

	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     newTreeNode(1, nil, nil),
			expected: []int{1},
		},
		{
			name: "Example case - [1,null,2,2]",
			root: newTreeNode(1,
				nil,
				newTreeNode(2,
					newTreeNode(2, nil, nil),
					nil)),
			expected: []int{2},
		},
		{
			name: "Multiple modes",
			root: newTreeNode(1,
				newTreeNode(1, nil, nil),
				newTreeNode(2,
					newTreeNode(2, nil, nil),
					newTreeNode(3, nil, nil))),
			expected: []int{1, 2}, // Both 1 and 2 appear twice
		},
		{
			name: "All elements appear once",
			root: newTreeNode(2,
				newTreeNode(1, nil, nil),
				newTreeNode(3, nil, nil)),
			expected: []int{1, 2, 3}, // All appear once, so all are modes
		},
		{
			name: "Complex case with multiple occurrences",
			root: newTreeNode(5,
				newTreeNode(2,
					newTreeNode(1, nil, nil),
					newTreeNode(3, nil, nil)),
				newTreeNode(7,
					newTreeNode(6, nil, nil),
					newTreeNode(7, nil, nil))),
			// Values: 1(1), 2(1), 3(1), 5(1), 6(1), 7(2)
			// Mode is 7 with 2 occurrences
			expected: []int{7},
		},
		{
			name: "BST with repeated values",
			root: newTreeNode(1,
				nil,
				newTreeNode(2,
					newTreeNode(2, nil, nil),
					newTreeNode(2, nil, nil))),
			// Values: 1(1), 2(3)
			// Mode is 2 with 3 occurrences
			expected: []int{2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindMode(test.root)

			// Sort both slices for comparison since order doesn't matter
			// We can compare lengths and check if all elements are present
			if len(result) != len(test.expected) {
				t.Errorf("FindMode() = %v; expected %v. Length mismatch: got %d, want %d",
					result, test.expected, len(result), len(test.expected))
				return
			}

			// Create maps for easier comparison
			resultMap := make(map[int]bool)
			expectedMap := make(map[int]bool)

			for _, v := range result {
				resultMap[v] = true
			}

			for _, v := range test.expected {
				expectedMap[v] = true
			}

			// Check if all elements in result are in expected
			for k := range resultMap {
				if !expectedMap[k] {
					t.Errorf("FindMode() = %v; expected %v. Found unexpected element: %d",
						result, test.expected, k)
					return
				}
			}

			// Check if all elements in expected are in result
			for k := range expectedMap {
				if !resultMap[k] {
					t.Errorf("FindMode() = %v; expected %v. Missing element: %d",
						result, test.expected, k)
					return
				}
			}
		})
	}
}
