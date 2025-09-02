package lc_00105_construct_binary_tree_from_preorder_and_inorder_traversal

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree 从前序遍历和中序遍历构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}

	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	return BuildTree(preorder, inorder, inorderMap)
}

func BuildTree(preorder []int, inorder []int, inorderMap map[int]int) *TreeNode {
	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rootIdx, _ := inorderMap[rootVal]

	startIdx, _ := inorderMap[inorder[0]]

	node := &TreeNode{
		Val: rootVal,
	}

	leftSize := len(inorder[:rootIdx-startIdx])

	left := BuildTree(preorder[1:leftSize+1], inorder[:rootIdx-startIdx], inorderMap)
	right := BuildTree(preorder[leftSize+1:], inorder[rootIdx-startIdx+1:], inorderMap)

	node.Left = left
	node.Right = right

	return node
}
