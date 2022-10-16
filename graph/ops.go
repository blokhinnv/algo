package graph

func Add[T Number](a, b T) T {
	return a + b
}

func Neg[T Number](a, b T) T {
	return -a
}

func Mul[T Number](a, b T) T {
	return a * b
}

func wrap[T Number](op func(a, b T) T, prefix string) func(a, b *Node[T]) *Node[T] {
	return func(a, b *Node[T]) *Node[T] {
		newParents := []*Node[T]{a, b}
		return NewNode(op(a.value, b.value), Add[T], newParents, nil, prefix)
	}

}

func AddNode[T Number](a, b *Node[T]) *Node[T] {
	return wrap(Add[T], "add__")(a, b)
}

func MulNode[T Number](a, b *Node[T]) *Node[T] {
	return wrap(Mul[T], "mul__")(a, b)
}

func NegNode[T Number](a, b *Node[T]) *Node[T] {
	return wrap(Neg[T], "neg__")(a, b)
}
