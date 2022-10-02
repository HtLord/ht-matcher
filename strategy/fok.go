package strategy

import (
	"fmt"
	"ht-matcher/model"
)

/*
*
doFOK is implementation of FOK business logic by DFS for fitting FIFO.

todo: 1. Modify recursive style to while-loop style(will remove global constant dependencies)
 2. Enhance logging strategy and message
 3. Add abstracted price matching logic
 4. Survey, figure out, and fix why enum not working(items which is model.OrderStatus can not be
    used in here
*/
func doFOK(inputs []model.SimpleOrder) []int {

	orders = inputs

	if orders == nil {
		return nil
	}

	testee := orders[0]
	var idxes []int
	if testee.Status == "neutral" {
		idxes = recursive(1, testee.Type, testee.Price, testee.Quantity, []int{})
	}

	return idxes
}

func recursive(i int, orderType model.OrderType, p model.Price, q int, memo []int) []int {
	if q == 0 {
		memo = append(memo, 0)
		return memo
	}
	if i == len(orders) {
		if Debug {
			fmt.Println("Hit end of orders")
		}
		return nil
	}

	var result1 []int
	var result2 []int

	tester := orders[i]

	if tester.Status != "neutral" ||
		orderType == tester.Order.Type ||
		p.Number != tester.Order.Price.Number ||
		q < tester.Quantity {
		result1 = recursive(i+1, orderType, p, q, memo)
	} else {
		result2 = recursive(i+1, orderType, p, q-tester.Quantity, append(memo, i))
	}

	if result2 != nil {
		return result2
	}

	if result1 != nil {
		return result1
	}

	if Debug {
		fmt.Printf("Result 1 (lo priority)[%d] %v\n", i, result1)
		fmt.Printf("Result 2 (hi priority)[%d] %v\n", i, result2)
		fmt.Println("None match")
	}

	return nil
}
