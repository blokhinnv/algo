package equation

import (
	"fmt"
	"math"
	"strconv"
)

// Структура для итерации алгоритма с уточнением интервала
type RangeIteration struct {
	k   int
	a   float64
	b   float64
	x   float64
	y   float64
	eps float64
}

// Структура для результата решения уравнения с уточнением интервала
type NumericalRangeResult struct {
	iters []RangeIteration
	x     float64
	eps   float64
}

// Метод для строкового представления ответа
func (r NumericalRangeResult) String() string {
	s := fmt.Sprintf("Решение: x=%v eps=%v\n\n", r.x, r.eps)
	s += fmt.Sprintf(
		"%-6s %-10s %-10s %-10s %-10s %-10s\n",
		"k", "a", "b", "x", "f(x)", "|b-a|",
	)
	for _, iter := range r.iters {
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

// Структура для итерации алгоритма с уточнением самого корня
type SimpleIteration struct {
	k   int
	x   float64
	y   float64
	eps float64
}

// Структура для ответа алгоритма с уточнением самого корня
type NumericalResult struct {
	iters []SimpleIteration
	x     float64
	eps   float64
}

// Метод для строкового представления ответа
func (r NumericalResult) String() string {
	s := fmt.Sprintf("Решение: x=%v eps=%v\n\n", r.x, r.eps)
	s += fmt.Sprintf(
		"%-6s %-10s %-10s %-10s\n",
		"k", "x", "f(x)", "|b-a|",
	)
	for _, iter := range r.iters {
		s += fmt.Sprintf(
			"%-3s %-6s %-6s %-6s\n",
			strconv.Itoa(iter.k),
			fmt.Sprintf("%10.6f", iter.x),
			fmt.Sprintf("%10.6f", iter.y),
			fmt.Sprintf("%10.6f", iter.eps),
		)
	}
	return s

}

var NEGATIVE_RESULT = NumericalResult{nil, math.Inf(1), math.Inf(1)}
