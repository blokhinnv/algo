// Пакет для вспомогательных математических функций
// и объектов
package math

import "gonum.org/v1/gonum/diff/fd"

// Структура для функции и ее производной
type Function struct {
	Y  func(x float64) float64
	Dy func(x float64) float64
}

// Конструктор функции
func NewFunction(
	y func(x float64) float64,
	dy func(x float64) float64,
) Function {
	// если функция производной не задана
	// используем оценку на основе конечных
	// разностей
	if dy == nil {
		dy = func(x float64) float64 {
			return fd.Derivative(y, x, &fd.Settings{
				Formula: fd.Forward,
				Step:    1e-3,
			})
		}
		return Function{
			Y:  y,
			Dy: dy,
		}
	}
	return Function{
		Y:  y,
		Dy: dy,
	}
}
