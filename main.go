package main

import (
	"flag"
	"fmt"
	"ht-matcher/model"
	"ht-matcher/strategy"
	"ht-matcher/util"
)

/*
*
Concepts:
It should be able to take stream, Restful, file, or other protocol to provide windowed data as
a transaction. If data fulfill the window then start running whole process. Scaling will depends on way
how data transport into this. In my opinion, will chose using scaling the worker(put this cli running in
pod which served by k8s) to fit concurrency requirement and consistency of the same transaction.

Implementation:
In the example only implement 1 time running cli. There is more things to do below:

todo: 1. Add configuration for strategy, window size from cli
 2. Looping whole main function for always running and waiting for enough size of window of data
*/
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
