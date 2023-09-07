package main

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{
		"Leonardo": 1000,
		"Edi":      2000,
		"Maria":    3000,
	}

	m2 := map[string]float64{
		"Leonardo": 100.10,
		"Edi":      200.00,
		"Maria":    300.00,
	}

	println(Soma(m))
	println(Soma(m2))
}
