package loglintplugin

import (
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"github.com/leoscrowi/loglint/internal/loglint"
	"golang.org/x/tools/go/analysis"
)

// https://github.com/golangci/example-plugin-module-linter/blob/main/example.go
// https://golangci-lint.run/docs/plugins/module-plugins/

type MySettings struct {
	Rules    []string `json:"rules"`
	Keywords []string `json:"keywords"`
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
	if len(f.settings.Rules) > 0 {
		_ = loglint.Analyzer.Flags.Set("rules", strings.Join(f.settings.Rules, ","))
	}

	if len(f.settings.Keywords) > 0 {
		_ = loglint.Analyzer.Flags.Set("keywords", strings.Join(f.settings.Keywords, ","))
	}

	return []*analysis.Analyzer{
		loglint.Analyzer,
	}, nil
}
