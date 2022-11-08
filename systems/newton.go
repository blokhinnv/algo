package systems

import (
	m "algo/math"
)

// Метод Ньютона для векторной функции
func SolveNewton(F m.VariadicFunction, x0 m.Vector, eps float64) NumericalResult {
	xPrev := x0
	xCurr := x0
	iters := []SimpleIteration{{0, xCurr, F.Call(xCurr), nil}}
	k := 1
	for {
		xPrev = xCurr
		J := F.Jacobian(xPrev)
		fAtx := F.Call(xPrev)
		delta := m.MatMul(m.InverseMatrix(J, false), fAtx.ToMatrix()).ToVector()
		xCurr = xPrev.Minus(delta)
		iters = append(
			iters,
			SimpleIteration{k, xCurr, F.Call(xCurr), delta},
		)
		k += 1
		if delta.MaxNorm() < eps {
			break
		}
	}
	return NumericalResult{
		iters,
		iters[len(iters)-1].x,
		eps,
	}

}
