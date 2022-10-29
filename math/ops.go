package math

// Возвращает true, если f дает true хотя бы на одном примере
func Any(vs []float64, f func(float64) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// Возвращает true, если f дает true на всех примерах
func All(vs []float64, f func(float64) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}
