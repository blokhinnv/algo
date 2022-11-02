package systems

import (
	m "algo/math"
	"fmt"
)

// Решение системы при помощи матричного умножения
func SolveMmul(A m.Matrix, b []float64, verbose bool) []float64 {
	A_inv := m.InverseMatrix(A, false)
	if verbose {
		fmt.Println("A=", A)
		fmt.Println("A_inv=", A_inv)
	}
	b_col := make(m.Matrix, len(b))
	for i := 0; i < len(b); i++ {
		b_col[i] = []float64{b[i]}
	}
	x_col := m.MatMul(A_inv, b_col)
	x := make([]float64, len(x_col))
	for i := 0; i < len(b); i++ {
		x[i] = x_col[i][0]
	}
	return x
}
