package config

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLogger(env string) {
	zerolog.TimeFieldFormat = time.RFC3339

	if env == "dev" {
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "2006-01-02 15:04:05",
		}
		Logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	} else {
		Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
}
