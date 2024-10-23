package main

import (
	"fmt"

	log "github.com/NikosGour/logging/src"
	loglevel "github.com/NikosGour/logging/src/LogLevel"
)

func main() {
	fmt.Println(loglevel.DEBUG, loglevel.INFO, loglevel.WARN, loglevel.ERROR)
	fmt.Println(loglevel.DEBUG.EnumIndex(), loglevel.INFO.EnumIndex(), loglevel.WARN.EnumIndex(), loglevel.ERROR.EnumIndex())

	x := 10
	log.Debug("This is a Debug: %d", x)
	log.Info("This is a Info: %v", x)
	log.Warn("This is a Warn: %v", x)
	log.Error("This is a Error: %v", x)

}
