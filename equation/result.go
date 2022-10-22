package equation

import (
	"fmt"
	"strconv"
)

// Структура для результата решения уравнения
type Iteration struct {
	k   int
	a   float64
	b   float64
	x   float64
	y   float64
	eps float64
}

type NumericalResult struct {
	iters []Iteration
	x     float64
	eps   float64
}

func (self NumericalResult) String() string {
	s := fmt.Sprintf("Решение: x=%v eps=%v\n\n", self.x, self.eps)
	s += fmt.Sprintf(
		"%-6s %-10s %-10s %-10s %-10s %-10s\n",
		"k", "a", "b", "x", "f(x)", "|b-a|",
	)
	for _, iter := range self.iters {
		s += fmt.Sprintf(
			"%-3s %-6s %-6s %-6s %-6s %-6s\n",
			strconv.Itoa(iter.k),
			fmt.Sprintf("%10.6f", iter.a),
			fmt.Sprintf("%10.6f", iter.b),
			fmt.Sprintf("%10.6f", iter.x),
			fmt.Sprintf("%10.6f", iter.y),
			fmt.Sprintf("%10.6f", iter.eps),
		)
	}
	return s

}
