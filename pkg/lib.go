package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	loglevel "github.com/NikosGour/logging/pkg/LogLevel"
	"gitlab.com/metakeule/fmtdate"
)

var _LOGLEVEL loglevel.LogLevel = loglevel.DEBUG

func SetLogLevel(level loglevel.LogLevel) {
	_LOGLEVEL = level
}

func print(level loglevel.LogLevel, format string, a ...any) (int, error) {
	if _LOGLEVEL > level {
		return 0, nil
	}

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
	case loglevel.FATAL:
		level_prefix = "❌"
	}

	time := time.Now()
	millis := time.UnixMilli() % 1000

	time_formated := fmt.Sprintf("%s:%d", fmtdate.Format("DD-MM-YYYY hh:mm:ss", time), millis)

	_, file, line, ok := runtime.Caller(2)
	file = filepath.Base(file)

	if ok {
		return fmt.Printf("%s | %-5s | %v | %s:%d | %s\n", level_prefix, level, time_formated, file, line, msg)
	}

	return fmt.Printf("%s | %-5s | %v | %s\n", level_prefix, level, time_formated, msg)
}

func Debug(format string, a ...any) int {
	n, _ := print(loglevel.DEBUG, format, a...)
	return n
}
func Info(format string, a ...any) int {

	n, _ := print(loglevel.INFO, format, a...)
	return n
}
func Warn(format string, a ...any) int {
	n, _ := print(loglevel.WARN, format, a...)
	return n
}
func Error(format string, a ...any) int {

	n, _ := print(loglevel.ERROR, format, a...)
	return n
}

func Fatal(format string, a ...any) {

	_, _ = print(loglevel.FATAL, format, a...)
	os.Exit(1)
}
