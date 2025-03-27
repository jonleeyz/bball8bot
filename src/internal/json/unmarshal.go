package json

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"github.com/jonleeyz/bbball8bot/internal/logging"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// GetTelegramUpdateFromRecord unmarshals an SQS message and returns the contained Telegram update.
func GetTelegramUpdateFromSQSMessage(sqsMessage events.SQSMessage) (*tgbotapi.Update, error) {
	// 1. Formatting SQS Message body payload for unmarshal
	payload := sqsMessage.Body
	logging.Debugf("SQS message body pre-unmarshal: %s", payload)
	// slice off last char in record.Body; record.Body is invalid json due to extra lagging " char
	payload = payload[:len(payload)-1]
	logging.Debugf("SQS message body without quote pre-unmarshal: %s", payload)

	// 2. Execute unmarshal
	var unmarshaledSQSMessageBody SQSMessageBody
	if err := json.Unmarshal([]byte(payload), &unmarshaledSQSMessageBody); err != nil {
		logging.Errorf("error when unmarshaling Telegram Update object: %v", err)
		return nil, err
	}
	logging.Infof("Unmarshal sqsMessageBody post-unmarshal: %+v", unmarshaledSQSMessageBody)

	// 3. Return Update object
	update := unmarshaledSQSMessageBody.Body
	logging.LogUpdateObject(update)
	return &update, nil
}

// SQSMessageBody is the unmarshal destination for an SQSMessage.Body JSON payload.
type SQSMessageBody struct {
	Method      string          `json:"method"`
	Body        tgbotapi.Update `json:"body-json"`
	QueryParams QueryParams     `json:"queryParams"`
	PathParams  PathParms       `json:"pathParams"`
}

type QueryParams struct {
}

type PathParms struct {
}
