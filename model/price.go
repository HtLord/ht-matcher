package model

type PriceType int

const (
	SamePrice = iota
	MarketPrice
	LimitPrice
)

type PriceMatcher interface {
	GetType() PriceType
	Matching(orders []Order) bool
}

type SamePriceMatcher interface {
	PriceMatcher
}
