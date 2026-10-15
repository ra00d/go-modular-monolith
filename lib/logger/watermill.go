package logger

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
)

// watermillLogger adapts the glog-based logger to watermill.LoggerAdapter so
// messaging infrastructure logs land in the same pipeline as app logs.
type watermillLogger struct {
	fields watermill.LogFields
}

// NewWatermillLogger builds a watermill.LoggerAdapter backed by glog.
func NewWatermillLogger() watermill.LoggerAdapter {
	return &watermillLogger{}
}

func (l *watermillLogger) Error(msg string, err error, fields watermill.LogFields) {
	Error(fmt.Sprintf("watermill %s: %v %s", msg, err, l.render(l.merge(fields))))
}

func (l *watermillLogger) Info(msg string, fields watermill.LogFields) {
	Info(fmt.Sprintf("watermill %s %s", msg, l.render(l.merge(fields))))
}

func (l *watermillLogger) Debug(msg string, fields watermill.LogFields) {
	Debug(fmt.Sprintf("watermill %s %s", msg, l.render(l.merge(fields))))
}

func (l *watermillLogger) Trace(msg string, fields watermill.LogFields) {
	l.Debug(msg, fields)
}

// With returns a copy carrying additional fields.
func (l *watermillLogger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	cp := &watermillLogger{fields: l.merge(fields)}
	return cp
}

func (l *watermillLogger) merge(extra watermill.LogFields) watermill.LogFields {
	out := make(watermill.LogFields, len(l.fields)+len(extra))
	for k, v := range l.fields {
		out[k] = v
	}
	for k, v := range extra {
		out[k] = v
	}
	return out
}

func (l *watermillLogger) render(fields watermill.LogFields) string {
	if len(fields) == 0 {
		return ""
	}
	s := "{"
	first := true
	for k, v := range fields {
		if !first {
			s += " "
		}
		first = false
		s += fmt.Sprintf("%s=%v", k, v)
	}
	return s + "}"
}
