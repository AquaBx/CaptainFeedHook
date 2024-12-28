package utils

import "flag"

var Flags flags

type flags struct {
	Debug bool
}

func InitFlags() {
	debug := flag.Bool("debug", false, "Debug flag")
	flag.Parse()
	Flags = flags{Debug: *debug}
}
