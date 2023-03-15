package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	size, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Size: %d bytes", size)
	f.Close()

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// Leitura byte a byte
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
