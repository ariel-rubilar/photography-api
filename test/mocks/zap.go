package mocks

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewNoOpLogger() *zap.Logger {
	core := zapcore.NewNopCore()
	logger := zap.New(core)
	return logger
}
