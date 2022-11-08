package math

import (
	"gonum.org/v1/gonum/diff/fd"
	"gonum.org/v1/gonum/mat"
)

// Структура для векторной функции векторного аргумента
type VariadicFunction struct {
	Y func(dst, x []float64)
	M int
	N int
}

// Преобразует BLAS матрицу в матрицу Matrix
func matToMatrix(jac *mat.Dense) Matrix {
	m, _ := jac.Dims()
	res := make(Matrix, m)
	for i := 0; i < m; i++ {
		res[i] = jac.RawRowView(i)
	}
	return res

}

// Возвращает Якобиан векторной функции
func (self VariadicFunction) Jacobian(x Vector) Matrix {
	jac := mat.NewDense(self.M, self.N, nil)
	fd.Jacobian(jac, self.Y, x, &fd.JacobianSettings{
		Formula:    fd.Central,
		Concurrent: true,
	})

	return matToMatrix(jac)
}

// Рассчитывает значение векторной функции в точке
func (self VariadicFunction) Call(x Vector) Vector {
	r := make([]float64, self.M)
	self.Y(r, x)
	return Vector(r).Copy()
}
