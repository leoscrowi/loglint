package data_slog

import (
	"context"
	"io"
	"log/slog"
)

const (
	NotEnglish     = "запуск сервера"
	SpecialSymbols = "server started!🚀"
	LowerCase      = "Connection failed"
	ApiKey         = "qwerty123"
)

func pkgSlogFunctions() {
	ctx := context.Background()

	h := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{})
	l := slog.New(h)

	slog.SetDefault(l)
	_ = slog.NewLogLogger(h, slog.LevelInfo)

	slog.Debug(NotEnglish)     // want "english check rule: запуск сервера"
	slog.Debug(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.Debug(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	slog.Debug(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	slog.DebugContext(ctx, NotEnglish)     // want "english check rule: запуск сервера"
	slog.DebugContext(ctx, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.DebugContext(ctx, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	slog.DebugContext(ctx, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	slog.Error(NotEnglish)     // want "english check rule: запуск сервера"
	slog.Error(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.Error(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	slog.Error(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	slog.ErrorContext(ctx, NotEnglish)     // want "english check rule: запуск сервера"
	slog.ErrorContext(ctx, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.ErrorContext(ctx, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	slog.ErrorContext(ctx, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	slog.Info(NotEnglish)     // want "english check rule: запуск сервера"
	slog.Info(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.Info(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	slog.Info(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	slog.InfoContext(ctx, NotEnglish)     // want "english check rule: запуск сервера"
	slog.InfoContext(ctx, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.InfoContext(ctx, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	slog.InfoContext(ctx, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	slog.Warn(NotEnglish)     // want "english check rule: запуск сервера"
	slog.Warn(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.Warn(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	slog.Warn(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	slog.WarnContext(ctx, NotEnglish)     // want "english check rule: запуск сервера"
	slog.WarnContext(ctx, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.WarnContext(ctx, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	slog.WarnContext(ctx, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	slog.Log(ctx, slog.LevelInfo, NotEnglish)     // want "english check rule: запуск сервера"
	slog.Log(ctx, slog.LevelInfo, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.Log(ctx, slog.LevelInfo, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	slog.Log(ctx, slog.LevelInfo, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	slog.LogAttrs(ctx, slog.LevelInfo, NotEnglish)     // want "english check rule: запуск сервера"
	slog.LogAttrs(ctx, slog.LevelInfo, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.LogAttrs(ctx, slog.LevelInfo, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	slog.LogAttrs(ctx, slog.LevelInfo, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Debug(NotEnglish)     // want "english check rule: запуск сервера"
	l.Debug(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Debug(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Debug(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.DebugContext(ctx, NotEnglish)     // want "english check rule: запуск сервера"
	l.DebugContext(ctx, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.DebugContext(ctx, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.DebugContext(ctx, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Error(NotEnglish)     // want "english check rule: запуск сервера"
	l.Error(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Error(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Error(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.ErrorContext(ctx, NotEnglish)     // want "english check rule: запуск сервера"
	l.ErrorContext(ctx, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.ErrorContext(ctx, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.ErrorContext(ctx, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Info(NotEnglish)     // want "english check rule: запуск сервера"
	l.Info(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Info(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Info(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.InfoContext(ctx, NotEnglish)     // want "english check rule: запуск сервера"
	l.InfoContext(ctx, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.InfoContext(ctx, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.InfoContext(ctx, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Warn(NotEnglish)     // want "english check rule: запуск сервера"
	l.Warn(SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Warn(LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Warn(ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.WarnContext(ctx, NotEnglish)     // want "english check rule: запуск сервера"
	l.WarnContext(ctx, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.WarnContext(ctx, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.WarnContext(ctx, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.Log(ctx, slog.LevelInfo, NotEnglish)     // want "english check rule: запуск сервера"
	l.Log(ctx, slog.LevelInfo, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Log(ctx, slog.LevelInfo, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.Log(ctx, slog.LevelInfo, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""

	l.LogAttrs(ctx, slog.LevelInfo, NotEnglish)     // want "english check rule: запуск сервера"
	l.LogAttrs(ctx, slog.LevelInfo, SpecialSymbols) // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.LogAttrs(ctx, slog.LevelInfo, LowerCase)      // want "first uppercase letter is not allowed: Connection failed"
	l.LogAttrs(ctx, slog.LevelInfo, ApiKey)         // want "sensitive data not allowed: \"ApiKey\""
}
