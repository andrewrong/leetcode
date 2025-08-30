package kit

import (
	"reflect"
	"testing"
)

func TestTreeTraversals(t *testing.T) {
	// Create a tree:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5

	root := &treeNode[int]{
		Val: 1,
		Left: &treeNode[int]{
			Val: 2,
			Left: &treeNode[int]{
				Val: 4,
			},
			Right: &treeNode[int]{
				Val: 5,
			},
		},
		Right: &treeNode[int]{
			Val: 3,
		},
	}

	tree := Tree[int](root)

	t.Run("PreOrderTraversal", func(t *testing.T) {
		var result []int
		tree.PreOrderTraversal(&result)
		expected := []int{1, 2, 4, 5, 3} // Root, left subtree, right subtree
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("PreOrderTraversal() = %v, want %v", result, expected)
		}
	})

	t.Run("PreOrderTraversalByIterate", func(t *testing.T) {
		result := tree.PreOrderTraversalByIterate()
		expected := []int{1, 2, 4, 5, 3} // Root, left subtree, right subtree
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("PreOrderTraversalByIterate() = %v, want %v", result, expected)
		}
	})

	t.Run("InOrderTraversal", func(t *testing.T) {
		var result []int
		tree.InOrderTraversal(&result)
		expected := []int{4, 2, 5, 1, 3} // Left subtree, root, right subtree
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("InOrderTraversal() = %v, want %v", result, expected)
		}
	})

	t.Run("InOrderTraversalByIterate", func(t *testing.T) {
		result := tree.InOrderTraversalByIterate()
		expected := []int{4, 2, 5, 1, 3} // Left subtree, root, right subtree
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("InOrderTraversalByIterate() = %v, want %v", result, expected)
		}
	})

	t.Run("PostOrderTraversal", func(t *testing.T) {
		var result []int
		tree.PostOrderTraversal(&result)
		expected := []int{4, 5, 2, 3, 1} // Left subtree, right subtree, root
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("PostOrderTraversal() = %v, want %v", result, expected)
		}
	})

	t.Run("PostOrderTraversalByIterate", func(t *testing.T) {
		result := tree.PostOrderTraversalByIterate()
		expected := []int{4, 5, 2, 3, 1} // Left subtree, right subtree, root
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("PostOrderTraversalByIterate() = %v, want %v", result, expected)
		}
	})

	t.Run("LevelTraversal", func(t *testing.T) {
		var result []int
		tree.LevelTraversal(&result)
		expected := []int{1, 2, 3, 4, 5} // Level by level from top to bottom
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("LevelTraversal() = %v, want %v", result, expected)
		}
	})

	// Test with empty tree
	t.Run("EmptyTree", func(t *testing.T) {
		var emptyTree Tree[int] = &treeNode[int]{}
		// Initialize with zero value
		emptyTree.(*treeNode[int]).Val = 0
		emptyTree.(*treeNode[int]).Left = nil
		emptyTree.(*treeNode[int]).Right = nil

		var result []int
		emptyTree.PreOrderTraversal(&result)
		if len(result) != 1 || result[0] != 0 {
			t.Errorf("EmptyTree PreOrderTraversal() = %v, want [0]", result)
		}

		// Test iterative version with empty tree
		iterResult := emptyTree.PreOrderTraversalByIterate()
		if len(iterResult) != 1 || result[0] != 0 {
			t.Errorf("EmptyTree PreOrderTraversalByIterate() = %v, want []", iterResult)
		}
	})

	// Test with string values
	t.Run("StringTree", func(t *testing.T) {
		stringRoot := &treeNode[string]{
			Val: "root",
			Left: &treeNode[string]{
				Val: "left",
			},
			Right: &treeNode[string]{
				Val: "right",
			},
		}

		stringTree := Tree[string](stringRoot)
		var result []string
		stringTree.InOrderTraversal(&result)
		expected := []string{"left", "root", "right"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("StringTree InOrderTraversal() = %v, want %v", result, expected)
		}

		// Test iterative version with string values
		iterResult := stringTree.InOrderTraversalByIterate()
		if !reflect.DeepEqual(iterResult, expected) {
			t.Errorf("StringTree InOrderTraversalByIterate() = %v, want %v", iterResult, expected)
		}
	})

	// Test with single node
	t.Run("SingleNode", func(t *testing.T) {
		singleNode := &treeNode[int]{
			Val: 42,
		}
		singleTree := Tree[int](singleNode)

		// Test recursive version
		var result []int
		singleTree.InOrderTraversal(&result)
		expected := []int{42}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SingleNode InOrderTraversal() = %v, want %v", result, expected)
		}

		// Test iterative version
		iterResult := singleTree.InOrderTraversalByIterate()
		if !reflect.DeepEqual(iterResult, expected) {
			t.Errorf("SingleNode InOrderTraversalByIterate() = %v, want %v", iterResult, expected)
		}
	})
}
