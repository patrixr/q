package q

type Node[T any] struct {
	Root     bool
	Data     T
	Children []*Node[T]
}

func Tree[T any](val T) *Node[T] {
	return &Node[T]{
		Root: true,
		Data: val,
	}
}

func (tree *Node[T]) Add(child T) *Node[T] {
	node := Node[T]{
		Data: child,
	}
	tree.Children = append(tree.Children, &node)
	return &node
}

func (tree *Node[T]) FindChild(predicate func(T) bool) *Node[T] {
	for _, child := range tree.Children {
		if predicate(child.Data) {
			return child
		}
	}

	return nil
}

func (tree *Node[T]) Traverse(cb func(it T)) {
	cb(tree.Data)

	for _, child := range tree.Children {
		child.Traverse(cb)
	}
}
