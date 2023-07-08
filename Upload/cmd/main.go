package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	for i < 10 {
		f, err := os.Create(fmt.Sprintf("file%d.txt", i))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString(fmt.Sprintf("file%d", i))
		i++
	}
}
