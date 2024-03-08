package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SecretResponse struct {
	SecretString string `json:"SecretString"`
}

// getBotToken requests the Telegram bot token from AWS Secrets Manager and returns it for use.
// This implementation makes use of the AWS Parameters and Secrets Lambda Extension and only works when a Lambda function
// executes it.
// Ref: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieving-secrets_lambda.html
func getBotToken() (string, error) {
	client := &http.Client{}

	// 1. Prepare HTTP request to get bot token
	req, err := http.NewRequest("GET", "http://localhost:2773/secretsmanager/get?secretId=telegram_bot_token", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("X-Aws-Parameters-Secrets-Token", os.Getenv("AWS_SESSION_TOKEN"))

	// 2. Execute request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// 3. Unmarshal for token
	var secret SecretResponse
	if err := json.NewDecoder(resp.Body).Decode(&secret); err != nil {
		return "", err
	}

	return secret.SecretString, nil
}

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
		log.Fatalf("error when creating Telegram bot: %v", err)
		return err
	}

	logInManualDebugMode("Number of records in event: %d", len(event.Records))
	logInManualDebugMode("Eevent: %v", event.Records)

	for _, record := range event.Records {
		var update tgbotapi.Update
		if err := json.Unmarshal([]byte(record.Body), &update); err != nil {
			log.Printf("error while unmarshaling Telegram Update object: %v", err)
			continue
		}

		logInManualDebugMode("Update: %v", update)

		// if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)); err != nil {
		chattable := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		sendOutcome, err := bot.Send(chattable)
		if err != nil {
			log.Printf("error when calling Telegram Bot API to send message: %v", err)
			continue
		}

		logInManualDebugMode("Chattable object: %v", chattable)
		logInManualDebugMode("bot.Send outcome: %v", sendOutcome)
	}
	return nil
}

// logInManualDebugMode is a syntatic wrapper around the log.Printf function that creates a log entry with debug tags.
func logInManualDebugMode(message string, debugObject ...interface{}) {
	debugLog := fmt.Sprintf(message, debugObject...)
	log.Printf("[MANUAL_DEBUG_LOG] %s", debugLog)
}

func main() {
	lambda.Start(HandleRequest)
}
