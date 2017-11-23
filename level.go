package clog

type Level int

func (l Level) String() string {
	res, ok := levelNames[l]
	if !ok {
		return "INVALID"
	}

	return res
}

func (l Level) IsValid() bool {
	_, ok := levelNames[l]

	return ok
}

func (l Level) IsEnabled() bool {
	return l.IsValid() && (l != DISABLED)
}

const (
	DISABLED Level = iota
	TRACE
	DEBUG
	MESSAGE
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	levelNames = map[Level]string{
		DISABLED: "DISABLED",
		TRACE:    "TRACE",
		DEBUG:    "DEBUG",
		MESSAGE:  "MESSAGE",
		INFO:     "INFO",
		WARNING:  "WARNING",
		ERROR:    "ERROR",
		FATAL:    "FATAL",
	}
)
