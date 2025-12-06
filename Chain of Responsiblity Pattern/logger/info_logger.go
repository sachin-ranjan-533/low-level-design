package logger

import "fmt"

type InfoLogger struct {
	nextLogger Logger
}

func (il *InfoLogger) SetNextLogger(next Logger) {
	il.nextLogger = next
}

func (il *InfoLogger) LogMessage(level string, message string) {
	if level == "Info" {
		fmt.Println("[INFO]:", message)
	} else if il.nextLogger != nil {
		il.nextLogger.LogMessage(level, message)
	}
}
