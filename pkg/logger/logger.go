package logger

import (
	"fmt"
	"log"
	"mux/pkg/colors"
	"os"
)

func ServerLogger() *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("%s%s", colors.CLIClear, colors.CLIPurple), log.LstdFlags)
}

func HandlersLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Llongfile|log.Lmsgprefix)
}
