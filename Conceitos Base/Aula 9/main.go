package main

type MyNumber int

type Number interface {
	~int | float64
}

func somaSal[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma = soma + v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Gustavo": 3000, "Leonardo": 2000}
	n := map[string]MyNumber{"Gustavo": 3000, "Leonardo": 2000}
	println(somaSal(m))
	println(somaSal(n))
	println(Compara(1, 2))
}
