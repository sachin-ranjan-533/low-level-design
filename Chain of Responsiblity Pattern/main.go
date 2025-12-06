package main

import "chain-of-responsibility/logger"

func main() {
	loggerChain := logger.InitializeLoggerChain()

	loggerChain.LogMessage("Info", "This is an informational message.")
	loggerChain.LogMessage("Warning", "This is a warning message.")
	loggerChain.LogMessage("ERROR", "This is an error message.")
}
