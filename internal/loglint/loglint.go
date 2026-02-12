package loglint

import (
	"errors"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "reports log messages",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, errors.New("not implemented")
}
