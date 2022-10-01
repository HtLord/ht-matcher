package util

import (
	"encoding/csv"
	"fmt"
	"ht-matcher/model"
	"io"
	"log"
	"strconv"
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

func ConvertTupleToPrice(data []string) model.Price {
	t := fmt.Sprint(data[0])

	n, err := strconv.ParseFloat(fmt.Sprint(data[1]), 64)
	if err != nil {
		panic("[Converter] bad csv data for price")
	}

	return model.Price{
		Type:   model.PriceType(t),
		Number: n,
	}
}
