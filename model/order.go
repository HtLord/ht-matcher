package model

import "time"

type OrderType bool

const (
	Sell = true
	Buy = false
)

type OrderStatus bool

const (
	Filled = true
	Killed = false
)

type Order struct {
	Sequence int
	Time     time.Time
	Type     OrderType
	Quantity int
	Price    float64
	Status   OrderStatus
}
