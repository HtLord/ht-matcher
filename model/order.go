package model

import "time"

type OrderType int

const (
	Sell = iota
	Buy
)

type Order struct {
	Sequence int
	Time     time.Time
	Type     OrderType
	Quantity int
	Price    float64
}
