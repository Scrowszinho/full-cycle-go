package main

import (
	"curso-go/matematica"
	"fmt"
)

func main() {
	soma := matematica.Soma(10, 20)
	carro := matematica.Carro{name: "test"}
	fmt.Printf(carro)
	fmt.Println("Resultado: %v", soma)
}
