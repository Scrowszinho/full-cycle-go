package main

import (
	"errors"
	"fmt"
)

func main() {
	resul, erro := sum(1, 2)
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(resul)
}

func sum(a int, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("Erro")
	}
	return a + b, nil
}
