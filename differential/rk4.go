package differential

const tol = 1e-6

// Метод Рунге-Кутта 4 порядка
func SolveRK4(f func(x, y float64) float64, x0, y0, h, b float64) TabularFunction {
	res := NewTabularFunction()
	currX, currY := x0, y0
	res.add(currX, currY)
	for currX+h < b+tol {
		b1 := f(currX, currY)
		b2 := f(currX+h/2, currY+h*b1/2)
		b3 := f(currX+h/2, currY+h*b2/2)
		b4 := f(currX+h, currY+h*b3)
		currY += h / 6 * (b1 + 2*b2 + 2*b3 + b4)
		currX += h
		res.add(currX, currY)
	}
	return res
}
