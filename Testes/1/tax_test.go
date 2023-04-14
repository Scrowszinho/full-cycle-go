package tax

import "testing"

func TestCalcTax(t *testing.T) {
	amount := 1000.0
	expectedAmount := 10.0
	result := CalcTax(amount)
	if result != expectedAmount {
		t.Errorf("expected %f, got %f", expectedAmount, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{3000.0, 10.0},
	}

	for _, tc := range table {
		result := CalcTax(tc.amount)
		if result != tc.expected {
			t.Errorf("expected %f, got %f", tc.expected, result)
		}
	}
}
