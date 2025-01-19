package utils

import (
	"log"
	"os"
)

func InitLogger(logLevel string) {
	var logFlags int
	switch logLevel {
	case "debug":
		logFlags = log.Ldate | log.Ltime | log.Lshortfile
	case "info":
		logFlags = log.Ldate | log.Ltime
	default:
		logFlags = log.Ldate | log.Ltime
	}

	log.SetFlags(logFlags)
	log.SetOutput(os.Stdout)
}
