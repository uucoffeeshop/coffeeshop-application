package logger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type LogrusHandler struct {
	logger *logrus.Logger
}

func NewLogrusHandler(logger *logrus.Logger) *LogrusHandler {
	return &LogrusHandler{
		logger: logger,
	}
}

func ConvertLogLevel(level string) logrus.Level {
	var l logrus.Level

	switch strings.ToLower(level) {
	case "error":
		l = logrus.ErrorLevel
	case "warn":
		l = logrus.WarnLevel
	case "info":
		l = logrus.InfoLevel
	case "debug":
		l = logrus.DebugLevel
	default:
		l = logrus.InfoLevel
	}

	return l
}

func (h *LogrusHandler) Enabled(_ logrus.Level) bool {
	// support all logging levels
	return true
}

func (h *LogrusHandler) Handle(rec *logrus.Entry) error {
	fields := rec.Data

	entry := h.logger.WithFields(fields)

	switch rec.Level {
	case logrus.DebugLevel:
		entry.Debug(rec.Message)
	case logrus.InfoLevel:
		entry.Info(rec.Message)
	case logrus.WarnLevel:
		entry.Warn(rec.Message)
	case logrus.ErrorLevel:
		entry.Error(rec.Message)
	}

	return nil
}