package utils

import (
	"fmt"
)

type logger struct {
	colors map[string]string
	debug  bool
}

var sLogger logger

func InitLogger() {
	sLogger = logger{}

	sLogger.colors = map[string]string{
		"debug": "\033[93m",
		"error": "\033[0;31m",
		"warn":  "\033[93m",
		"info":  "\033[0m",
		"":      "\033[0m",
	}
}

func (f logger) getColor(c string) string {
	return f.colors[c]
}

func hasToLog(t string) bool {
	return Flags.Debug && t == "debug" || t != "debug"
}

func Log(t string, v string) {
	if hasToLog(t) {
		fmt.Printf("%s Logger %s : %s\n", sLogger.getColor(t), t, v)
	}
}
