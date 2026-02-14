package zap

// Minimal stub of go used by internal/loglint/testdata/src/data_zap/data_zap.go.

import "go.uber.org/zap/zapcore"

type Logger struct{}

type SugaredLogger struct{}

// S returns a global sugared logger in real zap.
// For tests, just return a stub.
func S() *SugaredLogger { return &SugaredLogger{} }

// NewNop returns a no-op logger in real zap.
func NewNop() *Logger { return &Logger{} }

func (l *Logger) Sugar() *SugaredLogger { return &SugaredLogger{} }

// Logger methods used by data_zap.go
func (l *Logger) DPanic(msg string, fields ...Field) {}
func (l *Logger) Debug(msg string, fields ...Field)  {}
func (l *Logger) Error(msg string, fields ...Field)  {}
func (l *Logger) Fatal(msg string, fields ...Field)  {}
func (l *Logger) Info(msg string, fields ...Field)   {}
func (l *Logger) Panic(msg string, fields ...Field)  {}
func (l *Logger) Warn(msg string, fields ...Field)   {}

func (l *Logger) Log(level zapcore.Level, msg string, fields ...Field) {}

// Field is part of zap’s API. Your test file doesn’t construct any,
// but zap method signatures accept them, so define a placeholder.
type Field struct{}

// SugaredLogger methods used by data_zap.go
func (s *SugaredLogger) DPanic(args ...interface{})                       {}
func (s *SugaredLogger) DPanicf(template string, args ...interface{})     {}
func (s *SugaredLogger) DPanicln(args ...interface{})                     {}
func (s *SugaredLogger) DPanicw(msg string, keysAndValues ...interface{}) {}

func (s *SugaredLogger) Debug(args ...interface{})                       {}
func (s *SugaredLogger) Debugf(template string, args ...interface{})     {}
func (s *SugaredLogger) Debugln(args ...interface{})                     {}
func (s *SugaredLogger) Debugw(msg string, keysAndValues ...interface{}) {}

func (s *SugaredLogger) Error(args ...interface{})                       {}
func (s *SugaredLogger) Errorf(template string, args ...interface{})     {}
func (s *SugaredLogger) Errorln(args ...interface{})                     {}
func (s *SugaredLogger) Errorw(msg string, keysAndValues ...interface{}) {}

func (s *SugaredLogger) Fatal(args ...interface{})                       {}
func (s *SugaredLogger) Fatalf(template string, args ...interface{})     {}
func (s *SugaredLogger) Fatalln(args ...interface{})                     {}
func (s *SugaredLogger) Fatalw(msg string, keysAndValues ...interface{}) {}

func (s *SugaredLogger) Info(args ...interface{})                       {}
func (s *SugaredLogger) Infof(template string, args ...interface{})     {}
func (s *SugaredLogger) Infoln(args ...interface{})                     {}
func (s *SugaredLogger) Infow(msg string, keysAndValues ...interface{}) {}

func (s *SugaredLogger) Log(level zapcore.Level, args ...interface{})                       {}
func (s *SugaredLogger) Logf(level zapcore.Level, template string, args ...interface{})     {}
func (s *SugaredLogger) Logln(level zapcore.Level, args ...interface{})                     {}
func (s *SugaredLogger) Logw(level zapcore.Level, msg string, keysAndValues ...interface{}) {}

func (s *SugaredLogger) Panic(args ...interface{})                       {}
func (s *SugaredLogger) Panicf(template string, args ...interface{})     {}
func (s *SugaredLogger) Panicln(args ...interface{})                     {}
func (s *SugaredLogger) Panicw(msg string, keysAndValues ...interface{}) {}

func (s *SugaredLogger) Warn(args ...interface{})                       {}
func (s *SugaredLogger) Warnf(template string, args ...interface{})     {}
func (s *SugaredLogger) Warnln(args ...interface{})                     {}
func (s *SugaredLogger) Warnw(msg string, keysAndValues ...interface{}) {}
