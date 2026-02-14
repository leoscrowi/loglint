package data_zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	NotEnglish     = "запуск сервера"
	SpecialSymbols = "server started!🚀"
	LowerCase      = "Connection failed"
	ApiKey         = "qwerty123"
)

func pkgZapFunctions() {
	_ = zap.S()

	l := zap.NewNop()
	s := l.Sugar()

	l.DPanic(NotEnglish)     // want "english check rule: запуск сервера"
	l.DPanic(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.DPanic(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.DPanic(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Debug(NotEnglish)     // want "english check rule: запуск сервера"
	l.Debug(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Debug(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Debug(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Error(NotEnglish)     // want "english check rule: запуск сервера"
	l.Error(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Error(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Error(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Fatal(NotEnglish)     // want "english check rule: запуск сервера"
	l.Fatal(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Fatal(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Fatal(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Info(NotEnglish)     // want "english check rule: запуск сервера"
	l.Info(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Info(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Info(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Log(zapcore.InfoLevel, NotEnglish)     // want "english check rule: запуск сервера"
	l.Log(zapcore.InfoLevel, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Log(zapcore.InfoLevel, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Log(zapcore.InfoLevel, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Panic(NotEnglish)     // want "english check rule: запуск сервера"
	l.Panic(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Panic(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Panic(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Warn(NotEnglish)     // want "english check rule: запуск сервера"
	l.Warn(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Warn(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Warn(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.DPanic(NotEnglish)     // want "english check rule: запуск сервера"
	s.DPanic(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.DPanic(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.DPanic(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.DPanicf(NotEnglish)     // want "english check rule: запуск сервера"
	s.DPanicf(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.DPanicf(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.DPanicf(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.DPanicln(NotEnglish)     // want "english check rule: запуск сервера"
	s.DPanicln(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.DPanicln(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.DPanicln(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.DPanicw(NotEnglish, "k", "v")     // want "english check rule: запуск сервера"
	s.DPanicw(SpecialSymbols, "k", "v") // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.DPanicw(LowerCase, "k", "v")      // want "first uppercase letter is not allowed: Connection failed"
	s.DPanicw(ApiKey, "k", "v")         // want "sensitive data not allowed: \"ApiKey\""

	s.Debug(NotEnglish)     // want "english check rule: запуск сервера"
	s.Debug(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Debug(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Debug(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Debugf(NotEnglish)     // want "english check rule: запуск сервера"
	s.Debugf(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Debugf(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Debugf(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Debugln(NotEnglish)     // want "english check rule: запуск сервера"
	s.Debugln(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Debugln(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Debugln(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Debugw(NotEnglish, "k", "v")     // want "english check rule: запуск сервера"
	s.Debugw(SpecialSymbols, "k", "v") // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Debugw(LowerCase, "k", "v")      // want "first uppercase letter is not allowed: Connection failed"
	s.Debugw(ApiKey, "k", "v")         // want "sensitive data not allowed: \"ApiKey\""

	s.Error(NotEnglish)     // want "english check rule: запуск сервера"
	s.Error(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Error(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Error(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Errorf(NotEnglish)     // want "english check rule: запуск сервера"
	s.Errorf(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Errorf(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Errorf(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Errorln(NotEnglish)     // want "english check rule: запуск сервера"
	s.Errorln(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Errorln(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Errorln(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Errorw(NotEnglish, "k", "v")     // want "english check rule: запуск сервера"
	s.Errorw(SpecialSymbols, "k", "v") // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Errorw(LowerCase, "k", "v")      // want "first uppercase letter is not allowed: Connection failed"
	s.Errorw(ApiKey, "k", "v")         // want "sensitive data not allowed: \"ApiKey\""

	s.Fatal(NotEnglish)     // want "english check rule: запуск сервера"
	s.Fatal(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Fatal(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Fatal(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Fatalf(NotEnglish)     // want "english check rule: запуск сервера"
	s.Fatalf(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Fatalf(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Fatalf(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Fatalln(NotEnglish)     // want "english check rule: запуск сервера"
	s.Fatalln(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Fatalln(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Fatalln(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Fatalw(NotEnglish, "k", "v")     // want "english check rule: запуск сервера"
	s.Fatalw(SpecialSymbols, "k", "v") // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Fatalw(LowerCase, "k", "v")      // want "first uppercase letter is not allowed: Connection failed"
	s.Fatalw(ApiKey, "k", "v")         // want "sensitive data not allowed: \"ApiKey\""

	s.Info(NotEnglish)     // want "english check rule: запуск сервера"
	s.Info(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Info(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Info(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Infof(NotEnglish)     // want "english check rule: запуск сервера"
	s.Infof(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Infof(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Infof(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Infoln(NotEnglish)     // want "english check rule: запуск сервера"
	s.Infoln(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Infoln(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Infoln(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Infow(NotEnglish, "k", "v")     // want "english check rule: запуск сервера"
	s.Infow(SpecialSymbols, "k", "v") // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Infow(LowerCase, "k", "v")      // want "first uppercase letter is not allowed: Connection failed"
	s.Infow(ApiKey, "k", "v")         // want "sensitive data not allowed: \"ApiKey\""

	s.Log(zapcore.InfoLevel, NotEnglish)     // want "english check rule: запуск сервера"
	s.Log(zapcore.InfoLevel, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Log(zapcore.InfoLevel, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Log(zapcore.InfoLevel, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Logf(zapcore.InfoLevel, NotEnglish)     // want "english check rule: запуск сервера"
	s.Logf(zapcore.InfoLevel, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Logf(zapcore.InfoLevel, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Logf(zapcore.InfoLevel, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Logln(zapcore.InfoLevel, NotEnglish)     // want "english check rule: запуск сервера"
	s.Logln(zapcore.InfoLevel, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Logln(zapcore.InfoLevel, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Logln(zapcore.InfoLevel, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Logw(zapcore.InfoLevel, NotEnglish, "k", "v")     // want "english check rule: запуск сервера"
	s.Logw(zapcore.InfoLevel, SpecialSymbols, "k", "v") // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Logw(zapcore.InfoLevel, LowerCase, "k", "v")      // want "first uppercase letter is not allowed: Connection failed"
	s.Logw(zapcore.InfoLevel, ApiKey, "k", "v")         // want "sensitive data not allowed: \"ApiKey\""

	s.Panic(NotEnglish)     // want "english check rule: запуск сервера"
	s.Panic(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Panic(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Panic(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Panicf(NotEnglish)     // want "english check rule: запуск сервера"
	s.Panicf(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Panicf(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Panicf(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Panicln(NotEnglish)     // want "english check rule: запуск сервера"
	s.Panicln(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Panicln(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Panicln(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Panicw(NotEnglish, "k", "v")     // want "english check rule: запуск сервера"
	s.Panicw(SpecialSymbols, "k", "v") // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Panicw(LowerCase, "k", "v")      // want "first uppercase letter is not allowed: Connection failed"
	s.Panicw(ApiKey, "k", "v")         // want "sensitive data not allowed: \"ApiKey\""

	s.Warn(NotEnglish)     // want "english check rule: запуск сервера"
	s.Warn(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Warn(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Warn(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Warnf(NotEnglish)     // want "english check rule: запуск сервера"
	s.Warnf(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Warnf(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Warnf(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Warnln(NotEnglish)     // want "english check rule: запуск сервера"
	s.Warnln(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Warnln(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	s.Warnln(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	s.Warnw(NotEnglish, "k", "v")     // want "english check rule: запуск сервера"
	s.Warnw(SpecialSymbols, "k", "v") // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Warnw(LowerCase, "k", "v")      // want "first uppercase letter is not allowed: Connection failed"
	s.Warnw(ApiKey, "k", "v")         // want "sensitive data not allowed: \"ApiKey\""
}
