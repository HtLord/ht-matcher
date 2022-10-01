package strategy

import (
	"fmt"
	"ht-matcher/model"
)

var (
	orders []model.SimpleOrder
)

type Strategy int

const (
	FOK   Strategy = iota
	debug          = false
)

func MarkFilled(inputs []model.SimpleOrder, idxes []int) {
	for _, idx := range idxes {
		inputs[idx].Status = model.FILLED
	}
}

func MarkKilled(inputs []model.SimpleOrder) []model.SimpleOrder {
	var newOrder []model.SimpleOrder

	for _, input := range inputs {
		if input.Status == model.NEUTRAL {
			input.Status = model.KILLED
		}
		newOrder = append(newOrder, input)
	}

	return newOrder
}

func Display(inputs []model.SimpleOrder) {
	fmt.Println("[Display current round order(s)]")
	for _, input := range inputs {
		fmt.Printf("%v\n", input)
	}
}

func SwapTargetToFirst(target int, inputs []model.SimpleOrder) []model.SimpleOrder {
	newOrders := append(inputs)
	newOrders[0], newOrders[target] = newOrders[target], newOrders[0]
	return newOrders
}

func RecoverSequence(inputs []model.SimpleOrder) []model.SimpleOrder {
	return append(inputs[1:], inputs[0])
}

func Run(inputs []model.SimpleOrder, strategy Strategy) []model.SimpleOrder {
	nextRoundInputs := append(inputs)

	for i := 0; i < len(inputs); i++ {
		nextRoundInputs = SwapTargetToFirst(i, inputs)

		var idxes []int
		switch strategy {
		default:
			idxes = doFOK(nextRoundInputs)
			if idxes != nil {
				MarkFilled(nextRoundInputs, idxes)
			}
		case FOK:
			idxes = doFOK(nextRoundInputs)
			if idxes != nil {
				MarkFilled(nextRoundInputs, idxes)
			}
		}

		if debug {
			Display(nextRoundInputs)
		}
	}

	inputs = MarkKilled(inputs)
	inputs = RecoverSequence(inputs)
	if debug {
		Display(inputs)
	}

	return inputs
}
