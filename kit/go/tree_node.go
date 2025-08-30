package kit

type Tree[T any] interface {
	PreOrderTraversal(result *[]T)
	PreOrderTraversalByIterate() []T
	InOrderTraversal(result *[]T)
	InOrderTraversalByIterate() []T
	PostOrderTraversal(result *[]T)
	PostOrderTraversalByIterate() []T
	LevelTraversal(result *[]T)
}

// TreeNode is the definition for a binary tree node.
type treeNode[T any] struct {
	Val   T
	Left  *treeNode[T]
	Right *treeNode[T]
}

func NewTree[T any]() Tree[T] {
	var zero T
	root := &treeNode[T]{
		Val:   zero,
		Left:  nil,
		Right: nil,
	}
	return root
}

func (t *treeNode[T]) PreOrderTraversal(result *[]T) {
	if t == nil {
		return
	}
	*result = append(*result, t.Val)
	t.Left.PreOrderTraversal(result)
	t.Right.PreOrderTraversal(result)

}

func (t *treeNode[T]) PreOrderTraversalByIterate() []T {
	result := make([]T, 0)
	if t == nil {
		return result
	}

	st := NewStack[*treeNode[T]]()
	st.Push(t)

	for st.Size() != 0 {
		node, res := st.Pop()

		if !res {
			return result
		}

		if node == nil {
			node, res = st.Pop()
			if !res {
				return result
			}
			result = append(result, node.Val)
		} else {
			if node.Right != nil {
				st.Push(node.Right)
			}
			if node.Left != nil {
				st.Push(node.Left)
			}

			st.Push(node)
			st.Push(nil)
		}
	}

	return result
}

func (t *treeNode[T]) InOrderTraversal(result *[]T) {
	if t == nil {
		return
	}
	t.Left.InOrderTraversal(result)
	*result = append(*result, t.Val)
	t.Right.InOrderTraversal(result)
}

func (t *treeNode[T]) InOrderTraversalByIterate() []T {
	result := make([]T, 0)
	if t == nil {
		return result
	}

	st := NewStack[*treeNode[T]]()
	st.Push(t)

	for st.Size() != 0 {
		node, res := st.Pop()
		if !res {
			return result
		}

		if node == nil {
			tmp, res := st.Pop()
			if !res {
				return result
			}
			result = append(result, tmp.Val)
		} else {
			if node.Right != nil {
				st.Push(node.Right)
			}

			st.Push(node)
			st.Push(nil)

			if node.Left != nil {
				st.Push(node.Left)
			}
		}
	}

	return result
}

func (t *treeNode[T]) PostOrderTraversal(result *[]T) {
	if t == nil {
		return
	}
	t.Left.PostOrderTraversal(result)
	t.Right.PostOrderTraversal(result)
	*result = append(*result, t.Val)
}

func (t *treeNode[T]) PostOrderTraversalByIterate() []T {
	result := make([]T, 0)
	if t == nil {
		return result
	}

	st := NewStack[*treeNode[T]]()
	st.Push(t)

	for st.Size() != 0 {
		node, res := st.Pop()
		if !res {
			return result
		}

		if node == nil {
			node, res = st.Pop()
			if !res {
				return result
			}
			result = append(result, node.Val)
		} else {
			st.Push(node)
			st.Push(nil)

			if node.Right != nil {
				st.Push(node.Right)
			}

			if node.Left != nil {
				st.Push(node.Left)
			}
		}
	}
	return result
}

func (t *treeNode[T]) LevelTraversal(result *[]T) {
	queue := NewQueue[*treeNode[T]]()
	if t == nil {
		return
	}

	queue.Enqueue(t)
	for !queue.IsEmpty() {
		node, res := queue.Dequeue()
		if !res {
			return
		}
		*result = append(*result, node.Val)
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}
}
