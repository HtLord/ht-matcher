package strategy

import (
	"fmt"
	"ht-matcher/model"
)

var (
	Debug  = false
	orders []model.SimpleOrder
)

type Strategy int

const (
	FOK Strategy = iota
)

/*
*
MarkFilled will modify order status from neutral to filled and direct effect on input
*/
func MarkFilled(inputs []model.SimpleOrder, idxes []int) {
	for _, idx := range idxes {
		inputs[idx].Status = model.FILLED
	}
}

/*
*
MarkKilled will modify order status from neutral to filled and return new slice as result

todo: Figure and fix why its' behavior not like MarkFilled
*/
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

/*
*
Display orders to console also it can be add customized message by adding variadic string
*/
func Display(inputs []model.SimpleOrder, customizedMessage ...string) {
	if customizedMessage != nil {
		fmt.Println(customizedMessage)
	} else {
		fmt.Println("[Display current round order(s)]")
	}
	for _, input := range inputs {
		fmt.Printf("%v\n", input)
	}
}

/**
SwapTargetToFirst swap item at target to the top of slice(at 0 position).

This is because of FOK will always treat first input as testee and find available combination
from rest items of slice
*/
func SwapTargetToFirst(target int, inputs []model.SimpleOrder) []model.SimpleOrder {
	newOrders := append(inputs)
	newOrders[0], newOrders[target] = newOrders[target], newOrders[0]
	return newOrders
}

/**
RecoverSequence remove the top of slice to the buttom

According to FOK will use SwapTargetToFirst. In the final round it will make latest item to top. For
more easy way to compare orders. Then, using this to recovery orders' order to fit original order.
*/
func RecoverSequence(inputs []model.SimpleOrder) []model.SimpleOrder {
	return append(inputs[1:], inputs[0])
}

/*
*
Run for matching orders according to selected strategy and price type of order(not implement others
yet except same price).
*/
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

		if Debug {
			Display(nextRoundInputs)
		}
	}

	inputs = MarkKilled(inputs)
	inputs = RecoverSequence(inputs)

	return inputs
}
