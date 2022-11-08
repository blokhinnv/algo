package systems

import (
	m "algo/math"
	"fmt"
	"strconv"
)

// Структура для итерации алгоритма с уточнением самого корня
type SimpleIteration struct {
	k     int
	x     m.Vector
	y     m.Vector
	delta m.Vector
}

// Структура для ответа алгоритма с уточнением самого корня
type NumericalResult struct {
	iters []SimpleIteration
	x     m.Vector
	eps   float64
}

// Метод для строкового представления ответа
func (r NumericalResult) String() string {
	s := fmt.Sprintf("Решение: x=%v; eps=%v\n\n", r.x.StringPrecision(r.eps), r.eps)
	n := len(r.iters[0].x)
	s += fmt.Sprintf(
		"%-6[1]s %-[7]*[2]s %-[6]*[3]s %-[5]*[4]s\n",
		"k", "x", "F(x)", "delta",
		6*n, 6*n, 6*n,
	)

	for _, iter := range r.iters {
		s += fmt.Sprintf(
			"%-3s %-6s %-6s %-6s \n",
			strconv.Itoa(iter.k),
			fmt.Sprintf("%v", iter.x),
			fmt.Sprintf("%v", iter.y),
			fmt.Sprintf("%v", iter.delta),
		)
	}
	return s

}
