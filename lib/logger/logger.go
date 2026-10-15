// Package logger
package logger

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type ILogger interface {
	Log(msg ...any)
	Warn(msg ...any)
	Error(msg ...any)
	Info(msg ...any)
	Debug(msg ...any)
	Dump(...any)
}

var _ ILogger = (*logger)(nil)

type logger struct {
	name string
	ctx  context.Context
}

// Dumb implements [ILogger].
func (l *logger) Dump(vals ...any) {
	g.Dump(vals)
}

var defaultLogger = DefaultLogger()

func Dump(msg ...any) {
	defaultLogger.Dump(msg...)
}

func Log(msg ...any) {
	defaultLogger.Log(msg)
}

func Debug(msg ...any) {
	defaultLogger.Debug(msg)
}

func Info(msg ...any) {
	defaultLogger.Info(msg)
}

func Error(msg ...any) {
	defaultLogger.Error(msg)
}

func New(name string) *logger {
	ctx := context.Background()
	return &logger{
		name,
		ctx,
	}
}

func NewWithContext(ctx context.Context, name string) *logger {
	return &logger{
		name,
		ctx,
	}
}

func DefaultLogger() *logger {
	name := "App"
	ctx := context.Background()
	return &logger{
		name,
		ctx,
	}
}

// Debug implements [ILogger].
func (l *logger) Debug(msg ...any) {
	glog.Debug(l.ctx, msg...)
}

// Error implements [ILogger].
func (l *logger) Error(msg ...any) {
	glog.Error(l.ctx, msg)
}

// Info implements [ILogger].
func (l *logger) Info(msg ...any) {
	glog.Info(l.ctx, msg)
}

// Log implements [ILogger].
func (l *logger) Log(msg ...any) {
	glog.Print(l.ctx, msg)
}

// Warn implements [ILogger].
func (l *logger) Warn(msg ...any) {
	glog.Warning(l.ctx, msg)
}
