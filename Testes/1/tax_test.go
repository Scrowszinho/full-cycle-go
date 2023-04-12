package tax

import "testing"

func TestCalcTax(t *testing.T) {
	amount := 1000
	expectedAmount := 10
	result := tax.CalcTax(amount)
	if result != expectedAmount {
		t.Errorf("expected %d, got %d", expectedAmount, result)
	}
}
