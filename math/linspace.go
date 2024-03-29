package math

import "math"

const tol = 1e-6

// Генерирует N точек на полуинтервале
func Linspace(start, stop float64, N int) []float64 {
	rnge := make([]float64, N)
	step := (stop - start) / float64(N)
	i := 0
	for x := start; x < stop; x += step {
		rnge[i] = x
		i += 1
	}
	return rnge
}

// Генерирует точки на полуинтервале с заданным шагов
func Arange(start, stop, step float64) []float64 {
	N := int(math.Ceil((stop - start) / step))
	rnge := make([]float64, N+1)
	i := 0
	for x := start; x < stop+tol; x += step {
		rnge[i] = x
		i += 1
	}
	return rnge[:N]
}
