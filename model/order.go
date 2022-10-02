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

/*
*
After finish the workflow turn out it is not like Java to use struct as class(it will treat them are
different object in go). Thus, it is not a good idea to use embedded struct to simulate hierarchy of Java.
It seems need to use interface there will need more survey to fit go community usage or company principle.
Finally, it will looks like:
---------                  ---------------
| Order | <---interface--- | SimpleOrder |
---------        |         ---------------

	 |  	   --------------
	 |---------| OtherOrder |
			   --------------

todo: 1. survey interface examples, usages, and docuemnt in go
 2. or check what company real want
 3. try and error
*/
type RunnableOrder interface {
	getPrice() Price
	getType() OrderType
	getQuantity() int
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
