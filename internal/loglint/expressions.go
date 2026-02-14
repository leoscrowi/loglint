package loglint

var (
	LogPackages = []string{"log", "log/slog", "go.uber.org/zap"}

	DefaultKeyWords = []string{
		"apikey",
		"password",
		"token",
		"api_key",
		"api_token",
		"apitoken",
	}
)
