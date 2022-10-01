package strategy

import (
	"fmt"
	"ht-matcher/model"
)

const (
	recursiveLimit = 10
)

func doFOK(inputs []model.Order) []int {

	orders = inputs

	if orders == nil {
		return nil
	}

	testee := orders[0]
	idxes := recursive(1, testee.Type, testee.Price, testee.Quantity, []int{})
	if len(idxes) > 0 {
		idxes = append(idxes, 0)
	}

	return idxes
}

func recursive(i int, orderType model.OrderType, p float64, q int, memo []int) []int {
	if q == 0 {
		memo = append(memo, 0)
		return memo
	}
	if i == len(orders) {
		fmt.Println("Hit end of orders")
		return nil
	}

	var result1 []int
	var result2 []int

	tester := orders[i]

	if tester.Status != model.Neutral ||
		orderType == tester.Type ||
		p != tester.Price ||
		q < tester.Quantity {
		result1 = recursive(i+1, orderType, p, q, memo)
	} else {
		result2 = recursive(i+1, orderType, p, q-tester.Quantity, append(memo, i))
	}

	if result2 != nil {
		fmt.Printf("[%d] %v\n", i, result2)
		return result2
	}

	if result1 != nil {
		fmt.Printf("[%d] %v\n", i, result1)
		return result1
	}

	fmt.Println("None match")

	return nil
}
