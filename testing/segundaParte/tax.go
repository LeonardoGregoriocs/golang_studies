package tax

import "errors"

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculaTaxTwo(amount)
	return repository.SaveTax(tax)
}

func CalculaTaxTwo(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20.0
	}
	return 5.0
}

func CalculaTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("amout must be greater than 0")
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0, nil
	}
	if amount >= 20000 {
		return 20.0, nil
	}
	return 5.0, nil
}
