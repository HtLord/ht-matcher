package strategy

import (
	"fmt"
	"ht-matcher/model"
)

func Reducer() {
	var newOrders []model.Order
	for _, order := range orders {
		if order.Status != model.Filled {
			newOrders = append(newOrders, order)
		}
	}
	orders = newOrders
}

func DisplayOrders() {
	fmt.Println("[Rest order(s)]")
	for _, order := range orders {
		fmt.Printf("%v\n", order)
	}
}
