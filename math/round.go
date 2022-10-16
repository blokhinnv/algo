package math

import "math"

// Функция для преобразования точности
// в кол-во десятичных знаков для округления
func DeltaToPrecision(eps float64) (p uint) {
	for eps < 1 {
		p += 1
		eps *= 10
	}
	return
}

// Функция для округления числа до указанного
// кол-ва знаков
func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
