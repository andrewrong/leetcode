package lc_00106_construct_binary_tree_from_inorder_and_postorder_traversal

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree 从中序遍历和后序遍历构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	return BuildTree(inorder, postorder, inorderMap)
}

func BuildTree(inorder []int, postorder []int, inorderMap map[int]int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	rootIdx, _ := inorderMap[rootVal]

	startIdx, _ := inorderMap[inorder[0]]

	node := &TreeNode{
		Val: rootVal,
	}

	leftSize := len(inorder[:rootIdx-startIdx])

	left := BuildTree(inorder[:rootIdx-startIdx], postorder[:leftSize], inorderMap)
	right := BuildTree(inorder[rootIdx-startIdx+1:], postorder[leftSize:len(postorder)-1], inorderMap)

	node.Left = left
	node.Right = right

	return node
}
