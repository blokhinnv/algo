package math

func Any(vs []float64, f func(float64) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []float64, f func(float64) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}
