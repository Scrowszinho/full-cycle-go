package tax

import (
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

func (t *TaxRepositoryMock) SaveTax(amount float64) error {
	args := t.Called(amount)
	return args.Error(0)
}
