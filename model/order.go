package model

import "time"

type OrderType bool

const (
	Sell = true
	Buy  = false
)

type OrderStatus int

const (
	Neutral = iota
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
