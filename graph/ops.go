package graph

// Сложение над числами
func Add[T Number](a, b T) T {
	return a + b
}

// Обратный элемент относительно сложения
func Neg[T Number](a, b T) T {
	return -a
}

// Умножение над числами
func Mul[T Number](a, b T) T {
	return a * b
}

// Декоратор над операцией, который позволяет возвращать узел
// создаваемого графа вычислений
func wrap[T Number](
	op func(a, b T) T, 
	prefix string, 
	ignoreB bool,
) func(a, b any) *Node[T] {
	return func(a, b any) *Node[T] {
		// если a - узел, то ок; если нет, сделаем узел
		aNode, ok := a.(*Node[T])
		if !ok{
			aNode = NewNode(a.(T), nil, nil, nil, nil)
		}

		// если b - узел, то ок; если нет, сделаем узел
		bNode, ok := b.(*Node[T])
		if !ok{
			bNode = NewNode(b.(T), nil, nil, nil, nil)
		}
		
		// Родители создаваемого узла
		newParents := []*Node[T]{aNode}
		// ignoreB => унарная операция
		if !ignoreB{
			newParents = append(newParents, bNode)
		}
		return NewNode(op(aNode.value, bNode.value), op, newParents, nil, prefix)
	}

}

// Задекорированное сложение над числами
func AddNode[T Number](a *Node[T], b any) *Node[T] {
	return wrap(Add[T], "add__", false)(a, b)
}

// Задекорированное умножение над числами
func MulNode[T Number](a *Node[T], b any) *Node[T] {
	return wrap(Mul[T], "mul__", false)(a, b)
}

// Задекорированный унарный минус
func NegNode[T Number](a *Node[T]) *Node[T] {
	return wrap(Neg[T], "neg__", true)(a, 0.0)
}
