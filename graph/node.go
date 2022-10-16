package graph

import (
	"fmt"

	gp "github.com/dominikbraun/graph"
	"github.com/google/uuid"
	"github.com/kr/pretty"
)

type Number interface {
	int | int32 | float32 | int64 | float64
}

type Node[T Number] struct {
	value   T
	name    string
	op      func(T, T) T
	parents []*Node[T]
}

func NewNode[T Number](
	value T,
	op func(T, T) T,
	parents []*Node[T],
	name any,
	prefix any,
) *Node[T] {
	prefixStr, ok := prefix.(string)
	if !ok {
		prefixStr = ""
	}

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
