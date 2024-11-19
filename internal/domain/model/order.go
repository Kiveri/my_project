package model

import "time"

type OrderStatus int

const (
	CreatedOrderStatus   OrderStatus = 1
	BuildingOrderStatus  OrderStatus = 2
	SentOrderStatus      OrderStatus = 3
	DeliveredOrderStatus OrderStatus = 4
)

type Order struct {
	ID          int
	Positions   []Position
	CreatedAt   time.Time
	BuildingAt  *time.Time
	SentAt      *time.Time
	DeliveryAt  *time.Time
	OrderStatus OrderStatus
}

func NewOrder(positions []Position, now time.Time) *Order {
	return &Order{
		Positions:   positions,
		CreatedAt:   now,
		OrderStatus: CreatedOrderStatus,
	}
}

func (o *Order) Building(now time.Time) {
	o.OrderStatus = BuildingOrderStatus
	o.BuildingAt = &now
}
