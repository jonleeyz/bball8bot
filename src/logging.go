package main

import (
	"fmt"
	"log"
)

// logInManualDebugMode is a syntatic wrapper around the log.Printf function that creates a log entry with debug tags.
func logInManualDebugMode(message string, debugObject ...interface{}) {
	if !isDebugLoggingEnabled {
		return
	}

	debugLog := fmt.Sprintf(message, debugObject...)
	log.Printf("[MANUAL_DEBUG_LOG] %s", debugLog)
}
