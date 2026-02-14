package loglint

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, Analyzer,
		"data_log",
		"data_slog",
		"data_zap",
	)
}
