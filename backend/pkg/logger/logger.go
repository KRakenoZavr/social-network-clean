package logger

import (
	"fmt"
	"log"
	"os"

	"mux/pkg/colors"
)

func ServerLogger() *log.Logger {
	runEnv := os.Getenv("RUN_ENV")

	if runEnv == "test" {
		return log.New(os.Stdout, colors.CLIPurple, log.LstdFlags)
	}

	return log.New(os.Stdout, fmt.Sprintf("%s%s", colors.CLIClear, colors.CLIPurple), log.LstdFlags)
}

func HandlersLogger() *log.Logger {
	runEnv := os.Getenv("RUN_ENV")

	if runEnv == "test" {
		return log.New(os.Stdout, colors.CLIOrange, log.LstdFlags|log.Llongfile)
	}

	return log.New(os.Stdout, "", log.LstdFlags|log.Llongfile|log.Lmsgprefix)
}
