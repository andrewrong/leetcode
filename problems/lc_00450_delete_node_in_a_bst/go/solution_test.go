package solution

import (
	"reflect"
	"testing"
)

// Helper function to create a binary tree from a slice representation
// nil values represent missing nodes
func createTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	
	// Create all nodes
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v != nil {
			val := v.(int)
			nodes[i] = &TreeNode{Val: val}
		}
	}
	
	// Connect nodes
	for i := 0; i < len(nodes)/2; i++ {
		if nodes[i] != nil {
			leftIndex := 2*i + 1
			rightIndex := 2*i + 2
			
			if leftIndex < len(nodes) {
				nodes[i].Left = nodes[leftIndex]
			}
			if rightIndex < len(nodes) {
				nodes[i].Right = nodes[rightIndex]
			}
		}
	}
	
	return nodes[0]
}

// Helper function to convert a binary tree to a slice representation (level-order traversal)
func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}
	
	result := []interface{}{}
	queue := []*TreeNode{root}
	
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		
		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}
	
	// Trim trailing nil values
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}
	
	return result
}

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		key      int
		expected []interface{}
	}{
		{
			name:     "Delete leaf node",
			input:    []interface{}{5, 3, 6, 2, 4, nil, 7},
			key:      2,
			expected: []interface{}{5, 3, 6, nil, 4, nil, 7},
		},
		{
			name:     "Delete node with one child",
			input:    []interface{}{5, 3, 6, 2, 4, nil, 7},
			key:      3,
			expected: []interface{}{5, 4, 6, 2, nil, nil, 7},
		},
		{
			name:     "Delete node with two children",
			input:    []interface{}{5, 3, 6, 2, 4, nil, 7},
			key:      5,
			expected: []interface{}{6, 3, 7, 2, 4}, // One possible result
		},
		{
			name:     "Delete non-existent node",
			input:    []interface{}{5, 3, 6, 2, 4, nil, 7},
			key:      10,
			expected: []interface{}{5, 3, 6, 2, 4, nil, 7},
		},
		{
			name:     "Delete root node from single node tree",
			input:    []interface{}{5},
			key:      5,
			expected: []interface{}{},
		},
		{
			name:     "Delete from empty tree",
			input:    []interface{}{},
			key:      5,
			expected: []interface{}{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := createTree(test.input)
			result := DeleteNode(root, test.key)
			resultSlice := treeToSlice(result)
			
			// For the case of deleting node with two children, there can be multiple valid results
			// We'll check if the result is a valid BST with the correct elements
			if test.name == "Delete node with two children" {
				// Special validation for this case
				// Just check that the result is not nil and the deleted key is not present
				if result != nil && !containsValue(result, test.key) {
					// This is a simplified check - in a real test we might want more thorough validation
					return
				}
				t.Errorf("DeleteNode() = %v, expected a valid BST without key %d", resultSlice, test.key)
			} else {
				if !reflect.DeepEqual(resultSlice, test.expected) {
					t.Errorf("DeleteNode() = %v, expected %v", resultSlice, test.expected)
				}
			}
		})
	}
}

// Helper function to check if a value exists in the tree
func containsValue(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if root.Val == val {
		return true
	}
	return containsValue(root.Left, val) || containsValue(root.Right, val)
}