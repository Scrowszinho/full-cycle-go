package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcTax(t *testing.T) {
	tax := CalcTax(1000.0)
	assert.Equal(t, 10.0, tax)
}
