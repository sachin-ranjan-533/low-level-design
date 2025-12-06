package logger

type Logger interface {
	LogMessage(level string, message string)
}

func InitializeLoggerChain() Logger {
	infoLogger := &InfoLogger{}
	warningLogger := &WarningLogger{}
	errorLogger := &ErrorLogger{}

	infoLogger.SetNextLogger(warningLogger)
	warningLogger.SetNextLogger(errorLogger)

	return infoLogger
}
