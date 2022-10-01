package strategy

import (
	"fmt"
	"ht-matcher/model"
)

const (
	recursiveLimit = 10
)

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
