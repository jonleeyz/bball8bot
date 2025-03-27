package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/jonleeyz/bbball8bot/internal/json"
	"github.com/jonleeyz/bbball8bot/internal/logging"
	"github.com/jonleeyz/bbball8bot/internal/secrets"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	logging.Init()
	lambda.StartWithOptions(HandleRequest, lambda.WithContext(context.TODO()))
}

func HandleRequest(ctx context.Context, event *events.SQSEvent) error {
	if event == nil {
		errMessage := "error: event is nil"
		logging.Fatal(errMessage)
		return fmt.Errorf("%s", errMessage)
	}

	bot, err := getBot()
	if err != nil {
		return err
	}

	logging.Debugf("Number of records in event: %d", len(event.Records))
	logging.Debugf("Event(s): %+v", event.Records)

	for _, sqsMessage := range event.Records {
		update, err := json.GetTelegramUpdateFromSQSMessage(sqsMessage)
		if err != nil {
			continue
		}

		handleUpdate(ctx, update, bot)
	}

	// must return nil as Telegram will retry posting the Update to the webhook if something other than 2xx is returned.
	// Ref: https://core.telegram.org/bots/api#setwebhook
	return nil
}

// getBot retrieves the Telegram bot token and creates a valid bot API instance.
func getBot() (*tgbotapi.BotAPI, error) {
	token, err := secrets.GetBotToken()
	if err != nil {
		logging.Fatalf("error when retrieving Telegram bot token: %v", err)
		return nil, err
	}

	bot, err := tgbotapi.NewBotAPI(token)
	// TODO @jonlee: Refactor this into a proper config module
	bot.Debug = logging.IS_DEBUG_LOGGING_ENABLED
	if err != nil {
		logging.Fatalf("error when creating Telegram bot object: %v", err)
		return nil, err
	}
	return bot, nil
}
