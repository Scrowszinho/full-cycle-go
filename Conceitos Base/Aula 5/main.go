package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 0) * 2
	}()
	fmt.Println(total)
}

func sum(nums ...int) int {
	total := 0
	for _, nume := range nums {
		total += nume
	}
	return total
}
