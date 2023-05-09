package main

import "fmt"

func main() {
	hello := make(chan string)
	go recebe("hello", hello)
	ler(hello)
}

func recebe(name string, hello chan<- string) {
	hello <- name
}

func ler(data <-chan string) {
	fmt.Println(<-data)
}
