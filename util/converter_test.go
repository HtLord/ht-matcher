package util

import (
	"fmt"
	"ht-matcher/model"
	"strconv"
	"testing"
	"time"
)

/**
Test Scope looks like
- Setup data as tuples to simulate csv and model
- Convert csv to model
- Convert tuple to field of model
- Validate result of previous 2 steps
Currently will just copy-and-paste coding style for convenient adjust to better style latter.
*/

func equalPriceCsv(t *testing.T, record []string, results ...any) {
	if len(record) != len(results) {
		t.Error("Reading record fail cause length is not equal")
	}
	for i := 0; i < len(record); i++ {
		if record[i] != results[i] {
			t.Errorf(
				"Reading record data at column %d, \"%s\" is not eqaul \"%s\"",
				i,
				record[i],
				results[i],
			)
		}
	}
}

func generateCsvFormatString(number int) string {
	s := ""
	s += "%s"
	for i := 1; i < number; i++ {
		s += "|%s"
	}
	s += "\n"
	return s
}

func generateRecords(datas [][]interface{}, debug bool) [][]string {
	var s string
	for _, data := range datas {
		s += fmt.Sprintf(generateCsvFormatString(2), data...)
	}
	if debug {
		fmt.Printf(s)
	}
	return ConvertCsvToTuples(s)
}

func TestConvertCsvToTuples(t *testing.T) {
	datas := [][]interface{}{
		[]interface{}{"same", "1"},
		[]interface{}{"limit", "1.2"},
		[]interface{}{"market", "12.3"},
	}
	rs := generateRecords(datas, false)
	for i := range rs {
		equalPriceCsv(
			t,
			rs[i],
			datas[i][0],
			datas[i][1],
		)
	}
}

func TestConvertCsvToPrice(t *testing.T) {
	datas := [][]interface{}{
		[]interface{}{"same", "1"},
		[]interface{}{"limit", "1.2"},
		[]interface{}{"market", "12.3"},
	}
	s := ""
	for _, data := range datas {
		s += fmt.Sprintf(generateCsvFormatString(len(datas[0])), data...)
	}
	rs := ConvertCsvToTuples(s)
	for _, r := range rs {
		// Converting tuple to model
		theModel := ConvertTupleToPrice(r)
		// Converting tuple to type of field of model
		n, err := strconv.ParseFloat(r[1], 64)
		if err != nil {
			t.Error("Bad test case need to fix test first")
		}
		// Validating
		if theModel.Type != model.PriceType(r[0]) ||
			theModel.Number != n {
			t.Errorf("Tuple %v is not equals to %v", theModel, r)
		}
	}
}

func TestConvertCsvToOrder(t *testing.T) {
	datas := [][]interface{}{
		[]interface{}{"same", "1", "buy", "12"},
		[]interface{}{"limit", "1.2", "sell", "2"},
		[]interface{}{"market", "12.3", "buy", "1"},
	}
	s := ""
	for _, data := range datas {
		s += fmt.Sprintf(generateCsvFormatString(len(datas[0])), data...)
	}
	rs := ConvertCsvToTuples(s)
	for _, r := range rs {
		// Converting tuple to model
		theModel := ConvertTupleToOrder(r)
		// Converting tuple to type of field of model
		pt := model.PriceType(fmt.Sprintf(r[0]))
		pn, err := strconv.ParseFloat(fmt.Sprint(r[1]), 64)
		if err != nil {
			t.Errorf("Bad csv for order.price.number %v", r)
		}
		ot := model.OrderType(fmt.Sprintf(r[2]))
		q, err := strconv.ParseInt(fmt.Sprint(r[3]), 32, 32)
		if err != nil {
			t.Errorf("Bad csv for order.quantity %v", r)
		}
		// Validating
		if theModel.Price.Type != pt ||
			theModel.Price.Number != pn ||
			theModel.Type != ot ||
			theModel.Quantity != int(q) {
			t.Errorf("Tuple %v is not equals to %v", theModel, r)
		}
	}
}

func TestConvertCsvToSimpleOrder(t *testing.T) {
	datas := [][]interface{}{
		[]interface{}{"same", "1", "buy", "12", "1", "neutral", "*"},
		[]interface{}{"limit", "1.2", "sell", "2", "2", "filled", "2022-01-01 15:15"},
		[]interface{}{"market", "12.3", "buy", "1", "3", "killed", "*"},
	}
	s := ""
	for _, data := range datas {
		s += fmt.Sprintf(generateCsvFormatString(len(datas[0])), data...)
	}
	rs := ConvertCsvToTuples(s)
	for _, r := range rs {
		// Converting tuple to model
		theModel := ConvertTupleToSimpleOrder(r)
		// Converting tuple to type of field of model
		pt := model.PriceType(fmt.Sprintf(r[0]))
		pn, err := strconv.ParseFloat(fmt.Sprint(r[1]), 64)
		if err != nil {
			t.Errorf("Bad csv for order.price.number %v", r)
		}
		ot := model.OrderType(fmt.Sprintf(r[2]))
		q, err := strconv.ParseInt(fmt.Sprint(r[3]), 32, 32)
		if err != nil {
			t.Errorf("Bad csv for order.quantity %v", r)
		}
		sq, err := strconv.ParseInt(fmt.Sprint(r[4]), 32, 32)
		st := model.OrderStatus(fmt.Sprint(r[5]))

		tms := fmt.Sprintf(r[6])
		var tm time.Time
		if tms != "*" {
			tm, err = time.Parse("2006-01-02 15:04", tms)
		}
		if err != nil {
			t.Errorf("Bad csv for order.time %v", r)
		}
		// Validating
		if tms != "*" &&
			theModel.Time != tm {
			t.Errorf("Tuple %v is not equals to %v", theModel, r)
		}
		if theModel.Order.Price.Type != pt ||
			theModel.Order.Price.Number != pn ||
			theModel.Order.Type != ot ||
			theModel.Order.Quantity != int(q) ||
			theModel.Sequence != int(sq) ||
			theModel.Status != st {
			t.Errorf("Tuple %v is not equals to %v", theModel, r)
		}
	}
}
