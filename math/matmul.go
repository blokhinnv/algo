package math

// матричное умножение
func MatMul(A, B Matrix) Matrix {
	m, n, p := len(A), len(A[0]), len(B[0])
	r := make(Matrix, m)
	for i := 0; i < m; i++ {
		r[i] = make([]float64, p)
	}

	for i := 0; i < m; i++ { // строки A
		for j := 0; j < p; j++ { // столбцы B
			for k := 0; k < n; k++ { // строки B
				r[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return r
}
