package tax

import "time"

func CalculaTax(amount float64) float64 {
	if amount == 0 {
		return 0
	}
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}

func CalculaTa2(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount == 0 {
		return 0
	}
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}