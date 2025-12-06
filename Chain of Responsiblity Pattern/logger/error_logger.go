package logger

import "fmt"

type ErrorLogger struct {
	nextLogger Logger
}

func (er *ErrorLogger) SetNextLogger(next Logger) {
	er.nextLogger = next
}

func (er *ErrorLogger) LogMessage(level string, message string) {
	if level == "ERROR" {
		fmt.Println("[ERROR]:", message)
	} else if er.nextLogger != nil {
		er.nextLogger.LogMessage(level, message)
	}
}
