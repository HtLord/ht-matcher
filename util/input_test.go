package util

import (
	"os"
	"strings"
	"testing"
)

func TestReadCsvFromPath(t *testing.T) {
	pwd, _ := os.Getwd()
	fs := strings.Split(pwd, "/")
	fs = append(fs[:len(fs)-1], "data", "test.csv")
	p := strings.Join(fs, "/")

	ReadCsvFromPath(p)
}
