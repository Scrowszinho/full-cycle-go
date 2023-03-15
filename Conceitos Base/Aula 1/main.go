package main

import "fmt"

type ID int

var (
	a                 = true
	b         ID      = 1
	c                 = "Gustavo"
	d         float64 = 1.33
	testArray [3]int
)

func main() {
	testArray[0] = 1
	testArray[1] = 2
	testArray[2] = 3
	for i, v := range testArray {
		fmt.Printf("Index %d = %d\n", i, v)
	}
}
