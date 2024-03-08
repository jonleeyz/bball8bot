package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// isDebugLoggingEnabled toggles debug logging on if true, and false otherwise. Is read from the respective Lambda env var.
var isDebugLoggingEnabled bool

func HandleRequest(ctx context.Context, event *events.SQSEvent) error {
	if event == nil {
		errString := "error: event is nil"
		log.Fatal(errString)
		return fmt.Errorf(errString)
	}

	token, err := getBotToken()
	if err != nil {
		log.Fatalf("error when retrieving Telegram bot token: %v", err)
		return err
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("error when creating Telegram bot object: %v", err)
		return err
	}

	logInManualDebugMode("Number of records in event: %d", len(event.Records))
	logInManualDebugMode("Event(s): %+v", event.Records)

	for _, sqsMessage := range event.Records {
		update, err := getTelegramUpdateFromSQSMessage(sqsMessage)
		if err != nil {
			log.Printf("error when unmarshaling SQS message: %v", err)
		}

		// if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)); err != nil {
		chattable := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		sendOutcome, err := bot.Send(chattable)
		if err != nil {
			log.Printf("error when calling Telegram Bot API to send message: %v", err)
			continue
		}

		logInManualDebugMode("Chattable object: %+v", chattable)
		logInManualDebugMode("bot.Send outcome: %+v", sendOutcome)
	}
	return nil
}

func main() {
	readDebugLoggingFlag()
	lambda.Start(HandleRequest)
}
