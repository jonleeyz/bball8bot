package main

import (
	"encoding/json"
	"net/http"
	"os"
)

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

type SecretResponse struct {
	SecretString string `json:"SecretString"`
}
