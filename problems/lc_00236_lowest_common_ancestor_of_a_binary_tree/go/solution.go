package lc_00236_lowest_common_ancestor_of_a_binary_tree

import "fmt"

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor finds the lowest common ancestor of two nodes in a binary tree.
// p and q are guaranteed to exist in the tree and are distinct.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// Implementation will be added by the user
	// pPath := make([]*TreeNode, 0)
	// qPath := make([]*TreeNode, 0)
	// pOff := false
	// qOff := false

	// GetPath(root, p, q, &pPath, &qPath, &pOff, &qOff)

	// pSize := len(pPath)
	// qSize := len(qPath)

	// size := min(pSize, qSize)
	// i := 0
	// for ; i < size; i += 1 {
	// 	if pPath[i] != qPath[i] {
	// 		break
	// 	}
	// }

	// if i == 0 {
	// 	return nil
	// }
	// return pPath[i-1]

	post := make([]*TreeNode, 0)
	postOrder(root, &post)

	for i, v := range post {
		fmt.Printf("%d: %d\n", i, v.Val)
	}
	return nil
}

func postOrder(node *TreeNode, res *[]*TreeNode) {
	if node == nil {
		return
	}

	postOrder(node.Left, res)
	postOrder(node.Right, res)
	*res = append(*res, node)
}

func GetPath(node, q, p *TreeNode, pPath, qPath *[]*TreeNode, pOff, qOff *bool) {
	if node == nil {
		return
	}

	if q == node && !*qOff {
		*qOff = true
		*qPath = append(*qPath, node)
	}

	if p == node && !*pOff {
		*pOff = true
		*pPath = append(*pPath, node)
	}

	if *pOff && *qOff {
		return
	}

	if !*pOff {
		*pPath = append(*pPath, node)
	}

	if !*qOff {
		*qPath = append(*qPath, node)
	}

	GetPath(node.Left, q, p, pPath, qPath, pOff, qOff)
	if *pOff && *qOff {
		return
	}
	GetPath(node.Right, q, p, pPath, qPath, pOff, qOff)
	if *pOff && *qOff {
		return
	}

	if !*pOff {
		*pPath = (*pPath)[:len(*pPath)-1]
	}
	if !*qOff {
		*qPath = (*qPath)[:len(*qPath)-1]
	}
}
