// Пакет для интегрирования функций
package integrate

import (
	m "algo/math"
	"math"
)

// Структура для результата интегрированя
type Result struct {
	value   float64
	abs_err float64
}

// Функция для расчета абсолютной ошибки интегрирования
func AbsErrorTrapezoids(
	df func(x float64) float64,
	a float64,
	b float64,
	n int,
) float64 {
	return -math.Pow(b-a, 2) / (12 * math.Pow(float64(n), 2)) * (df(b) - df(a))
}

// Функция для расчета интеграла методом трапеций с указанием кол-ва интервалов
func ComputeTrapezoidsN(
	f func(float64) float64,
	a float64,
	b float64,
	n int,
) (res float64) {
	h := (b - a) / float64(n)

	for i := 0; i <= n; i++ {
		fx := f(a + float64(i)*h)
		res += 2 * fx
		if i == 0 || i == n {
			res -= fx
		}
	}
	res *= h / 2
	return
}

// Функция для расчета интеграла методом трапеций
// с указанием кол-ва интервалов и абсолютной ошибки
func TrapezoidsN(f m.Function, a float64, b float64, n int) Result {
	return Result{
		value:   ComputeTrapezoidsN(f.Y, a, b, n),
		abs_err: AbsErrorTrapezoids(f.Dy, a, b, n),
	}
}

// Функция для расчета интеграла методом трапеций
// с указанием ошибки метода
func TrapezoidsDelta(f m.Function, a float64, b float64, delta float64) Result {
	n := 10
	i1 := ComputeTrapezoidsN(f.Y, a, b, n)
	n *= 2
	i2 := ComputeTrapezoidsN(f.Y, a, b, n)
	for math.Abs(i2-i1) > delta {
		i1 = i2
		n *= 2
		i2 = ComputeTrapezoidsN(f.Y, a, b, n)
	}
	return Result{
		value:   m.RoundFloat(i2, m.DeltaToPrecision(delta)),
		abs_err: delta,
	}
}
