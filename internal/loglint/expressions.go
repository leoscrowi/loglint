package loglint

import (
	"github.com/leoscrowi/loglint/internal/rules"
	"github.com/leoscrowi/loglint/internal/rules/englishcheck"
	"github.com/leoscrowi/loglint/internal/rules/keywords"
	"github.com/leoscrowi/loglint/internal/rules/lowercase"
	"github.com/leoscrowi/loglint/internal/rules/specialsymbols"
)

var (
	logPackages = []string{"log", "log/slog", "go.uber.org/zap"}
	rulesMap    = map[string]rules.Rule{
		"englishcheck":   englishcheck.NewRule(),
		"specialsymbols": specialsymbols.NewRule(),
		"keywords":       keywords.NewRule(),
		"lowercase":      lowercase.NewRule(),
	}
)
