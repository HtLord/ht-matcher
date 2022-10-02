/**
Generate input by csv for testing or real workflow by loading data from csv file
*/
package util

import (
	"encoding/csv"
	"fmt"
	"ht-matcher/model"
	"io"
	"log"
	"os"
)

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

/**
ReadCsvFromPath will reading csv and convert it into simple orders. The sequence of csv must fit model's
order from top to buttom.
e.g.
type A struct {
	a1 string
	B
}
type B struct {
	b1 string
    b2 string
}
---
a1|b1|b2
*/
func ReadCsvFromPath(filePath string) []model.SimpleOrder {
	var orders []model.SimpleOrder

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	var csvstring string
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		tmp := fmt.Sprintf("%s\n", rec)
		csvstring = csvstring + tmp[1:len(tmp)-2] + "\n"
	}

	rs := ConvertCsvToTuples(csvstring)
	for _, r := range rs {
		so := ConvertTupleToSimpleOrder(r)
		orders = append(orders, so)
	}

	return orders
}
