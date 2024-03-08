package main

import (
	"log"
	"os"
)

func readDebugLoggingFlag() {
	isDebugLoggingEnabledString, ok := os.LookupEnv("IS_DEBUG_LOGGING_ENABLED")
	log.Printf("isDebugLoggingEnabled: %s; ok: %v", isDebugLoggingEnabledString, ok)

	if isDebugLoggingEnabledString == "true" && ok {
		isDebugLoggingEnabled = true
	}
}
