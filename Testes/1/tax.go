package tax

func CalcTax(amount float64) float64 {
	if amount >= 1000 {
		return 10
	} else {
		return amount * 0.05
	}

}
