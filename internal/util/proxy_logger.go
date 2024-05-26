package util

import (
	"fmt"
	"log/slog"

	"github.com/elazarl/goproxy"
)

type proxyLogger struct {
	logger *slog.Logger
}

func (l *proxyLogger) Printf(format string, v ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, v...))
}

func NewProxyLogger(log *slog.Logger) goproxy.Logger {
	return &proxyLogger{
		logger: log,
	}
}
