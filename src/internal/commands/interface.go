package commands

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TODO @jonlee: Improve this description
// CommandHandler defines an abstraction layer to handle Telegram Bot Commands.
type CommandHandler interface {
	Init(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error
	Handle(ctx context.Context) error
}
