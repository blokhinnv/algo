package differential

import (
	m "algo/math"
	"fmt"
	"math"
	"reflect"

	"gonum.org/v1/plot/plotter"
)

// Табличное представление функции
type TabularFunction struct {
	x []float64
	y []float64
}

// Метод для добавления новой точки
func (tf *TabularFunction) add(x, y float64) {
	tf.x = append(tf.x, x)
	tf.y = append(tf.y, y)
}

// Преобразование TabularFunction в plotter.XYs
func (tf TabularFunction) toXYs() plotter.XYs {
	n := len(tf.x)
	pts := make(plotter.XYs, n)
	for i := range pts {
		pts[i].X = tf.x[i]
		pts[i].Y = tf.y[i]
	}
	return pts
}

// Метод для строкового представления ответа
func (tf TabularFunction) String() string {
	n := len(tf.x)
	s := fmt.Sprintf(
		"%5s %5s\n",
		"x", "y",
	)
	for i := 0; i < n; i++ {
		s += fmt.Sprintf(
			"%-6s %-6s\n",
			fmt.Sprintf("%6.5f", tf.x[i]), fmt.Sprintf("%6.5f", tf.y[i]),
		)
	}
	return s
}

// Метод для вычисления расстояния между двумя табличными функциями
func (tf TabularFunction) ComputeDissimilarity(other TabularFunction) float64 {
	if !reflect.DeepEqual(tf.x, other.x) {
		return math.Inf(1)
	}
	m := len(tf.x)
	total_diff := 0.0
	for i := 0; i < m; i++ {
		total_diff += math.Abs(tf.y[i] - other.y[i])
	}
	return total_diff
}

// Конструктор пустой табличной функции
func NewTabularFunction() TabularFunction {
	return TabularFunction{make([]float64, 0), make([]float64, 0)}
}

// Конструктор табличной функции
func NewTabularFunctionFromF(
	f func(x float64) float64,
	a, b float64,
	h float64,
) TabularFunction {
	tf := TabularFunction{make([]float64, 0), make([]float64, 0)}
	for _, x := range m.Arange(a, b, h) {
		tf.add(x, f(x))
	}
	return tf
}
