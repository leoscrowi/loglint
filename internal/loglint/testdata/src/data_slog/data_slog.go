package data_slog

import (
	"context"
	"io"
	"log/slog"
)

func pkgSlogFunctions() {
	ctx := context.Background()

	h := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{})
	l := slog.New(h)

	slog.SetDefault(l)
	_ = slog.NewLogLogger(h, slog.LevelInfo)

	slog.Debug("запуск сервера")    // want "english check rule: запуск сервера"
	slog.Debug("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.Debug("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	slog.Debug("apikey")            // want "sensitive data not allowed: \"apikey\""

	slog.DebugContext(ctx, "запуск сервера")    // want "english check rule: запуск сервера"
	slog.DebugContext(ctx, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.DebugContext(ctx, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	slog.DebugContext(ctx, "apikey")            // want "sensitive data not allowed: \"apikey\""

	slog.Error("запуск сервера")    // want "english check rule: запуск сервера"
	slog.Error("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.Error("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	slog.Error("apikey")            // want "sensitive data not allowed: \"apikey\""

	slog.ErrorContext(ctx, "запуск сервера")    // want "english check rule: запуск сервера"
	slog.ErrorContext(ctx, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.ErrorContext(ctx, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	slog.ErrorContext(ctx, "apikey")            // want "sensitive data not allowed: \"apikey\""

	slog.Info("запуск сервера")    // want "english check rule: запуск сервера"
	slog.Info("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.Info("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	slog.Info("apikey")            // want "sensitive data not allowed: \"apikey\""

	slog.InfoContext(ctx, "запуск сервера")    // want "english check rule: запуск сервера"
	slog.InfoContext(ctx, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.InfoContext(ctx, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	slog.InfoContext(ctx, "apikey")            // want "sensitive data not allowed: \"apikey\""

	slog.Warn("запуск сервера")    // want "english check rule: запуск сервера"
	slog.Warn("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.Warn("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	slog.Warn("apikey")            // want "sensitive data not allowed: \"apikey\""

	slog.WarnContext(ctx, "запуск сервера")    // want "english check rule: запуск сервера"
	slog.WarnContext(ctx, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.WarnContext(ctx, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	slog.WarnContext(ctx, "apikey")            // want "sensitive data not allowed: \"apikey\""

	slog.Log(ctx, slog.LevelInfo, "запуск сервера")    // want "english check rule: запуск сервера"
	slog.Log(ctx, slog.LevelInfo, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.Log(ctx, slog.LevelInfo, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	slog.Log(ctx, slog.LevelInfo, "apikey")            // want "sensitive data not allowed: \"apikey\""

	slog.LogAttrs(ctx, slog.LevelInfo, "запуск сервера")    // want "english check rule: запуск сервера"
	slog.LogAttrs(ctx, slog.LevelInfo, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	slog.LogAttrs(ctx, slog.LevelInfo, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	slog.LogAttrs(ctx, slog.LevelInfo, "apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Debug("запуск сервера")    // want "english check rule: запуск сервера"
	l.Debug("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Debug("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Debug("apikey")            // want "sensitive data not allowed: \"apikey\""

	l.DebugContext(ctx, "запуск сервера")    // want "english check rule: запуск сервера"
	l.DebugContext(ctx, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.DebugContext(ctx, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.DebugContext(ctx, "apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Error("запуск сервера")    // want "english check rule: запуск сервера"
	l.Error("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Error("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Error("apikey")            // want "sensitive data not allowed: \"apikey\""

	l.ErrorContext(ctx, "запуск сервера")    // want "english check rule: запуск сервера"
	l.ErrorContext(ctx, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.ErrorContext(ctx, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.ErrorContext(ctx, "apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Info("запуск сервера")    // want "english check rule: запуск сервера"
	l.Info("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Info("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Info("apikey")            // want "sensitive data not allowed: \"apikey\""

	l.InfoContext(ctx, "запуск сервера")    // want "english check rule: запуск сервера"
	l.InfoContext(ctx, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.InfoContext(ctx, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.InfoContext(ctx, "apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Warn("запуск сервера")    // want "english check rule: запуск сервера"
	l.Warn("server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Warn("Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Warn("apikey")            // want "sensitive data not allowed: \"apikey\""

	l.WarnContext(ctx, "запуск сервера")    // want "english check rule: запуск сервера"
	l.WarnContext(ctx, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.WarnContext(ctx, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.WarnContext(ctx, "apikey")            // want "sensitive data not allowed: \"apikey\""

	l.Log(ctx, slog.LevelInfo, "запуск сервера")    // want "english check rule: запуск сервера"
	l.Log(ctx, slog.LevelInfo, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.Log(ctx, slog.LevelInfo, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.Log(ctx, slog.LevelInfo, "apikey")            // want "sensitive data not allowed: \"apikey\""

	l.LogAttrs(ctx, slog.LevelInfo, "запуск сервера")    // want "english check rule: запуск сервера"
	l.LogAttrs(ctx, slog.LevelInfo, "server started!🚀")  // want "special symbols are not allowed: !🚀 in \"server started!🚀\""
	l.LogAttrs(ctx, slog.LevelInfo, "Connection failed") // want "first uppercase letter is not allowed: Connection failed"
	l.LogAttrs(ctx, slog.LevelInfo, "apikey")            // want "sensitive data not allowed: \"apikey\""
}
