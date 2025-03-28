package handlers

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/internal/logging"
)

// HandleUpdate parses the input Update object, and responds accordingly.
// No error is returned from this function, any error is logged but is not fatal; other updates can still be processed.
// TODO @jonlee: Change signature to return error; not returning because of Telegram's requirement does not allow the function to be general purpose
func HandleUpdate(ctx context.Context, update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	h, err := getUpdateHandler(ctx, bot, update)
	if err != nil {
		logging.Errorf("%s", err.Error())
		return
	}
	h.Handle(ctx, update)
}

type UpdateHandler interface {
	Handle(ctx context.Context, update *tgbotapi.Update) error
}
