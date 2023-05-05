package tax

func CalcTax(amount float64) float64 {
	if amount >= 1000 {
		return 10.0
	}
	if amount <= 0 {
		return 0.0
	} else {
		return amount * 0.05
	}

}

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalcTax(amount)
	return repository.SaveTax(tax)
}
