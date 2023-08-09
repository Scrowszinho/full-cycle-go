package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	order, err := NewOrder("123", 123, 23)
	assert.NoError(t, err)
	assert.Equal(t, order.FinalPrice, 146)
}
