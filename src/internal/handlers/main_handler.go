package handlers

import (
	"context"
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jonleeyz/bball8bot/internal/handlers/callbacks"
	"github.com/jonleeyz/bball8bot/internal/handlers/messages"
	"github.com/jonleeyz/bball8bot/internal/logging"
)

// HandleUpdate parses the input Update object, and responds accordingly.
// No error is returned from this function, any error is logged but is not fatal; other updates can still be processed.
// TODO @jonlee: Change signature to return error; not returning because of Telegram's requirement does not allow the function to be general purpose
func HandleUpdate(ctx context.Context, update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.CallbackQuery != nil {
		handler := callbacks.Init(bot)
		handler.Handle(ctx, update)
	}

	if update.Message == nil {
		return
	}

	messageHandler := messages.Init(bot)
	messageHandler.Handle(ctx, update)
}

// Pre-condition: Exactly 1 optional field in an Update object will be non-nil.
// Reference: https://core.telegram.org/bots/api#update
// getHandler returns a new handler appropriate to handle the input Update's content.
func getHandler(ctx context.Context, bot *tgbotapi.BotAPI, update *tgbotapi.Update) (UpdateHandler, error) {
	if isUpdateACallbackQuery(ctx, update) {
		return callbacks.Init(bot), nil
	}
	if isUpdateAMessage(ctx, update) {
		return messages.Init(bot), nil
	}

	logging.Errorf("no appropriate update handler found; update: %+v", *update)
	return nil, errors.New("no matching handler found")
}
