package main

import (
	"ht-matcher/strategy"
	"ht-matcher/util"
)

func main() {
	inputs := util.GenerateSampleOrders()
	nextRoundInputs := append(inputs)

	for i := 0; i < len(inputs); i++ {
		nextRoundInputs = strategy.SwapTargetToFirst(i, inputs)
		idxes := strategy.FOK(nextRoundInputs)

		if idxes != nil {
			strategy.MarkFilled(nextRoundInputs, idxes)
		}

		strategy.Display(nextRoundInputs)
	}

	inputs = strategy.MarkKilled(inputs)
	inputs = strategy.RecoverSequence(inputs)
	strategy.Display(inputs)
}
