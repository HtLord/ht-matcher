package strategy

import (
	"ht-matcher/model"
	"ht-matcher/util"
	"testing"
)

func equals(o1 model.Order, o2 model.Order) bool {
	return o1.Type == o2.Type &&
		o1.Quantity == o2.Quantity &&
		o1.Price == o2.Price &&
		o1.Status == o2.Status
}

func checkResult(t *testing.T, inputs []model.Order, outputs []model.Order) {
	inputs = Run(inputs, FOK)
	for i := 0; i < len(inputs); i++ {
		if !equals(inputs[i], outputs[i]) {
			t.Errorf("Not fit when sequence is %d", i)
		}
	}
}

func TestSingle1To1InTransaction(t *testing.T) {
	checkResult(
		t,
		util.GenerateSampleOrders1(),
		util.GenerateResultForSampleOrders1(),
	)
}

func TestMultiple1To1InTransaction(t *testing.T) {
	checkResult(
		t,
		util.GenerateSampleOrders2(),
		util.GenerateResultForSampleOrders2(),
	)
}

func TestSingle1ToNInTransaction(t *testing.T) {
	checkResult(
		t,
		util.GenerateSampleOrders3(),
		util.GenerateResultForSampleOrders3(),
	)
}
