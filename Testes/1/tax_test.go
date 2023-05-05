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

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalcTax(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1500.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalcTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("expected %f, got %f", amount, result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("expected %f, got %f", amount, result)
		}
	})
}
