package main

import (
	"ht-matcher/strategy"
	"ht-matcher/util"
)

func main() {
	inputs := append(util.GenerateSampleOrders1())
	strategy.Run(inputs, strategy.FOK)
}
