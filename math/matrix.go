package math

import "fmt"

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
