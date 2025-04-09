package loglevel

type LogLevel uint

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

func (level LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[level]
}

func (level LogLevel) EnumIndex() uint {
	return uint(level)
}
