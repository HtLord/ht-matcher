package main

import (
	"ht-matcher/strategy"
	"ht-matcher/util"
)

func main() {
	// Generate value from sample csv
	origin := util.GenerateSampleOrders(0)
	strategy.Display(origin, "Origin Data")

	// Run strategy
	result := strategy.Run(origin, strategy.FOK)
	strategy.Display(result, "Final Result")
}
