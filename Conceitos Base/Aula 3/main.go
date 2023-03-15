package main

import "fmt"

func main() {
	data := map[string]int{"Gustavo": 1000}
	fmt.Println(data["Gustavo"])
	delete(data, "Gustavo")
	data["Gustav"] = 3000
	fmt.Println(data["Gustav"])

	for name, salar := range data {
		fmt.Printf("Nome: %s - Salario: %d\n", name, salar)
	}

	for _, salar := range data {
		fmt.Printf("Salario: %d\n", salar)
	}
}
