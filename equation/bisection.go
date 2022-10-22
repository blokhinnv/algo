package equation

import (
	m "algo/math"
	"math"
)

func BisectionEps(f m.Function, a float64, b float64, eps float64) NumericalResult {
	iters := make([]Iteration, 0)
	n_iters := int(math.Ceil(math.Log2((b - a) / eps)))
	for k := 0; k < n_iters; k++ {
		x := (a + b) / 2
		y := f.Y(x)

		if f.Y(a)*f.Y(x) > 0 {
			a = x
		} else {
			b = x
		}

		delta := math.Abs(b - a)

		iters = append(iters, Iteration{k, a, b, x, y, delta})
	}
	return NumericalResult{
		iters,
		m.RoundFloat(
			iters[len(iters)-1].x,
			m.DeltaToPrecision(eps),
		),
		eps,
	}

}
