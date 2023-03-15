package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Addres Address
}

type People interface {
	Desactive()
}

func (c Client) Desactive() {
	c.Active = false
	fmt.Printf("Client: %s is desactivated", c.Name)
}

func Desativacao(pessoa People) {
	pessoa.Desactive()
}

func main() {
	gustavo := Client{
		Name:   "Gustavo",
		Age:    18,
		Active: true,
	}
	gustavo.Desactive()
	Desativacao(gustavo)
	fmt.Println(gustavo.Addres.City)
}
