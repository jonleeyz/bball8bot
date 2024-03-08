package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

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

func HandleRequest(ctx context.Context, event *MyEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	message := fmt.Sprintf("Hello %s!", event.Name)
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
