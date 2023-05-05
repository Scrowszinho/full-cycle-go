package main

import "time"

func task(name string) {
	for i := 0; i < 10; i++ {
		println(name, ":", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go task("A")
	go task("B")
	go func() {
		for i := 0; i < 5; i++ {
			println("anonymous", ":", i)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(15 * time.Second)
}
