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
	log.Info("This is a Info: %d", x)
	log.Warn("This is a Warn: %d", x)
	log.Error("This is a Error: %d", x)

	// log.Fatal(fmt.Errorf("Me lene niko"))
	log.Fatal("This is a Fatal: %d", x)

}
