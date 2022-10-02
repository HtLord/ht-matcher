package util

import (
	"log"
	"os"
)

func RemoveFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Panicf("Fail to remove file %v", err)
	}
}
