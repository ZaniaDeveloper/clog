package clog

import (
	"io"
)

type Pusher func(LoggerInfos, *LogRecord)

type LoggerInfos interface {
	Config() Configuration
	Theme() *Theme
	Writer() io.Writer
	Path() string
	TimeFormat() string
	Layout() string
	Level() Level
	Rotation() Rotation
	RotationSize() Size
	RotationTimestamp() bool
	RotationPrefix() string
	RotationSuffix() string
	RotationKeepExtension() bool
}

type LoggerModifier interface {
	ThemeString(s string)
	ThemeObject(o *Theme)
	Path(p string)
	TimeFormat(f string)
	Layout(l string)
	Writer(w io.Writer)
	Stdout()
	Stderr()
	Level(l Level)
	Rotation(r *Rotation)
	RotationSize(s Size)
	RotationTimestamp(v bool)
	RotationPrefix(p string)
	RotationSuffix(s string)
	RotationKeepExtension(ke bool)
	Pusher(p Pusher)
}

type LogSender interface {
	Datas(fields ...string) LogSender
	WithDatas(fields Fields) LogSender

	Log(l Level, message string, args ...interface{})
	Trace(message string, args ...interface{})
	Debug(message string, args ...interface{})
	Msg(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

type Logger interface {
	LogSender
	IsNull() bool
	GetInfos() LoggerInfos
	Modify() LoggerModifier
}

func NewLogger(config *Configuration) Logger {
	return nil
}

func NewLoggerWithEngine(engine string, config *Configuration) Logger {
	return nil
}
