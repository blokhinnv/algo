package systems

import (
	m "algo/math"
	"fmt"
)

// Выводит на экран текущую матрицу системы и вектор свободных членов в
// форматированном виде
func printIter(k int, A m.Matrix, b m.Vector, isForward bool) {
	n := len(A)

	arrow := "↓"
	if !isForward {
		arrow = "↑"
	}
	fmt.Printf("[Iteration %d %v]\n", k, arrow)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%5.2f ", m.RoundFloat(A[i][j], 2))
		}

		fmt.Printf("| %5.2f \n", m.RoundFloat(b[i], 2))
	}
	fmt.Println()
}

// Классический метод Гаусса
func SolveGauss(A_in m.Matrix, b_in m.Vector, verbose bool) m.Vector {
	A := A_in.Copy()
	b := b_in.Copy()

	n := len(A)
	for k := 0; k < n-1; k++ { // итерации прямого хода
		for i := k + 1; i < n; i++ { // проход строк ниже k
			m := A[i][k] / A[k][k]
			for j := 0; j < n; j++ { // обход столбцов iй строки
				A[i][j] = A[i][j] - m*A[k][j]
			}

			b[i] = b[i] - m*b[k]
		}
		if verbose {
			printIter(k, A, b, true)
		}
	}

	xs := make(m.Vector, n)

	for k := n - 1; k >= 0; k-- { // итерация обратного хода
		x := b[k]
		for i := n - 1; i >= k; i-- { // вычитаем известные для данной строчки x
			x -= A[k][i] * xs[i]
		}
		x /= A[k][k] // нормируем и получаем новый x
		xs[k] = x
	}
	return xs
}

// Метод Жордана-Гаусса
func SolveGaussBackward(A_in m.Matrix, b_in m.Vector, verbose bool) m.Vector {
	A := A_in.Copy()
	b := b_in.Copy()

	n := len(A)
	nForwardIters := n - 1
	for k := 0; k < n-1; k++ { // итерации прямого хода
		for i := k + 1; i < n; i++ { // проход строк ниже k
			m := A[i][k] / A[k][k]
			for j := 0; j < n; j++ { // обход столбцов iй строки
				A[i][j] = A[i][j] - m*A[k][j]
			}

			b[i] = b[i] - m*b[k]
		}
		if verbose {
			printIter(k, A, b, true)
		}
	}

	for k := n - 1; k >= 0; k-- { // итерация обратного хода
		b[k] /= A[k][k]
		A[k][k] = 1
		for i := k - 1; i >= 0; i-- {
			b[i] -= A[i][k] * b[k]
			A[i][k] = 0
		}
		printIter(2*nForwardIters-k, A, b, false)
	}
	return b
}
