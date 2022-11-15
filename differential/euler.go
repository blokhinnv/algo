package differential

// Метод Эйлера
func SolveEuler(f func(x, y float64) float64, x0, y0, h, b float64) TabularFunction {
	res := NewTabularFunction()
	currX, currY := x0, y0
	res.add(currX, currY)
	for currX+h < b {
		currY += h * f(currX, currY)
		currX += h
		res.add(currX, currY)
	}
	return res
}
