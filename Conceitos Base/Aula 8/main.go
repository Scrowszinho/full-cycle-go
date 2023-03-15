package main

import "fmt"

func main() {
	var x interface{} = 10
	showType(x)
}

func showType(t interface{}) {
	fmt.Printf("Type: %T - Value: %v\n", t, t)
}
