package strategy

import (
	"fmt"
	"ht-matcher/model"
)

var (
	orders []model.Order
)

type Strategy int

const (
	FOK Strategy = iota
)

type Options struct {
	DebugMode bool
}

func MarkFilled(inputs []model.Order, idxes []int) {
	for _, idx := range idxes {
		inputs[idx].Status = model.Filled
	}
}

func MarkKilled(inputs []model.Order) []model.Order {
	var newOrder []model.Order

	for _, input := range inputs {
		if input.Status == model.Neutral {
			input.Status = model.Killed
		}
		newOrder = append(newOrder, input)
	}

	return newOrder
}

func Display(inputs []model.Order) {
	fmt.Println("[Display current round order(s)]")
	for _, input := range inputs {
		fmt.Printf("%v\n", input)
	}
}

func SwapTargetToFirst(target int, inputs []model.Order) []model.Order {
	newOrders := append(inputs)
	newOrders[0], newOrders[target] = newOrders[target], newOrders[0]
	return newOrders
}

func RecoverSequence(inputs []model.Order) []model.Order {
	return append(inputs[1:], inputs[0])
}

func Run(inputs []model.Order, strategy Strategy) []model.Order {
	nextRoundInputs := append(inputs)

	for i := 0; i < len(inputs); i++ {
		nextRoundInputs = SwapTargetToFirst(i, inputs)
		if inputs[0].Status != model.Neutral {
			continue
		}

		var idxes []int
		switch strategy {
		default:
			idxes = doFOK(nextRoundInputs)
		case FOK:
			idxes = doFOK(nextRoundInputs)
		}

		if idxes != nil {
			MarkFilled(nextRoundInputs, idxes)
		}

		Display(nextRoundInputs)
	}

	inputs = MarkKilled(inputs)
	inputs = RecoverSequence(inputs)
	Display(inputs)

	return inputs
}
