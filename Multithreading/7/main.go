package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Teste"}
			time.Sleep(time.Second * 1)
			c1 <- msg
		}
	}()
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Teste 2"}
			time.Sleep(time.Second * 1)
			c2 <- msg
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("recived from %d - %s\n", msg1.id, msg1.Msg)
		case msg2 := <-c2:
			fmt.Printf("recived from %d - %s\n", msg2.id, msg2.Msg)
		case <-time.After(time.Second * 3):
			println("timeout")
			// default:
			// 	println("default")
		}
	}
}
