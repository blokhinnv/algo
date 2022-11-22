package differential

import (
	"math"
)

// Метод прогноза и коррекции
func SolvePredictorCorrector(
	f func(x, y float64) float64,
	x0, y0, h, b, eps float64,
) TabularFunction {
	correct := func(currX, currY, yPred float64) float64 {
		yCorr := yPred
		for {
			yCorrNext := currY + h/2*(f(currX, currY)+f(currX+h, yCorr))
			if math.Abs(yCorrNext-yCorr) < eps {
				return yCorrNext
			}
			yCorr = yCorrNext
		}

	}
	res := NewTabularFunction()
	currX, currY, prevY := x0, y0, y0
	res.add(currX, currY)
	useEuler := true
	k := 0
	for currX+h < b+tol {
		// fmt.Printf("k=%v x_k=%v y_k=%v y_k-1=%v\n", k, currX, currY, prevY)
		if useEuler {
			prevY = currY
			currY += h * f(currX, currY)
			useEuler = false
		} else {
			currYPred := prevY + 2*h*f(currX, currY)
			prevY = currY
			currY = correct(currX, currY, currYPred)
		}
		currX += h
		res.add(currX, currY)
		k += 1
	}

	return res
}
