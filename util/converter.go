package util

import (
	"encoding/csv"
	"io"
	"log"
	"strings"
)

func ConvertToRecords(s string) [][]string {
	var records [][]string
	r := csv.NewReader(strings.NewReader(s))
	r.Comma = '|'
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, record)
	}

	return records
}
