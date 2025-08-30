package lc_00257_binary_tree_paths

import (
	"strconv"
	"strings"
)

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// binaryTreePaths returns all root-to-leaf paths in a binary tree.
func binaryTreePaths(root *TreeNode) []string {
	nodes := make([]string, 0)
	paths := make([]string, 0)

	if root == nil {
		return paths
	}

	getPath(root, &nodes, &paths)

	return paths
}

func getPath(node *TreeNode, nodes *[]string, path *[]string) {
	if node == nil {
		return
	}

	// 单层处理, 回溯
	*nodes = append(*nodes, strconv.FormatInt(int64(node.Val), 10))
	size := len(*nodes)
	// node == left node
	if node.Left == nil && node.Right == nil {
		*path = append(*path, strings.Join(*nodes, "->"))
		return
	}
	getPath(node.Left, nodes, path)
	*nodes = (*nodes)[:size]
	getPath(node.Right, nodes, path)
}
