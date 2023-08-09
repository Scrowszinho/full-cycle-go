package usecase

import (
	"teste/internal/entity"
	"teste/internal/events"
)

type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderInputDTO   entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderRepository(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderInputDTO:   OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *CreateOrderUseCase) Execute(input OrderInputDTO) (OrderInputDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}

	order.CalculateFinalPrice()
	err := c.OrderRepository.Save(order); err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID: order.ID,
		Price: order.Price,
		Tax: order.Tax,
		FinalPrice: order.Price + order.Tax,
	}

	c.OrderCreated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderCreated)

	return dto, nil
}
