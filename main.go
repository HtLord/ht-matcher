package main

import (
	"ht-matcher/model"
	"ht-matcher/strategy"
	"ht-matcher/util"
)

func main() {
	inputs := util.GenerateSampleOrders3()
	nextRoundInputs := append(inputs)

	for i := 0; i < len(inputs); i++ {
		nextRoundInputs = strategy.SwapTargetToFirst(i, inputs)
		if inputs[0].Status != model.Neutral {
			continue
		}

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
