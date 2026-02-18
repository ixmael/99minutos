package logger

import (
	"sync"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type InMemoryLoggerService struct {
	shimpmentLock sync.Mutex
	logger        []*string
}

// NewInMemoryLoggerServices creates a new instance of the Logger service.
func NewInMemoryLoggerServices() (ports.Logger, error) {
	l := InMemoryLoggerService{
		logger: make([]*string, 0),
	}

	return &l, nil
}

func (l *InMemoryLoggerService) Debug(msg string, args ...any) {
	l.logger = append(l.logger, &msg)
}

func (l *InMemoryLoggerService) Info(msg string, args ...any) {
	l.logger = append(l.logger, &msg)
}

func (l *InMemoryLoggerService) Warn(msg string, args ...any) {
	l.logger = append(l.logger, &msg)
}

func (l *InMemoryLoggerService) Error(msg string, args ...any) {
	l.logger = append(l.logger, &msg)
}

func (l *InMemoryLoggerService) Fatal(msg string, args ...any) {
	l.logger = append(l.logger, &msg)
}
