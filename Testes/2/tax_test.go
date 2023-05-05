package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcTax(t *testing.T) {
	tax := CalcTax(1000.0)
	assert.Equal(t, 10.0, tax)
}

func TestCalcTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("Error saving tax"))
	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	repository.AssertExpectations(t)

}
