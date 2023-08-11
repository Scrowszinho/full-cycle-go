package main

import (
	"database/sql"
	"fmt"
	"teste/configs"
	"teste/internal/events"

	"github.com/streadway/amqp"
)

func main() {
	configs, err := configs.GetConfigs(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort))
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()
	eventDispatcher := events.NewEventDispatcher()

	eventDispatcher.Register("OrderCreated", &handlerOrderCreated{
		RabbitMQChannel: rabbitMQChannel,
	})
}

func getRabbitMQChannel() *amqp.Channel {
	configs, err := configs.GetConfigs(".")
	if err != nil {
		panic(err)
	}
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", configs.RabbitUser, configs.RabbitPassword, configs.RabbitHost, configs.RabbitPort))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
