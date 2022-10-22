package equation

import (
	m "algo/math"
	"math"
)

func rightSideChords(f m.Function, a float64, b float64, eps float64) []Iteration {
	iters := make([]Iteration, 1)
	iters[0] = Iteration{0, a, b, b, f.Y(b), b - a}
	fa := f.Y(a)
	k := 1
	xCurr, xPrev := b, b
	for {
		xPrev = xCurr
		xCurr = a - fa/(f.Y(xPrev)-fa)*(xPrev-a)
		delta := math.Abs(xCurr - xPrev)

		iters = append(
			iters,
			Iteration{k, a, xCurr, xCurr, f.Y(xCurr), delta},
		)

		if delta < eps {
			break
		}
		k += 1
	}
	return iters
}

func leftSideChords(f m.Function, a float64, b float64, eps float64) []Iteration {
	iters := make([]Iteration, 1)
	iters[0] = Iteration{0, a, b, a, f.Y(a), b - a}
	fb := f.Y(b)
	k := 1
	xCurr, xPrev := a, a
	for {
		xPrev = xCurr
		xCurr = xPrev - f.Y(xPrev)/(fb-f.Y(xPrev))*(b-xPrev)
		delta := math.Abs(xCurr - xPrev)

		iters = append(
			iters,
			Iteration{k, xCurr, b, xCurr, f.Y(xCurr), delta},
		)

		if delta < eps {
			break
		}
		k += 1
	}
	return iters
}

func ChordsEps(f m.Function, a float64, b float64, eps float64) NumericalResult {
	fa := f.Y(a)
	space := m.Linspace(a, b, 100)
	var iters []Iteration = nil
	if m.All(space, func(x float64) bool { return f.D2y(x) >= 0 }) {
		// функция выпукла вверх
		if fa > 0 {
			// формула 5
			iters = rightSideChords(f, a, b, eps)
		} else {
			// формула 6
			iters = leftSideChords(f, a, b, eps)
		}
	} else if m.All(space, func(x float64) bool { return f.D2y(x) <= 0 }) {
		// функция выпукла вниз
		if fa > 0 {
			// формула 6
			iters = leftSideChords(f, a, b, eps)
		} else {
			// формула 5
			iters = rightSideChords(f, a, b, eps)
		}
	} else {
		return NumericalResult{}
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
