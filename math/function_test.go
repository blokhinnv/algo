package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncConstructor(t *testing.T) {
	t.Run("x^2", func(t *testing.T) {
		var f = NewFunctionNoDerivatives(
			func(x float64) float64 {
				return x * x
			},
		)

		assert.Equal(t, 4.0, f.Y(2))
		assert.InDelta(t, 4.0, f.Dy(2), 1e-3)
		assert.InDelta(t, 2, f.D2y(2), 1e-3)
	})
}
