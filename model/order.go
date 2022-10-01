package model

import "time"

type OrderType string

const (
	SELL OrderType = "sell"
	BUY            = "buy"
)

type OrderStatus string

const (
	NEUTRAL = "neutral"
	FILLED  = "filled"
	KILLED  = "killed"
)

type RunnableOrder interface {
	getPrice() Price
	getType() OrderType
	getQuantity() int
}

type Order struct {
	Price
	Type     OrderType
	Quantity int
}

type SimpleOrder struct {
	Order
	Sequence int
	Status   OrderStatus
	Time     time.Time
}

func (so SimpleOrder) getPrice() Price {
	return so.Price
}

func (so SimpleOrder) getType() OrderType {
	return so.Order.Type
}

func (so SimpleOrder) getQuantity() int {
	return so.Order.Quantity
}
