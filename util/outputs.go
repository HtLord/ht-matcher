package util

import (
	"ht-matcher/model"
)

var onputDataSet = []string{
	`
same|5.1|buy|10|1|filled|*
same|5.2|buy|10|2|killed|*
same|5.3|buy|10|3|killed|*
same|5.1|sell|10|4|filled|*
same|5.1|buy|10|5|killed|*
same|5.1|sell|70|6|killed|*
`,
	`
same|5.1|buy|10|1|filled|*
same|5.1|buy|10|2|filled|*
same|5.1|buy|50|3|killed|*
same|5.1|sell|10|4|filled|*
same|5.2|buy|10|5|killed|*
same|5.1|sell|10|6|filled|*
`,
	`
same|5.2|buy|10|1|killed|*
same|5.3|buy|10|2|killed|*
same|5.1|buy|50|3|filled|*
same|5.1|sell|10|4|filled|*
same|5.2|buy|10|5|killed|*
same|5.1|sell|40|6|filled|*
`,
}

func GenerateResultSampleOrders(dataSetIdx int) []model.SimpleOrder {
	var orders []model.SimpleOrder

	for _, r := range ConvertCsvToTuples(onputDataSet[dataSetIdx]) {
		simpleOrder := ConvertTupleToSimpleOrder(r)
		orders = append(orders, simpleOrder)
	}
	return orders
}
