package math

import (
	"fmt"
)

type Matrix [][]float64

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

// Выводит на экран текущую матрицу преобразования
// в форматированном виде
func printIterInv(k int, A Matrix, isForward bool) {
	n := len(A)

	arrow := "↓"
	if !isForward {
		arrow = "↑"
	}
	fmt.Printf("[Iteration %d %v]\n", k, arrow)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%5.2f ", RoundFloat(A[i][j], 2))
		}

		fmt.Printf(" | ")
		for j := n; j < 2*n; j++ {
			fmt.Printf("%5.2f ", RoundFloat(A[i][j], 2))
		}
		fmt.Println()
	}
	fmt.Println()

}

// Вычисляет обратную матрицу путем элементарных преобразований
func InverseMatrix(A_in Matrix, verbose bool) Matrix {
	A := make([][]float64, len(A_in))
	copy(A, A_in)

	n := len(A)
	for i := 0; i < n; i++ {
		eye_row := make([]float64, n)
		eye_row[i] = 1
		A[i] = append(A[i], eye_row...)
	}
	if verbose {
		printIterInv(0, A, true)
	}
	nForwardIters := n - 1
	for k := 0; k < n-1; k++ { // итерации прямого хода
		for i := k + 1; i < n; i++ { // проход строк ниже k
			m := A[i][k] / A[k][k]
			for j := 0; j < 2*n; j++ { // обход столбцов iй строки
				A[i][j] = A[i][j] - m*A[k][j]
			}
		}
		if verbose {
			printIterInv(k+1, A, true)
		}
	}

	for k := n - 1; k >= 0; k-- { // итерация обратного хода
		for j := n; j < 2*n; j++ {
			A[k][j] /= A[k][k]
		}
		A[k][k] = 1
		for i := k - 1; i >= 0; i-- {
			for j := n; j < 2*n; j++ { // Обновление левой части
				A[i][j] = A[i][j] - A[k][j]*A[i][k]
			}
			A[i][k] = 0
		}
		if verbose {
			printIterInv(2*nForwardIters-k+1, A, false)
		}
	}

	invA := make([][]float64, n)
	for i := 0; i < n; i++ {
		invA[i] = A[i][n:]
	}
	return invA
}
