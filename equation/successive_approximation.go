package equation

import (
	m "algo/math"
	"math"
)

func SuccessiveApprox(f m.Function, a, b, eps float64) NumericalResult {
	space := m.Linspace(a, b, 100)

	// метод сходится, если производная на отрезке по модулю меньше 1
	if !m.All(space, func(x float64) bool { return math.Abs(f.Dy(x)) < 1 }) {
		return NEGATIVE_RESULT
	}
	iters := make([]SimpleIteration, 0)
	xCurr := (a + b) / 2
	xPrev := xCurr
	k := 0
	for {
		xPrev = xCurr
		xCurr = f.Y(xPrev)
		delta := math.Abs(xCurr - xPrev)

		iters = append(
			iters,
			SimpleIteration{k, xCurr, f.Y(xCurr), delta},
		)

		if delta < eps {
			break
		}
		k += 1
	}
	return NumericalResult{
		iters,
		m.RoundFloat(
			iters[len(iters)-1].x,
			m.DeltaToPrecision(eps)+1,
		),
		eps,
	}

}
