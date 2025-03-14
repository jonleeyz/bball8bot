package logging

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// isDebugLoggingEnabled toggles debug logging on if true, and false otherwise. Is read from the respective Lambda env var.
var IS_DEBUG_LOGGING_ENABLED bool

func Init() {
	isDebugLoggingEnabledString, ok := os.LookupEnv("IS_DEBUG_LOGGING_ENABLED")
	Printf("[LAMBDA ENV VAR INIT] IS_DEBUG_LOGGING_ENABLED: %s; ok: %v", isDebugLoggingEnabledString, ok)

	if isDebugLoggingEnabledString == "true" && ok {
		IS_DEBUG_LOGGING_ENABLED = true
	} else {
		IS_DEBUG_LOGGING_ENABLED = false
	}
}

// Printf is a syntatic wrapper around the log.Printf function.
func Printf(message string, debugObjects ...interface{}) {
	log.Printf(message, debugObjects...)
}

// Fatal is a syntatic wrapper around the log.Fatal function.
func Fatal(errMessage string) {
	log.Fatal(errMessage)
}

// Fatalf is a syntatic wrapper around the log.Fatalf function.
func Fatalf(errMessage string, debugObjects ...interface{}) {
	log.Fatalf(errMessage, debugObjects...)
}

// Debugf is a syntatic wrapper around the log.Printf function that creates a log entry with debug tags.
func Debugf(message string, debugObjects ...interface{}) {
	if !IS_DEBUG_LOGGING_ENABLED {
		return
	}

	debugLog := fmt.Sprintf(message, debugObjects...)
	log.Printf("[MANUAL_DEBUG_LOG] %s", debugLog)
}

func LogUpdateObject(update tgbotapi.Update) {
	Printf("Update: %+v", update)
	// TODO: Maybe make these debug level logs
	Printf("Update - Message payload: %+v", update.Message)
	Printf("Update - EditedMessage payload: %+v", update.EditedMessage)
	Printf("Update - InlineQuery payload: %+v", update.InlineQuery)
	Printf("Update - ChosenInlineResult payload: %+v", update.ChosenInlineResult)
	Printf("Update - CallbackQuery payload: %+v", update.CallbackQuery)
	Printf("Update - Poll payload: %+v", update.Poll)
	Printf("Update - PollAnswer payload: %+v", update.PollAnswer)
}
