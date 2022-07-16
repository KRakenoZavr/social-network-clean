package logger

import (
	"log"
	"mux/pkg/colors"
	"os"
)

func ServerLogger() *log.Logger {
	return log.New(os.Stdout, colors.CLIBlue, log.LstdFlags)
}

func HandlersLogger() *log.Logger {
	return log.New(os.Stdout, colors.CLIRed, log.LstdFlags|log.Llongfile)
}
