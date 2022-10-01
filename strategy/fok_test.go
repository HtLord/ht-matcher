package strategy

import (
	"ht-matcher/util"
	"testing"
)

func Test1To1(t *testing.T) {
	inputs := append(util.GenerateSampleOrders1())
	Run(inputs, FOK)
}

func Test1ToMany(t *testing.T) {
	inputs := append(util.GenerateSampleOrders1())
	Run(inputs, FOK)
}
