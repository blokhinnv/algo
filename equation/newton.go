package equation

import (
	m "algo/math"
	"math"
)

func Newton(f m.Function, a, b, eps float64) NumericalResult {
	var xCurr float64
	if f.D2y(a)*f.Y(a) > 0 {
		xCurr = a
	} else if f.D2y(b)*f.Y(b) > 0 {
		xCurr = b
	} else {
		return NEGATIVE_RESULT
	}

	iters := []SimpleIteration{{0, xCurr, f.Y(xCurr), math.Inf(1)}}
	xPrev := xCurr
	k := 1
	for {
		xPrev = xCurr
		xCurr = xPrev - f.Y(xPrev)/f.Dy(xPrev)
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
			m.DeltaToPrecision(eps),
		),
		eps,
	}

}
