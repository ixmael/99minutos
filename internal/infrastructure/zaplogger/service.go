package zaplogger

import (
	"go.uber.org/zap"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type zaploggerservice struct {
	logger *zap.SugaredLogger
}

// NewZapLogger creates a new instance of the ZapLogger service.
func NewZapLogger(environment string) (ports.Logger, error) {
	var logger *zap.Logger
	if environment != "development" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	l := zaploggerservice{
		logger: logger.Sugar(),
	}

	return &l, nil
}

func (l *zaploggerservice) Debug(msg string, args ...any) {
	l.logger.Debug(msg)
}

func (l *zaploggerservice) Info(msg string, args ...any) {
	l.logger.Info(msg)
}

func (l *zaploggerservice) Warn(msg string, args ...any) {
	l.logger.Warn(msg)
}

func (l *zaploggerservice) Error(msg string, args ...any) {
	l.logger.Error(msg)
}

func (l *zaploggerservice) Fatal(msg string, args ...any) {
	l.logger.Fatal(msg)
}
