package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(ctx context.Context, update *tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	h, err := getUpdateHandler(ctx, bot, update)
	if err != nil {
		return err
	}
	return h.Handle(ctx)
}

type UpdateHandler interface {
	Handle(ctx context.Context) error
}
