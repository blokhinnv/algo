package equation

import (
	m "algo/math"
	"math"
)

// Метод хорд для движения с правой стороны
func rightSideChords(f m.Function, a float64, b float64, eps float64) []RangeIteration {
	iters := make([]RangeIteration, 1)
	iters[0] = RangeIteration{0, a, b, b, f.Y(b), b - a}
	fa := f.Y(a)
	k := 1
	xCurr, xPrev := b, b
	for {
		xPrev = xCurr
		xCurr = a - fa/(f.Y(xPrev)-fa)*(xPrev-a)
		delta := math.Abs(xCurr - xPrev)

		iters = append(
			iters,
			RangeIteration{k, a, xCurr, xCurr, f.Y(xCurr), delta},
		)

		if delta < eps {
			break
		}
		k += 1
	}
	return iters
}

// Метод хорд для движения с левой стороны
func leftSideChords(f m.Function, a float64, b float64, eps float64) []RangeIteration {
	iters := make([]RangeIteration, 1)
	iters[0] = RangeIteration{0, a, b, a, f.Y(a), b - a}
	fb := f.Y(b)
	k := 1
	xCurr, xPrev := a, a
	for {
		xPrev = xCurr
		xCurr = xPrev - f.Y(xPrev)/(fb-f.Y(xPrev))*(b-xPrev)
		delta := math.Abs(xCurr - xPrev)

		iters = append(
			iters,
			RangeIteration{k, xCurr, b, xCurr, f.Y(xCurr), delta},
		)

		if delta < eps {
			break
		}
		k += 1
	}
	return iters
}

// Метод хорд на заданном отрезке с заданной точностью
func ChordsEps(f m.Function, a float64, b float64, eps float64) NumericalRangeResult {
	fa := f.Y(a)
	space := m.Linspace(a, b, 100)
	var iters []RangeIteration = nil
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
		return NumericalRangeResult{}
	}
	return NumericalRangeResult{
		iters,
		m.RoundFloat(
			iters[len(iters)-1].x,
			m.DeltaToPrecision(eps),
		),
		eps,
	}
}
