package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeltaToPrecision(t *testing.T) {
	t.Run("1e-2", func(t *testing.T) {
		assert.Equal(t, uint(2), DeltaToPrecision(0.01))
	})

	t.Run("5e-2", func(t *testing.T) {
		assert.Equal(t, uint(2), DeltaToPrecision(0.05))
	})

	t.Run("1e-5", func(t *testing.T) {
		assert.Equal(t, uint(5), DeltaToPrecision(0.00001))
	})
}

func TestRoundFloat(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 0.001, RoundFloat(0.00123, 3))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 0.008, RoundFloat(0.00765, 3))
	})

}

func TestNewFunction(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		var y, dy func(x float64) float64

		f := Function{
			Y:  y,
			Dy: dy,
		}
		f_from_func := NewFunction(y, dy)
		assert.Equal(t, f, f_from_func)
	})

}
