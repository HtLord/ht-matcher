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
			Status:   model.Neutral,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.2,
			Status:   model.Neutral,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 50,
			Price:    5.3,
			Status:   model.Neutral,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Sell,
			Quantity: 10,
			Price:    5.1,
			Status:   model.Neutral,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.2,
			Status:   model.Neutral,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Sell,
			Quantity: 70,
			Price:    5.3,
			Status:   model.Neutral,
		},
	)

	return result
}
