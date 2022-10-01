package util

import (
	"ht-matcher/model"
	"time"
)

func GenerateResultForSampleOrders1() []model.Order {
	var result []model.Order

	result = append(
		result,
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.1,
			Status:   model.Filled,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.2,
			Status:   model.Killed,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 50,
			Price:    5.3,
			Status:   model.Killed,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Sell,
			Quantity: 10,
			Price:    5.1,
			Status:   model.Filled,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.2,
			Status:   model.Killed,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Sell,
			Quantity: 70,
			Price:    5.3,
			Status:   model.Killed,
		},
	)

	return result
}

func GenerateResultForSampleOrders2() []model.Order {
	var result []model.Order

	result = append(
		result,
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.1,
			Status:   model.Filled,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.1,
			Status:   model.Filled,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 50,
			Price:    5.1,
			Status:   model.Killed,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Sell,
			Quantity: 10,
			Price:    5.1,
			Status:   model.Filled,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.2,
			Status:   model.Killed,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Sell,
			Quantity: 10,
			Price:    5.1,
			Status:   model.Filled,
		},
	)

	return result
}

func GenerateResultForSampleOrders3() []model.Order {
	var result []model.Order

	result = append(
		result,
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.2,
			Status:   model.Killed,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.3,
			Status:   model.Killed,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 50,
			Price:    5.1,
			Status:   model.Filled,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Sell,
			Quantity: 10,
			Price:    5.1,
			Status:   model.Filled,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Buy,
			Quantity: 10,
			Price:    5.2,
			Status:   model.Killed,
		},
		model.Order{
			Sequence: nextSeq(),
			Time:     time.Now(),
			Type:     model.Sell,
			Quantity: 40,
			Price:    5.1,
			Status:   model.Filled,
		},
	)

	return result
}
