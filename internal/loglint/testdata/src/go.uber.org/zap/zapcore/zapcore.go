package zapcore

// Minimal stub of go.uber.org/zap/zapcore for analysistest (GOPATH mode).

type Level int8

const (
	DebugLevel  Level = -1
	InfoLevel   Level = 0
	WarnLevel   Level = 1
	ErrorLevel  Level = 2
	DPanicLevel Level = 3
	PanicLevel  Level = 4
	FatalLevel  Level = 5
)
