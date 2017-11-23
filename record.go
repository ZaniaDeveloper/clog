package clog

import (
	"time"
)

type Fields map[string]string

type LogRecord struct {
	Timestamp time.Time
	Level     Level
	LogName   string
	Prefix    string
	Message   string
	Datas     Fields
}
