package clog

type Formatter func(*LogRecord) string
