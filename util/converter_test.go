package util

import (
	"fmt"
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

func TestConvertToRecords(t *testing.T) {
	datas := [][]interface{}{
		[]interface{}{"same", "1"},
		[]interface{}{"limit", "1.2"},
		[]interface{}{"market", "12.3"},
	}
	var s string
	for _, data := range datas {
		s += fmt.Sprintf(generateCsvFormatString(2), data...)
	}
	fmt.Printf(s)
	rs := ConvertToRecords(s)
	for i := range rs {
		equalPriceCsv(
			t,
			rs[i],
			datas[i][0],
			datas[i][1],
		)
	}
}
