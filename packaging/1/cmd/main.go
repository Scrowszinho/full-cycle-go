package main

import (
	"fmt"

	"github.com/scrowszinho/teste/math"
)

func main() {
	a := math.NewMath(1, 2)
	fmt.Println(a.Add())
}
