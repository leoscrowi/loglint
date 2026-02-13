package main

import (
	//"os"

	"github.com/leoscrowi/loglint/internal/loglint"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(loglint.Analyzer)
}
