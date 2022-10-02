package main

import (
	"flag"
	"fmt"
	"ht-matcher/model"
	"ht-matcher/strategy"
	"ht-matcher/util"
	"os"
	"time"
)

/*
*
Concepts:
It should be able to take stream, Restful, file, or other protocol to provide windowed data as
a transaction. If data fulfill the window then start running whole process. Scaling will depends on way
how data transport into this. In my opinion, will chose using scaling the worker(put this cli running in
pod which served by k8s) to fit concurrency requirement and consistency of the same transaction.

Implementation:
There is 2 running mode for this implementation.
- Mode 1 for running once
- Mode 2 while looping and waiting for csv file by given file path. After return the result, the file
will be remove. The server will keep listening till another file exist in the same path which be given
in cli.

todo: 1. Add configuration for strategy, window size from cli
 2. Looping whole main function for always running and waiting for enough size of window of data
*/
func doRunOnce(filePath string) {
	var origin []model.SimpleOrder
	if filePath != "" {
		fmt.Println("Open csv file from <", filePath, ">")
		origin = util.ReadCsvFromPath(filePath)
	} else {
		// Generate value from sample csv
		origin = util.GenerateSampleOrders(0)
	}
	strategy.Display(origin, "Origin Data")

	// Run strategy
	result := strategy.Run(origin, strategy.FOK)
	strategy.Display(result, "Final Result")
}

func doRunServer(filePath string) {
	for {
	ReadData:
		_, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("No data %v wait for a second\n", filePath)
			time.Sleep(1 * time.Second)
			goto ReadData
		}

		var origin []model.SimpleOrder
		origin = util.ReadCsvFromPath(filePath)
		strategy.Display(origin, "Origin Data")

		// Run strategy
		result := strategy.Run(origin, strategy.FOK)
		strategy.Display(result, "Final Result")
		util.RemoveFile(filePath)
	}
}

func main() {
	filePath := flag.String("f", "", "csv file path")
	runOnce := flag.Bool("once", true, "run once otherwise will listening the csv file")
	flag.Parse()

	if *runOnce {
		doRunOnce(*filePath)
	} else {
		doRunServer(*filePath)
	}

}
