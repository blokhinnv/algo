package integrate

import (
	"math"
	"testing"

	m "algo/math"

	"github.com/stretchr/testify/assert"
)

func TestTrapezoidsN(t *testing.T) {
	var f = m.NewFunction(
		func(x float64) float64 {
			return 7 / (x*x + 1)
		},
		nil,
	)

	var g = m.NewFunction(
		func(x float64) float64 {
			return 1 / math.Log(x)
		},
		func(x float64) float64 {
			return -1 / (x * math.Pow(math.Log(x), 2))
		},
	)

	// интеграл из http://www.cleverstudents.ru/integral/method_of_trapezoids.html
	t.Run("integral1", func(t *testing.T) {
		integral := TrapezoidsN(f, 0, 5, 10)
		assert.InDelta(t, 9.61172, integral.value, 1e-5)
		assert.InDelta(t, 0.002, integral.abs_err, 1e-3)
	})

	// интеграл из http://mathprofi.ru/formula_simpsona_metod_trapecij.html
	t.Run("integral2", func(t *testing.T) {
		integral := TrapezoidsN(g, 2, 5, 5)
		assert.InDelta(t, 2.617, integral.value, 1e-3)
	})

}

func TestTrapezoidsDelta(t *testing.T) {
	var h = m.NewFunction(
		func(x float64) float64 {
			return x * math.Exp(x)
		},
		func(x float64) float64 {
			return 0
		},
	)

	// интеграл из http://www.cleverstudents.ru/integral/method_of_trapezoids.html
	t.Run("integral1", func(t *testing.T) {
		integral := TrapezoidsDelta(h, 0, 2, 0.001)
		assert.InDelta(t, 8.389, integral.value, 1e-3)
	})
}
