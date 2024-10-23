package loglevel

type LogLevel uint

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

func (level LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[level]
}

func (level LogLevel) EnumIndex() uint {
	return uint(level)
}
