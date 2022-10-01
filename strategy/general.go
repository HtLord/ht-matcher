package strategy

import (
	"fmt"
	"ht-matcher/model"
)

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
