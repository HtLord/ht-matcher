package model

import "time"

type OrderType bool

const (
	Sell OrderType = true
	Buy  OrderType = false
)

type OrderStatus int

const (
	Neutral OrderStatus = iota
	Filled
	Killed
)

type Order struct {
	Sequence int
	Time     time.Time
	Type     OrderType
	Quantity int
	Price    float64
	Status   OrderStatus
}
