package main

import (
	"ht-matcher/strategy"
	"ht-matcher/util"
)

func main() {
	strategy.FOK(
		util.GenerateSampleOrders(),
	)
	strategy.Reducer()
	strategy.DisplayOrders()
	// todo: logic
	// todo: reading input stream
	// todo: output input stream
	// todo: looping with pause and restart
}
