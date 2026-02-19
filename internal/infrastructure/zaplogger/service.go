package zaplogger

import (
	"errors"

	"go.uber.org/zap"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type zaploggerservice struct {
	logger *zap.SugaredLogger
}

// NewZapLogger creates a new instance of the ZapLogger service.
func NewZapLogger(environment string) (ports.Logger, error) {
	config := zap.NewProductionConfig()
	if environment != "production" {
		config = zap.NewDevelopmentConfig()
	}

	logger, err := config.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		return nil, errors.New("cannot create logger")
	}

	l := zaploggerservice{
		logger: logger.Sugar(),
	}

	return &l, nil
}

func (l *zaploggerservice) Debug(msg string, args ...any) {
	l.logger.Debugw(msg, args...)
}

func (l *zaploggerservice) Info(msg string, args ...any) {
	l.logger.Infow(msg, args...)
}

func (l *zaploggerservice) Warn(msg string, args ...any) {
	l.logger.Warnw(msg, args...)
}

func (l *zaploggerservice) Error(msg string, args ...any) {
	l.logger.Errorw(msg, args...)
}

func (l *zaploggerservice) Fatal(msg string, args ...any) {
	l.logger.Fatalw(msg, args...)
}

func (l *zaploggerservice) Stop() {
	l.logger.Sync()
}
