package util

import "ht-matcher/model"

var inputDataSet = []string{
	`
same|5.1|buy|10|1|neutral|*
same|5.2|buy|10|2|neutral|*
same|5.3|buy|10|3|neutral|*
same|5.1|sell|10|4|neutral|*
same|5.1|buy|10|5|neutral|*
same|5.1|sell|70|6|neutral|*
`,
	`
same|5.1|buy|10|1|neutral|*
same|5.1|buy|10|2|neutral|*
same|5.1|buy|50|3|neutral|*
same|5.1|sell|10|4|neutral|*
same|5.2|buy|10|5|neutral|*
same|5.1|sell|10|6|neutral|*
`,
	`
same|5.2|buy|10|1|neutral|*
same|5.3|buy|10|2|neutral|*
same|5.1|buy|50|3|neutral|*
same|5.1|sell|10|4|neutral|*
same|5.2|buy|10|5|neutral|*
same|5.1|sell|40|6|neutral|*
`,
}

func GenerateSampleOrders(dataSetIdx int) []model.SimpleOrder {
	var orders []model.SimpleOrder

	for _, r := range ConvertCsvToTuples(inputDataSet[dataSetIdx]) {
		simpleOrder := ConvertTupleToSimpleOrder(r)
		orders = append(orders, simpleOrder)
	}
	return orders
}
