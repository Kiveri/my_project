package model

import "time"

type OrderStatus int
type DeliveryType int

const (
	CreatedOrderStatus          OrderStatus  = 1
	BuildingOrderStatus         OrderStatus  = 2
	SentOrderStatus             OrderStatus  = 3
	DeliveredOrderStatus        OrderStatus  = 4
	CourierDeliveryType         DeliveryType = 1
	SelfDeliveryDeliveryType    DeliveryType = 2
	PointOfDeliveryDeliveryType DeliveryType = 3
)

type Order struct {
	Id           int
	Positions    []Position
	CreatedAt    time.Time
	BuildingAt   *time.Time
	SentAt       *time.Time
	DeliveryAt   *time.Time
	OrderStatus  OrderStatus
	DeliveryType DeliveryType
}

func NewOrderCourierDeliveryType(positions []Position, now time.Time) *Order {
	return &Order{
		Positions:    positions,
		CreatedAt:    now,
		OrderStatus:  CreatedOrderStatus,
		DeliveryType: CourierDeliveryType,
	}
}

func NewOrderSelfDeliveryDeliveryType(positions []Position, now time.Time) *Order {
	return &Order{
		Positions:    positions,
		CreatedAt:    now,
		OrderStatus:  CreatedOrderStatus,
		DeliveryType: SelfDeliveryDeliveryType,
	}
}

func NewOrderPointOfDeliveryDeliveryType(positions []Position, now time.Time) *Order {
	return &Order{
		Positions:    positions,
		CreatedAt:    now,
		OrderStatus:  CreatedOrderStatus,
		DeliveryType: PointOfDeliveryDeliveryType,
	}
}

func (o *Order) Building(now time.Time) {
	o.OrderStatus = BuildingOrderStatus
	o.BuildingAt = &now
}

func (o *Order) Sent(now time.Time) {
	o.OrderStatus = SentOrderStatus
	o.SentAt = &now
}

func (o *Order) Delivered(now time.Time) {
	o.OrderStatus = DeliveredOrderStatus
	o.DeliveryAt = &now
}
