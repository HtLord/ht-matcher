package model

type PriceType string

const (
	SAME   PriceType = "same"
	MARKET PriceType = "market"
	LIMIT  PriceType = "limit"
)

type PriceMatcher interface {
	GetType() PriceType
	Matching(orders []Order) bool
}

type SamePriceMatcher interface {
	PriceMatcher
}

type Price struct {
	Type   PriceType
	Number float64
}
