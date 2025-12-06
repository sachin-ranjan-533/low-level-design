package logger

import "fmt"

type WarningLogger struct {
	nextLogger Logger
}

func (wl *WarningLogger) SetNextLogger(next Logger) {
	wl.nextLogger = next
}

func (wl *WarningLogger) LogMessage(level string, message string) {
	if level == "Warning" {
		fmt.Println("[WARNING]:", message)
	} else if wl.nextLogger != nil {
		wl.nextLogger.LogMessage(level, message)
	}
}
