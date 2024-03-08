package json

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"github.com/jonleeyz/bbball8bot/internal/logging"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// GetTelegramUpdateFromRecord unmarshals an SQS message and returns the contained Telegram update.
func GetTelegramUpdateFromSQSMessage(sqsMessage events.SQSMessage) (*tgbotapi.Update, error) {
	var destination RecordBody

	logging.Debugf("SQS message body pre-unmarshal: %s", sqsMessage.Body)
	// slice off last char in record.Body; record.Body is invalid json due to extra lagging " char
	sqsMessageBody := sqsMessage.Body[:len(sqsMessage.Body)-1]
	logging.Debugf("SQS message body without quote pre-unmarshal: %s", sqsMessageBody)

	if err := json.Unmarshal([]byte(sqsMessageBody), &destination); err != nil {
		logging.Printf("error when unmarshaling Telegram Update object: %v", err)
		return nil, err
	}
	logging.Debugf("Unmarshal destination post-unmarshal: %+v", destination)

	update := &destination.Body
	logging.Debugf("Update: %+v", update)
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
