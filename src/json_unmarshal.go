package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// getTelegramUpdateFromRecord unmarshals an SQS message and returns the contained Telegram update.
func getTelegramUpdateFromSQSMessage(sqsMessage events.SQSMessage) (*tgbotapi.Update, error) {
	var destination RecordBody

	logInManualDebugMode("SQS message body pre-unmarshal: %s", sqsMessage.Body)
	// slice off last char in record.Body; record.Body is invalid json due to extra lagging " char
	sqsMessageBody := sqsMessage.Body[:len(sqsMessage.Body)-1]
	logInManualDebugMode("SQS message body without quote pre-unmarshal: %s", sqsMessageBody)

	if err := json.Unmarshal([]byte(sqsMessageBody), &destination); err != nil {
		log.Printf("error when unmarshaling Telegram Update object: %v", err)
		return nil, err
	}
	logInManualDebugMode("Unmarshal destination post-unmarshal: %+v", destination)

	update := &destination.Body
	logInManualDebugMode("Update: %+v", update)
	return update, nil
}

// RecordBody is the unmarshal destination for an SQSEvent.Record.Body
type RecordBody struct {
	Method      string          `json:"method"`
	Body        tgbotapi.Update `json:"body-json"`
	QueryParams QueryParams     `json:"queryParams"`
	PathParams  PathParms       `json:"pathParams"`
}

type QueryParams struct {
}

type PathParms struct {
}
