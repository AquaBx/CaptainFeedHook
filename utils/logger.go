package utils

import (
    "github.com/rs/zerolog"
    "os"
)

var log zerolog.Logger

func InitLogger() {
    level := zerolog.InfoLevel
    if Flags.Debug {
        level = zerolog.DebugLevel
    }

    log = zerolog.New(os.Stdout).
        Level(level).
        With().
        Timestamp().
        Logger()
}

func Log(level string, msg string) {
    switch level {
    case "debug":
        log.Debug().Msg(msg)
    case "info":
        log.Info().Msg(msg)
    case "warn":
        log.Warn().Msg(msg)
    case "error":
        log.Error().Msg(msg)
    }
}

