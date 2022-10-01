package strategy

import (
	"ht-matcher/model"
	"ht-matcher/util"
	"testing"
)

func equals(o1 model.SimpleOrder, o2 model.SimpleOrder) bool {
	return o1.Type == o2.Type &&
		o1.Quantity == o2.Quantity &&
		o1.Price.Number == o2.Price.Number &&
		o1.Status == o2.Status
}

func checkResult(t *testing.T, inputs []model.SimpleOrder, outputs []model.SimpleOrder) {
	inputs = Run(inputs, FOK)
	for i := 0; i < len(inputs); i++ {
		if !equals(inputs[i], outputs[i]) {
			t.Errorf("Not fit when sequence is %d", i+1)
		}
	}
}

func TestSingle1To1InTransaction(t *testing.T) {
	checkResult(
		t,
		util.GenerateSampleOrders(0),
		util.GenerateResultSampleOrders(0),
	)
}

func TestMultiple1To1InTransaction(t *testing.T) {
	checkResult(
		t,
		util.GenerateSampleOrders(1),
		util.GenerateResultSampleOrders(1),
	)
}

func TestSingle1ToNInTransaction(t *testing.T) {
	checkResult(
		t,
		util.GenerateSampleOrders(2),
		util.GenerateResultSampleOrders(2),
	)
}
