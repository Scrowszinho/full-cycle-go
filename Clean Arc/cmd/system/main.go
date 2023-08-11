package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"teste/configs"
	"teste/internal/events"
	"teste/internal/infra/graph"
	"teste/internal/infra/web/webserver"

	"github.com/streadway/amqp"
	"golang.org/x/tools/playground"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.DBPort)
	webserverHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webserverHandler.Create)
	fmt.Println("Starting web server")
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase)
	pb.RegisterOrderService(grpcServer, createOrderService)
	reflection.Register(grpcServer)
	fmt.Sprintln("Grpc server started")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCHost))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolvers{
		CreateOrderUseCase: *createOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL playground")
	http.ListenAndServe(":"+configs.GraphQLPort, nil)

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
