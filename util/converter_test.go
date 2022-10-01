package util

import (
	"fmt"
	"ht-matcher/model"
	"strconv"
	"testing"
)

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
	return ConvertToRecords(s)
}

func TestConvertToRecords(t *testing.T) {
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
		s += fmt.Sprintf(generateCsvFormatString(2), data...)
	}
	rs := ConvertToRecords(s)
	for _, r := range rs {
		p := ConvertTupleToPrice(r)
		n, err := strconv.ParseFloat(r[1], 64)
		if err != nil {
			t.Error("Bad test case need to fix test first")
		}
		if p.Type != model.PriceType(r[0]) ||
			p.Number != n {
			t.Errorf("Tuple %v is not equals to %v", p, r)
		}
	}
}
