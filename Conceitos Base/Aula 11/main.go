package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numbers := []string{"Um", "Dois", "Tres"}
	for k, v := range numbers {
		println(k, v)
	}

}
