package loglintplugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/leoscrowi/loglint/internal/loglint"
	"golang.org/x/tools/go/analysis"
)

// https://github.com/golangci/example-plugin-module-linter/blob/main/example.go
// https://golangci-lint.run/docs/plugins/module-plugins/

type MySettings struct {
	One string `json:"one"`
}

type PluginLogLint struct {
	settings MySettings
}

func New(conf any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[MySettings](conf)
	if err != nil {
		return nil, err
	}

	return &PluginLogLint{s}, nil
}

func (f *PluginLogLint) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func init() {
	register.Plugin("loglint", New)
}

func (f *PluginLogLint) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		loglint.Analyzer,
	}, nil
}
