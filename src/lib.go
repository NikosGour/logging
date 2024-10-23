package logging

import (
	"fmt"
	"time"

	loglevel "github.com/NikosGour/logging/src/LogLevel"
	"gitlab.com/metakeule/fmtdate"
)

func print(level loglevel.LogLevel, format string, a ...any) (int, error) {
	msg := fmt.Sprintf(format, a...)
	var level_prefix string
	switch level {
	case loglevel.DEBUG:
		level_prefix = "🟩"
	case loglevel.INFO:
		level_prefix = "🟦"
	case loglevel.WARN:
		level_prefix = "🟨"
	case loglevel.ERROR:
		level_prefix = "🟥"
	}

	time := time.Now()
	millis := time.UnixMilli() % 1000

	time_formated := fmt.Sprintf("%s:%d", fmtdate.Format("DD-MM-YYYY hh:mm:ss", time), millis)

	return fmt.Printf("%s | %-5s | %v | %s\n", level_prefix, level, time_formated, msg)
}

func Debug(format string, a ...any) (int, error) {
	return print(loglevel.DEBUG, format, a...)
}
func Info(format string, a ...any) (int, error) {

	return print(loglevel.INFO, format, a...)
}
func Warn(format string, a ...any) (int, error) {
	return print(loglevel.WARN, format, a...)
}
func Error(format string, a ...any) (int, error) {

	return print(loglevel.ERROR, format, a...)
}
