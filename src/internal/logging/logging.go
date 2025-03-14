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
	log.Printf("[LAMBDA ENV VAR INIT] IS_DEBUG_LOGGING_ENABLED: %s; ok: %v", isDebugLoggingEnabledString, ok)

	if isDebugLoggingEnabledString == "true" && ok {
		IS_DEBUG_LOGGING_ENABLED = true
	} else {
		IS_DEBUG_LOGGING_ENABLED = false
	}
}

// Infof is a syntatic wrapper around the log.Printf function.
func Infof(message string, debugObjects ...interface{}) {
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

// Errorf is a syntatic wrapper around the log.Printf function that creates a log entry with error tags.
func Errorf(errMessage string, debugObjects ...interface{}) {
	errorLog := fmt.Errorf(errMessage, debugObjects...)
	log.Printf("[ERROR] %s", errorLog)
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
	Infof("Update: %+v", update)
	// TODO: Maybe make these debug level logs
	Debugf("Update - Message payload: %+v", update.Message)
	Debugf("Update - EditedMessage payload: %+v", update.EditedMessage)
	Debugf("Update - InlineQuery payload: %+v", update.InlineQuery)
	Debugf("Update - ChosenInlineResult payload: %+v", update.ChosenInlineResult)
	Debugf("Update - CallbackQuery payload: %+v", update.CallbackQuery)
	Debugf("Update - Poll payload: %+v", update.Poll)
	Debugf("Update - PollAnswer payload: %+v", update.PollAnswer)
}
