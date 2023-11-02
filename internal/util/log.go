package util

import (
	"strings"

	"go.uber.org/zap"
)

type Logger interface {
	Log(msg ...string)
	Warn(msg string, obj interface{})
	Error(msg string, err error)
	FatalError(msg string, err error)
}

type logger struct {
	zl *zap.Logger
}

func NewLogger(zl *zap.Logger) *logger {
	return &logger{zl}
}

func (l *logger) Log(msg ...string) {
	l.zl.Info(strings.Join(msg, " "))
}

func (l *logger) Warn(msg string, obj interface{}) {
	l.zl.Warn(msg,
		zap.Any("obj", obj),
	)
}

func (l *logger) Error(msg string, err error) {
	l.zl.Error(msg,
		zap.Error(err),
	)
}

func (l *logger) FatalError(msg string, err error) {
	l.zl.Fatal(msg,
		zap.Error(err),
	)
}
