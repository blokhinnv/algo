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

func wrap[T Number](
	op func(a, b T) T, 
	prefix string, 
	ignoreB bool,
) func(a, b any) *Node[T] {
	return func(a, b any) *Node[T] {
		aNode, ok := a.(*Node[T])
		if !ok{
			aNode = NewNode(a.(T), nil, nil, nil, nil)
		}

		bNode, ok := b.(*Node[T])
		if !ok{
			bNode = NewNode(b.(T), nil, nil, nil, nil)
		}
		
		newParents := []*Node[T]{aNode}
		if !ignoreB{
			newParents = append(newParents, bNode)
		}
		return NewNode(op(aNode.value, bNode.value), op, newParents, nil, prefix)
	}

}

func AddNode[T Number](a *Node[T], b any) *Node[T] {
	return wrap(Add[T], "add__", false)(a, b)
}

func MulNode[T Number](a *Node[T], b any) *Node[T] {
	return wrap(Mul[T], "mul__", false)(a, b)
}

func NegNode[T Number](a *Node[T]) *Node[T] {
	return wrap(Neg[T], "neg__", true)(a, 0.0)
}
