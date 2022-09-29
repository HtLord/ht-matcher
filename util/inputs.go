package util

import (
	"ht-matcher/model"
	"time"
)

var (
	seq = 0
)

func nextSeq() int {
	seq += 1
	return seq
}

func GenerateSampleOrders() []model.Order {
	var result []model.Order

	result = append(
		result,
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.1,
			Status:   false,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.2,
			Status:   false,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 50,
			Price:    5.3,
			Status:   false,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Sell,
			Quantity: 10,
			Price:    5.1,
			Status:   false,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.2,
			Status:   false,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Sell,
			Quantity: 70,
			Price:    5.3,
			Status:   false,
		},
	)

	return result
}
