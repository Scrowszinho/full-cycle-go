package main

import (
	"database/sql"
	"teste/internal/entity"
	"teste/internal/events"
	"teste/internal/infra/database"
	"teste/internal/infra/web"
	"teste/internal/usecase"

	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcher = wire.NewSet(
	events.NewEventDispatcher,
	events.NewOrderCreated,
	wire.Bind(new(entity.EventInterface), new(*events.OrderCreated)),
	wire.Bind(new(entity.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	events.NewOrderCreated,
	wire.Bind(new(entity.EventInterface), new(*events.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
