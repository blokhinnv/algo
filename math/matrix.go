package math

import (
	"fmt"
	"math"
	"strings"
)

type Matrix [][]float64
type Vector []float64

// Метод для красивого вывода матрицы
func (r Matrix) String() string {
	m, n := len(r), len(r[0])
	s := "\n"
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			s += fmt.Sprintf("%5.2f ", RoundFloat(r[i][j], 2))
		}
		s += "\n"
	}
	return s
}

// Метод для копирования матрицы
func (r Matrix) Copy() Matrix {
	m, n := len(r), len(r[0])
	x := make(Matrix, m)
	for i := 0; i < m; i++ {
		x[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			x[i][j] = r[i][j]
		}
	}
	return x
}

// Функция для вычитания двух матриц
func (s Matrix) Minus(o Matrix) Matrix {
	// TODO: проверить размерности!
	r := s.Copy()
	m, n := len(r), len(r[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r[i][j] -= o[i][j]
		}
	}
	return r
}

// Преобразует матрицу из одного столбца в вектор
func (r Matrix) ToVector() Vector {
	m, n := len(r), len(r[0])
	if n > 1 {
		panic("Trying to convert matrix to vector")
	} else {
		v := make(Vector, m)
		for i := 0; i < m; i++ {
			v[i] = r[i][0]
		}
		return v
	}

}

// Метод для красивого вывода вектора
func (r Vector) String() string {
	m := len(r)
	s := ""
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("%5.2f ", RoundFloat(r[i], 2))
	}
	return s
}

// Метод для копирования вектора
func (r Vector) Copy() Vector {
	x := make(Vector, len(r))
	copy(x, r)
	return x
}

// Преобразует вектор в матрицу из одного столбца
func (r Vector) ToMatrix() Matrix {
	m := make(Matrix, len(r))
	for i := 0; i < len(m); i++ {
		m[i] = []float64{r[i]}
	}
	return m
}

// Вычитает один вектор из другого
func (s Vector) Minus(o Vector) Vector {
	r := s.Copy()
	m := len(r)
	for i := 0; i < m; i++ {
		r[i] -= o[i]
	}
	return r
}

// Считает l_inf норму вектора
func (v Vector) MaxNorm() float64 {
	currMax := 0.0
	for _, v := range v {
		x := math.Abs(v)
		if x > currMax {
			currMax = x
		}
	}
	return currMax
}

// Преобразует вектор к строке, предварительно
// округляя значения вектора
func (v Vector) StringPrecision(eps float64) string {
	r := v.Copy()
	xs := make([]string, len(v))
	for _, v := range r {
		xs = append(
			xs,
			fmt.Sprintf(
				"%v",
				RoundFloat(
					v,
					DeltaToPrecision(eps),
				),
			),
		)
	}

	return strings.Join(xs, " ")
}
