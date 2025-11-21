package common

import (
	"log"
	"log/slog"
	"os"

	"github.com/UncleJunVIP/gabagool/pkg/gabagool"
)

func LogStandardFatal(msg string, err error) {
	log.SetOutput(os.Stderr)
	log.Fatalf("%s: %v", msg, err)
}

func GetLoggerInstance() *slog.Logger {
	return gabagool.GetLogger()
}
