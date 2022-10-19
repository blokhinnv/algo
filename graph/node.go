// Пакет для построения графа вычислений
package graph

import (
	"fmt"

	gp "github.com/dominikbraun/graph"
	"github.com/google/uuid"
	"github.com/kr/pretty"
)

// Числовой тип
type Number interface {
	int | int32 | float32 | int64 | float64
}

// Базовый тип для узла графа вычислений
type Node[T Number] struct {
	value   T
	name    string
	op      func(T, T) T
	parents []*Node[T]
}

// Общий конструктор узлов графа вычислений
func NewNode[T Number](
	value T,
	op func(T, T) T,
	parents []*Node[T],
	name any,
	prefix any,
) *Node[T] {
	// prefix обычно означает операцию
	prefixStr, ok := prefix.(string)
	if !ok {
		prefixStr = ""
	}
	// генерируем имя узла, если оно не задано
	nodeName, ok := name.(string)
	if !ok {
		nodeName = uuid.New().String()
	}

	finalName := fmt.Sprintf("%v%v", prefixStr, nodeName)

	return &Node[T]{
		value:   value,
		name:    finalName,
		op:      op,
		parents: parents,
	}
}

// Конструктор именованных узлов без операций
func NewNamedNode[T Number](value T, name string) *Node[T] {
	return NewNode(value, nil, nil, name, nil)
}

func (self *Node[T]) GetParents() []*Node[T] {
	return self.parents
}

func (self *Node[T]) GetValue() T {
	return self.value
}

func (self *Node[T]) GetName() string {
	return self.name
}

func (self *Node[T]) String() string {
	return fmt.Sprintf("%# v", pretty.Formatter(self))

}

// Реализует обход в ширину вверх, начиная с узла self
// для построения графа
func (self *Node[T]) ToGraph() gp.Graph[string, string] {
	g := gp.New(gp.StringHash, gp.Directed())
	queue := []*Node[T]{self}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		g.AddVertex(v.GetName())
		for _, parent := range v.GetParents() {
			g.AddVertex(parent.GetName())
			g.AddEdge(parent.GetName(), v.GetName())
			queue = append(queue, parent)
		}
	}
	return g
}
