// Пакет для вспомогательных математических функций
// и объектов
package math

import "gonum.org/v1/gonum/diff/fd"

// Структура для функции и ее производной
type Function struct {
	Y   func(x float64) float64
	Dy  func(x float64) float64
	D2y func(x float64) float64
}

// Конструктор функции
func NewFunctionNoDerivatives(
	y func(x float64) float64,
) Function {
	// если функция производной не задана
	// используем оценку на основе конечных
	// разностей
	dy := func(x float64) float64 {
		return fd.Derivative(y, x, &fd.Settings{
			Formula: fd.Forward,
			Step:    1e-3,
		})
	}

	d2y := func(x float64) float64 {
		return fd.Derivative(dy, x, &fd.Settings{
			Formula: fd.Forward,
			Step:    1e-3,
		})
	}
	return Function{
		Y:   y,
		Dy:  dy,
		D2y: d2y,
	}
}

// Конструктор функции
func NewFunctionWithDerivative(
	y func(x float64) float64,
	dy func(x float64) float64,
) Function {
	return Function{
		Y:   y,
		Dy:  dy,
		D2y: nil,
	}
}
