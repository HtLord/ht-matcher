package main

import (
	"flag"
	"fmt"
	"ht-matcher/model"
	"ht-matcher/strategy"
	"ht-matcher/util"
)

func main() {
	filePath := flag.String("file", "", "csv file path")
	flag.Parse()

	var origin []model.SimpleOrder
	if *filePath != "" {
		fmt.Println("Open csv file from <", *filePath, ">")
		origin = util.GenerateSampleOrders(0)
	} else {
		// Generate value from sample csv
		origin = util.GenerateSampleOrders(0)
	}
	strategy.Display(origin, "Origin Data")

	// Run strategy
	result := strategy.Run(origin, strategy.FOK)
	strategy.Display(result, "Final Result")
}
