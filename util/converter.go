/*
Convert csv will be cut into tuples by newline and returned as 2d-string and rest of converters
will take its' returned data for converting to models

Advance, more flexity for model converter, it may need ability to parse struct and auto convert from
string to target type by using reflect mechanism.
*/
package util

import (
	"encoding/csv"
	"fmt"
	"ht-matcher/model"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

func ConvertCsvToTuples(s string) [][]string {
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
	if len(data) != 2 {
		panic("Bad csv format for price")
	}

	// Type
	t := fmt.Sprint(data[0])

	// Number
	n, err := strconv.ParseFloat(fmt.Sprint(data[1]), 64)
	if err != nil {
		panic("[Converter] bad csv data for price")
	}

	return model.Price{
		Type:   model.PriceType(t),
		Number: n,
	}
}

func ConvertTupleToOrder(data []string) model.Order {
	if len(data) != 4 {
		panic("Bad csv format for order")
	}

	// Price
	p := ConvertTupleToPrice(data[:2])

	// Type
	t := fmt.Sprint(data[2])

	// Quantity
	q, err := strconv.ParseInt(fmt.Sprint(data[3]), 10, 32)
	if err != nil {
		panic("Bad csv format for int")
	}

	return model.Order{
		Price:    p,
		Type:     model.OrderType(t),
		Quantity: int(q),
	}
}

func ConvertTupleToSimpleOrder(data []string) model.SimpleOrder {
	if len(data) != 7 {
		panic("Bad csv format for order")
	}

	// Order
	o := ConvertTupleToOrder(data[:4])

	// Sequence
	sq, err := strconv.ParseInt(fmt.Sprint(data[4]), 10, 32)
	if err != nil {
		panic("Bad csv format for int")
	}

	// Status
	s := model.OrderStatus(fmt.Sprint(data[5]))

	// Time
	tms := fmt.Sprint(data[6])
	var tm time.Time
	if tms == "*" {
		tm = time.Now()
	} else {
		tm, err = time.Parse("2006-01-02 15:04", tms)
	}
	if err != nil {
		panic("Bad csv format for time")
	}

	return model.SimpleOrder{
		Order:    o,
		Sequence: int(sq),
		Status:   s,
		Time:     tm,
	}
}
