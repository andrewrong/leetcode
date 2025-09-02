package lc_00701_insert_into_a_bst

import (
	"testing"
)

// Helper function to build a tree from slice representation (level order with nulls)
func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(vals) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Helper function to validate that a tree is a valid BST
func isValidBST(root *TreeNode) bool {
	var validate func(*TreeNode, int, int) bool
	validate = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		
		if node.Val <= min || node.Val >= max {
			return false
		}
		
		return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
	}
	
	return validate(root, -1<<31, 1<<31-1)
}

// Helper function to count nodes in a tree
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func Test_insertIntoBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		val  int
	}{
		{
			name: "Insert into empty tree",
			root: nil,
			val:  5,
		},
		{
			name: "Insert as right child",
			root: buildTree([]interface{}{4, 2, 7, 1, 3}),
			val:  5,
		},
		{
			name: "Insert as left child",
			root: buildTree([]interface{}{40, 20, 60, 10, 30, 50, 70}),
			val:  25,
		},
		{
			name: "Insert at deepest level",
			root: buildTree([]interface{}{4, 2, 7, 1, 3, nil, nil, nil, nil, nil, nil}),
			val:  5,
		},
		{
			name: "Insert smaller than all",
			root: buildTree([]interface{}{4, 2, 7, 1, 3}),
			val:  0,
		},
		{
			name: "Insert larger than all",
			root: buildTree([]interface{}{4, 2, 7, 1, 3}),
			val:  8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Count nodes in original tree
			originalNodeCount := countNodes(tt.root)
			
			// Perform insertion
			result := insertIntoBST(tt.root, tt.val)
			
			// Basic validation that the value was inserted
			if result == nil {
				t.Errorf("insertIntoBST() = nil, want a valid tree")
				return
			}
			
			// Check if the inserted value exists in the tree
			found := false
			var check func(*TreeNode)
			check = func(node *TreeNode) {
				if node == nil {
					return
				}
				if node.Val == tt.val {
					found = true
					return
				}
				check(node.Left)
				check(node.Right)
			}
			check(result)
			
			if !found {
				t.Errorf("insertIntoBST() did not insert value %d", tt.val)
			}
			
			// Validate that the result is still a valid BST
			if !isValidBST(result) {
				t.Errorf("insertIntoBST() produced an invalid BST")
			}
			
			// Check that the node count increased by 1
			newNodeCount := countNodes(result)
			if newNodeCount != originalNodeCount+1 {
				t.Errorf("insertIntoBST() node count mismatch: got %d, want %d", newNodeCount, originalNodeCount+1)
			}
		})
	}
}