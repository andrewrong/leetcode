package solution

import (
	"testing"
)

func TestMergeTrees(t *testing.T) {
	// Test case 1: Example from problem description
	// root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
	// Expected output: [3,4,5,5,4,null,7]
	
	// Creating root1 tree
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	
	// Creating root2 tree
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	
	// Expected merged tree
	expected := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	
	result := MergeTrees(root1, root2)
	if !treesEqual(result, expected) {
		t.Errorf("MergeTrees failed for test case 1")
	}
	
	// Test case 2: One tree is nil
	tree1 := &TreeNode{Val: 1}
	result2 := MergeTrees(tree1, nil)
	if !treesEqual(result2, tree1) {
		t.Errorf("MergeTrees failed for test case 2")
	}
	
	// Test case 3: Both trees are nil
	result3 := MergeTrees(nil, nil)
	if result3 != nil {
		t.Errorf("MergeTrees failed for test case 3")
	}
}

// treesEqual checks if two binary trees are structurally identical
func treesEqual(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && treesEqual(t1.Left, t2.Left) && treesEqual(t1.Right, t2.Right)
}