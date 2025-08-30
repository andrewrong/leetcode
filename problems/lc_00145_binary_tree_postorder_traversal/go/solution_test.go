package solution

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	// Helper function to create a tree node
	createNode := func(val int) *TreeNode {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	// Test case 1: [1,null,2,3] => [3,2,1]
	root1 := createNode(1)
	root1.Right = createNode(2)
	root1.Right.Left = createNode(3)
	
	// Test case 2: [] => []
	var root2 *TreeNode
	
	// Test case 3: [1] => [1]
	root3 := createNode(1)

	tests := []struct {
		root     *TreeNode
		expected []int
	}{
		{root1, []int{3, 2, 1}},
		{root2, []int{}},
		{root3, []int{1}},
	}

	for _, test := range tests {
		result := PostorderTraversal(test.root)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("PostorderTraversal() = %v; expected %v", result, test.expected)
		}
	}
}