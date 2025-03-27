package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/jonleeyz/bbball8bot/commands"
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

	token, err := secrets.GetBotToken()
	if err != nil {
		logging.Fatalf("error when retrieving Telegram bot token: %v", err)
		return err
	}

	bot, err := tgbotapi.NewBotAPI(token)
	// TODO @jonlee: Refactor this into a proper config module
	bot.Debug = logging.IS_DEBUG_LOGGING_ENABLED
	if err != nil {
		logging.Fatalf("error when creating Telegram bot object: %v", err)
		return err
	}

	logging.Debugf("Number of records in event: %d", len(event.Records))
	logging.Debugf("Event(s): %+v", event.Records)

	for _, sqsMessage := range event.Records {
		update, err := json.GetTelegramUpdateFromSQSMessage(sqsMessage)
		if err != nil {
			logging.Errorf("error when unmarshaling SQS message: %v", err)
			continue
		} else {
			logging.LogUpdateObject(*update)
		}

		// TODO @jonlee: Update, placeholder, just to ensure that callback queries are answered.
		if update.CallbackQuery != nil {
			callback := update.CallbackQuery
			callbackResponseString := fmt.Sprintf("button pressed: %s", callback.Data)

			var callbackAnswer tgbotapi.CallbackConfig
			if callbackResponseString == "button pressed: ATTENDING" {
				callbackAnswer = tgbotapi.NewCallbackWithAlert(callback.ID, callbackResponseString)
			} else {
				callbackAnswer = tgbotapi.NewCallback(callback.ID, callbackResponseString)
			}

			if _, err := bot.Request(callbackAnswer); err != nil {
				logging.Errorf("error when answering callback: %v", err)
			}
			continue
		}

		if update.Message == nil {
			continue
		}

		// if message is command, call command handler
		if update.Message.IsCommand() {
			if err := commands.HandleBotCommand(ctx, bot, update); err != nil {
				// TODO @jonlee: Tidy this log statement
				logging.Errorf("TEMP TOP level log: %v", err)
			}
			continue
		}

		// if message is not command, echo message as reply to original message
		newReply := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		newReply.BaseChat.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(newReply); err != nil {
			logging.Errorf("error when calling Telegram Bot API to send message: %v", err)
		}
	}
	return nil
}
