package data_zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func pkgZapFunctions() {
	_ = zap.S()

	l := zap.NewNop()
	s := l.Sugar()

	l.DPanic("запуск сервера")    // want "english check rule: запуск сервера"
	l.DPanic("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.DPanic("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.DPanic("apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Debug("запуск сервера")    // want "english check rule: запуск сервера"
	l.Debug("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Debug("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Debug("apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Error("запуск сервера")    // want "english check rule: запуск сервера"
	l.Error("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Error("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Error("apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Fatal("запуск сервера")    // want "english check rule: запуск сервера"
	l.Fatal("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Fatal("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Fatal("apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Info("запуск сервера")    // want "english check rule: запуск сервера"
	l.Info("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Info("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Info("apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Log(zapcore.InfoLevel, "запуск сервера")    // want "english check rule: запуск сервера"
	l.Log(zapcore.InfoLevel, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Log(zapcore.InfoLevel, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Log(zapcore.InfoLevel, "apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Panic("запуск сервера")    // want "english check rule: запуск сервера"
	l.Panic("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Panic("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Panic("apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Warn("запуск сервера")    // want "english check rule: запуск сервера"
	l.Warn("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Warn("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Warn("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.DPanic("запуск сервера")    // want "english check rule: запуск сервера"
	s.DPanic("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.DPanic("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.DPanic("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.DPanicf("запуск сервера")    // want "english check rule: запуск сервера"
	s.DPanicf("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.DPanicf("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.DPanicf("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.DPanicln("запуск сервера")    // want "english check rule: запуск сервера"
	s.DPanicln("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.DPanicln("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.DPanicln("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.DPanicw("запуск сервера", "k", "v")    // want "english check rule: запуск сервера"
	s.DPanicw("server started!🚀", "k", "v")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.DPanicw("Connection failed", "k", "v") // want "first uppercase letter is not allowed: Connection failed"
	s.DPanicw("apikey", "k", "v")            // want "sensitive data not allowed: \"apikey\""

	s.Debug("запуск сервера")    // want "english check rule: запуск сервера"
	s.Debug("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Debug("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Debug("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Debugf("запуск сервера")    // want "english check rule: запуск сервера"
	s.Debugf("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Debugf("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Debugf("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Debugln("запуск сервера")    // want "english check rule: запуск сервера"
	s.Debugln("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Debugln("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Debugln("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Debugw("запуск сервера", "k", "v")    // want "english check rule: запуск сервера"
	s.Debugw("server started!🚀", "k", "v")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Debugw("Connection failed", "k", "v") // want "first uppercase letter is not allowed: Connection failed"
	s.Debugw("apikey", "k", "v")            // want "sensitive data not allowed: \"apikey\""

	s.Error("запуск сервера")    // want "english check rule: запуск сервера"
	s.Error("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Error("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Error("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Errorf("запуск сервера")    // want "english check rule: запуск сервера"
	s.Errorf("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Errorf("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Errorf("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Errorln("запуск сервера")    // want "english check rule: запуск сервера"
	s.Errorln("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Errorln("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Errorln("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Errorw("запуск сервера", "k", "v")    // want "english check rule: запуск сервера"
	s.Errorw("server started!🚀", "k", "v")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Errorw("Connection failed", "k", "v") // want "first uppercase letter is not allowed: Connection failed"
	s.Errorw("apikey", "k", "v")            // want "sensitive data not allowed: \"apikey\""

	s.Fatal("запуск сервера")    // want "english check rule: запуск сервера"
	s.Fatal("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Fatal("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Fatal("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Fatalf("запуск сервера")    // want "english check rule: запуск сервера"
	s.Fatalf("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Fatalf("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Fatalf("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Fatalln("запуск сервера")    // want "english check rule: запуск сервера"
	s.Fatalln("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Fatalln("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Fatalln("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Fatalw("запуск сервера", "k", "v")    // want "english check rule: запуск сервера"
	s.Fatalw("server started!🚀", "k", "v")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Fatalw("Connection failed", "k", "v") // want "first uppercase letter is not allowed: Connection failed"
	s.Fatalw("apikey", "k", "v")            // want "sensitive data not allowed: \"apikey\""

	s.Info("запуск сервера")    // want "english check rule: запуск сервера"
	s.Info("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Info("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Info("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Infof("запуск сервера")    // want "english check rule: запуск сервера"
	s.Infof("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Infof("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Infof("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Infoln("запуск сервера")    // want "english check rule: запуск сервера"
	s.Infoln("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Infoln("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Infoln("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Infow("запуск сервера", "k", "v")    // want "english check rule: запуск сервера"
	s.Infow("server started!🚀", "k", "v")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Infow("Connection failed", "k", "v") // want "first uppercase letter is not allowed: Connection failed"
	s.Infow("apikey", "k", "v")            // want "sensitive data not allowed: \"apikey\""

	s.Log(zapcore.InfoLevel, "запуск сервера")    // want "english check rule: запуск сервера"
	s.Log(zapcore.InfoLevel, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Log(zapcore.InfoLevel, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Log(zapcore.InfoLevel, "apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Logf(zapcore.InfoLevel, "запуск сервера")    // want "english check rule: запуск сервера"
	s.Logf(zapcore.InfoLevel, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Logf(zapcore.InfoLevel, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Logf(zapcore.InfoLevel, "apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Logln(zapcore.InfoLevel, "запуск сервера")    // want "english check rule: запуск сервера"
	s.Logln(zapcore.InfoLevel, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Logln(zapcore.InfoLevel, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Logln(zapcore.InfoLevel, "apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Logw(zapcore.InfoLevel, "запуск сервера", "k", "v")    // want "english check rule: запуск сервера"
	s.Logw(zapcore.InfoLevel, "server started!🚀", "k", "v")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Logw(zapcore.InfoLevel, "Connection failed", "k", "v") // want "first uppercase letter is not allowed: Connection failed"
	s.Logw(zapcore.InfoLevel, "apikey", "k", "v")            // want "sensitive data not allowed: \"apikey\""

	s.Panic("запуск сервера")    // want "english check rule: запуск сервера"
	s.Panic("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Panic("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Panic("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Panicf("запуск сервера")    // want "english check rule: запуск сервера"
	s.Panicf("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Panicf("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Panicf("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Panicln("запуск сервера")    // want "english check rule: запуск сервера"
	s.Panicln("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Panicln("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Panicln("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Panicw("запуск сервера", "k", "v")    // want "english check rule: запуск сервера"
	s.Panicw("server started!🚀", "k", "v")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Panicw("Connection failed", "k", "v") // want "first uppercase letter is not allowed: Connection failed"
	s.Panicw("apikey", "k", "v")            // want "sensitive data not allowed: \"apikey\""

	s.Warn("запуск сервера")    // want "english check rule: запуск сервера"
	s.Warn("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Warn("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Warn("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Warnf("запуск сервера")    // want "english check rule: запуск сервера"
	s.Warnf("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Warnf("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Warnf("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Warnln("запуск сервера")    // want "english check rule: запуск сервера"
	s.Warnln("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Warnln("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	s.Warnln("apikey")            // want "sensitive data not allowed: \"apikey\""

	s.Warnw("запуск сервера", "k", "v")    // want "english check rule: запуск сервера"
	s.Warnw("server started!🚀", "k", "v")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	s.Warnw("Connection failed", "k", "v") // want "first uppercase letter is not allowed: Connection failed"
	s.Warnw("apikey", "k", "v")            // want "sensitive data not allowed: \"apikey\""
}
