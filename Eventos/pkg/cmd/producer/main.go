package main

import (
	"teste/Eventos/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	rabbitmq.Publish(ch, "Testando publish", "amq.direct")
}
